-- +goose Up
CREATE TABLE IF NOT EXISTS search_providers (
  id         INTEGER PRIMARY KEY,                 
  category   TEXT    NOT NULL DEFAULT '',
  domain     TEXT    NOT NULL DEFAULT '',
  rank       INTEGER NOT NULL DEFAULT 0,
  site_name  TEXT    NOT NULL DEFAULT '',
  subcategory TEXT   NOT NULL DEFAULT '',
  tag        TEXT    NOT NULL UNIQUE,
  url        TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS history (
  id           INTEGER PRIMARY KEY,
  query        TEXT NOT NULL DEFAULT '',
  provider_id  INT  NOT NULL DEFAULT 0,
  provider_tag TEXT NOT NULL DEFAULT '',
  timestamp    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (provider_id) REFERENCES providers(id)
);

CREATE TABLE IF NOT EXISTS shortcuts (
  id         INTEGER  PRIMARY KEY,
  alias      TEXT NOT NULL UNIQUE,
  url        TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS sessions (
  id         INTEGER PRIMARY KEY,
  alias      TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
