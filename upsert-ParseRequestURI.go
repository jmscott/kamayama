/*
 *  Synopsis:
 *	Upsert rows from stdin into various tables kamayama.ParseRequestURI*
 *  Exit Code:
 *	0	all rows upserted.
 *	1	parse error on a url read from stdin. does not stop parsing.
 *	2	unexpeccted database error
 *	3	unexpected system error
 *  Note:
 *	Oops.  rename to upsert-ParseRequestURI
 */
package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"strings"
	"net/url"
)

/*
 *  0:	all urls parsed ok and upserted into db (OK)
 *  1:	at least one url parse failed, error upserted into db (WARN)
 *  2:	fatal sql error (ERROR)
 *  3:	fatal unexpected system error (ERROR)
 */
var exit_code		int

//  current line being scanned
var scan_line_no	uint64

//  current db connection.  pgx is single threaded.
var db		*pgx.Conn

//  background context required by pgx.
//
//  Note: can ctx be gotten from *db?
var ctx			context.Context

func die(format string, args ...interface{}) {
	
	fmt.Fprintf(os.Stderr, "ERROR: " + format + "\n", args...)
	os.Exit(3)
}

//  die() on sql error
func dieq(format string, args ...interface{}) {
	
	format = fmt.Sprintf(
			"ERROR: SQL(%s): %s",
			os.Getenv("PGDATABASE"),
			format,
	)

	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(2)
}

func warn(format string, args ...interface{}) {
	
	format = fmt.Sprintf(
			"WARN: ParseRequestURI(stdin) failed: line %d: %s\n",
			scan_line_no,
			format,
	)
	fmt.Fprintf(os.Stderr, format, args...)
	exit_code = 1
}

func warn_url(rawURL string, parse_err error) {

	warn("%s",  parse_err)
	_, err := db.Exec(ctx, `
			INSERT INTO golang_net_url_ParseRequestURI(
				rawURL, error
			) VALUES ($1, $2)
			  ON CONFLICT
			  	DO NOTHING
			;`, rawURL, parse_err.Error(),
	)
	if err != nil {
		dieq("db.Exec(error) failed: %s", err)
	}
}

func commit(what string, txn pgx.Tx) {
	
	err := txn.Commit(ctx)
	if err != nil {
		dieq("txn.Commit(%s) failed: %s", what, err)
	}
}

//  upsert an error tuple for a parsed url
func upsert(rawURL string) {

	txn, err := db.Begin(ctx)
	if err != nil {
		dieq("db.Begin(ctx) failed: %s", err)
	}
	defer txn.Rollback(ctx)		// becomes null-op after commit 

	_, err = db.Exec(ctx, `SET search_path TO kamayama`)
	if err != nil {
		dieq("db.Exec(search_path) failed: %s", err)
	}

	//  parse url and issue warning if parse fails.
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		warn_url(rawURL, err)
		commit("parse", txn)
		return
	}
	userinfo := ""
	username := ""

	ui := u.User
	if ui != nil {
		userinfo = ui.String()
		username = ui.Username()
	}

	//  sql INSERT parsed url and croak on any database error
	_, err = db.Exec(ctx, `
		INSERT INTO golang_net_url_ParseRequestURI(
			rawURL,
			Scheme, UserInfo, Username, Host, Hostname, Path,
			Fragment, EscapedPath, EscapedFragment, RawQuery,
			RawPath, RawFragment, ForceQuery, OmitHost, IsAbs,
			Port
		) VALUES (
			$1,
			$2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11,
			$12, $13, $14, $15, $16,
			$17
		) ON CONFLICT
			DO NOTHING
		;`,
		rawURL,
		u.Scheme, userinfo, username, u.Host, u.Hostname(), u.Path,
		u.Fragment, u.EscapedPath(), u.EscapedFragment(), u.RawQuery,
		u.RawPath, u.RawFragment, u.ForceQuery, u.OmitHost, u.IsAbs(),
		u.Port(),
	)
	if err != nil {
		dieq("db.Exec(ParseRequestURI) failed: %s", err)
	}

	//  INSERT into golang_net_url_Query
	q := u.Query()
	for a, vals := range q {
		_, err = db.Exec(ctx, `
			INSERT INTO
				golang_net_url_Query(rawURL, arg)
				VALUES ($1, $2)
			  ON CONFLICT
			  	DO NOTHING
			;
		`, rawURL, a)
		if err != nil {
			dieq("db.Exec(Query) failed: %s", err)
		}
		for i, v := range vals {
			_, err = db.Exec(
				ctx,
				`INSERT INTO
					golang_net_url_Query_Values(
						rawURL,
						arg,
						array_order,
						value
					) VALUES ($1, $2, $3, $4)
				;`,
				rawURL, a, i, v)
			if err != nil {
				dieq("db.Exec(QueryValues) failed: %s", err)
			}
		}
	}
	commit("upsert", txn)
}

func main() {

	var err error

	argc := len(os.Args)
	if argc != 1 {
		die("wrong cli arg count: got %d, want 0", argc)
	}

	ctx = context.Background()
	db, err = pgx.Connect(
			ctx,
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
				os.Getenv("PGUSER"),
				os.Getenv("PGPASSWORD"),
				os.Getenv("PGHOST"),
				os.Getenv("PGPORT"),
				os.Getenv("PGDATABASE"),
			),
	)
	if err != nil {
		dieq(
			"pgx.Connect(%s) failed: %s",
			os.Getenv("PGDATABASE"),
			err,
		)
	}
	defer db.Close(context.Background()) // Close connection when done

	in := bufio.NewReader(os.Stdin)
	for {
		scan_line_no++
		rawURL, err := in.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			die("in.ReadString(new-line) failed: %s", err)
		}
		rawURL = strings.TrimRight(rawURL, "\n")
		if rawURL == "" {
			die("stdin: empty string near line %d", scan_line_no)
		}
		upsert(rawURL)
	}
	os.Exit(exit_code)
}
