package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"strings"
)

func die(format string, args ...interface{}) {
	
	fmt.Fprintf(os.Stderr, "ERROR: " + format + "\n", args...)
	os.Exit(1)
}

func main() {

	argc := len(os.Args)
	if argc != 1 {
		die("wrong cli arg count: got %d, want 0", argc)
	}

	conn, err := pgx.Connect(
			context.Background(),
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
				os.Getenv("PGUSER"),
				os.Getenv("PGPASSWORD"),
				os.Getenv("PGHOST"),
				os.Getenv("PGPORT"),
				os.Getenv("PGDATABASE"),
			),
	)
	if err != nil {
		die(
			"pgx.Connect(%s) failed: %s",
			os.Getenv("PGDATABASE"),
			err,
		)
	}
	defer conn.Close(context.Background()) // Close connection when done

	in := bufio.NewReader(os.Stdin)
	line_no := 0
	for {
		line_no++
		rawURL, err := in.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			die("in.ReadString(new-line) failed: %s", err)
		}
		rawURL = strings.TrimRight(rawURL, "\n")
		if rawURL == "" {
			die("stdin: empty string near line %d", line_no)
		}

	}

	os.Exit(0)
}
