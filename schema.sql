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

DROP TABLE IF EXISTS golang_net_url_ParseRequestURI CASCADE;
CREATE TABLE golang_net_url_ParseRequestURI
(
	rawURL		text PRIMARY KEY,
	error		text,

	Schema		text NOT NULL,

	--  the String() value of UserInfo structure
	UserInfo	text,

	Host		text NOT NULL,
	Hostname	text NOT NULL,

	Path		text NOT NULL,
	Fragment	text NOT NULL,

	EscapedPath	text NOT NULL,
	EscapedFragment	text NOT NULL,

	RawQuery	text NOT NULL,
	RawPath		text NOT NULL,
	RawFragment	text NOT NULL,
	ForceQuery	bool NOT NULL,
	OmitHost	bool NOT NULL,

	IsAbs		bool NOT NULL,

	Port		text NOT NULL,

	insert_time	timestamptz
				NOT NULL
				DEFAULT now()
);
CREATE INDEX idx_golang_net_url_ParseRequestURI_itime
  ON
  	golang_net_url_ParseRequestURI(insert_time)
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

DROP TABLE IF EXISTS golang_net_url_Query_Values CASCADE;
CREATE TABLE golang_net_url_Query_Values
(
	rawURL		text,
	arg		text,

	array_order	smallint CHECK (
				array_order >= 0
			),

	value		text NOT NULL,

	PRIMARY KEY	(rawURL, arg, array_order)
);

COMMIT TRANSACTION;
