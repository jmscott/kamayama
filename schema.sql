/*
 *  Synopsis:
 *	Schema for web site kamayama.com.
 */

\set ON_ERROR_STOP on

BEGIN TRANSACTION;

DROP SCHEMA IF EXISTS kamayama CASCADE;
CREATE SCHEMA kamayama;

SET search_path TO kamayama;

--  output of golang function net/url.ParseRequestURI(rawURL string)

DROP DOMAIN IF EXISTS kamatime CASCADE;
CREATE DOMAIN kamatime AS timestamp CHECK (

	-- birthday of kamayama
	value >= '2026-05-14 02:14:15.406082-05'
) NOT NULL DEFAULT now();

DROP TABLE IF EXISTS golang_net_url_ParseRequestURI CASCADE;
CREATE TABLE golang_net_url_ParseRequestURI
(
	rawURL		text PRIMARY KEY,

	error		text,

	Scheme		text,

	--  the String() value of UserInfo structure
	UserInfo	text,
	Username	text,

	Host		text,
	Hostname	text,

	Path		text,
	Fragment	text,

	EscapedPath	text,
	EscapedFragment	text,

	RawQuery	text,
	RawPath		text,
	RawFragment	text,
	ForceQuery	bool,
	OmitHost	bool,

	IsAbs		bool,

	Port		text CHECK (
				Port ~ '^[0-9]{1,5}$'
				OR
				Port = ''
			),

	insert_time	kamatime,

	/*
	 *  Can have either an error and no other values
	 *  OR
	 *  No error and some values.
	 */
	--  insure we do not have both errors and parsed values.
	CONSTRAINT error_nullness CHECK ((
		--  error happened, so other fields must be null
		error IS NOT NULL
		AND
		(
			Scheme, UserInfo,
			Host, Hostname,
			Path, EscapedPath, EscapedFragment,
			RawQuery, RawPath, RawFragment,
			ForceQuery, OmitHost, IsAbs,
			Port
		) IS NULL
	     ) OR (
	     	--  no error so some of other fields must exist
	     	error IS NULL
		AND
		(
			Scheme, UserInfo,
			Host, Hostname,
			Path, EscapedPath, EscapedFragment,
			RawQuery, RawPath, RawFragment,
			ForceQuery, OmitHost, IsAbs,
			Port
		) IS NOT NULL
	    )
	)
);
CREATE INDEX idx_golang_net_url_ParseRequestURI_itime
  ON
  	golang_net_url_ParseRequestURI(insert_time)
;
COMMENT ON TABLE golang_net_url_ParseRequestURI
  IS
  	'Request URL Frisjed by golang fun net.url.ParseRequestURI()'
;

DROP TABLE IF EXISTS golang_net_url_Query;
CREATE TABLE golang_net_url_Query
(
	rawURL		text
				REFERENCES golang_net_url_ParseRequestURI,
	--  arg can be ""
	arg		text NOT NULL,

	PRIMARY KEY	(rawURL, arg)
);
COMMENT ON TABLE golang_net_url_Query
  IS
  	'Query args for frisked request url'
;

DROP TABLE IF EXISTS golang_net_url_Query_Values CASCADE;
CREATE TABLE golang_net_url_Query_Values
(
	rawURL		text,
	arg		text,
	array_order	smallint,

	value		text NOT NULL,

	PRIMARY KEY	(rawURL, arg, array_order),

	FOREIGN KEY (rawURL, arg)
			REFERENCES golang_net_url_Query
);

COMMIT TRANSACTION;
