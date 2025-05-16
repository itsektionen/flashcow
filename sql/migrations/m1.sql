BEGIN;
CREATE TYPE currency AS ENUM ('SEK');

CREATE TABLE committee (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	full_name text NOT NULL,
	short_name text NOT NULL
);

CREATE TABLE budget (
	name text NOT NULL,
	start_time timestamptz NOT NULL,
	end_time timestamptz NOT NULL
);

CREATE TABLE budget_entry (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	committee_id int NOT NULL REFERENCES committee(id),
	name text NOT NULL,
	budgeted_amount int NOT NULL
);

CREATE TABLE receipt_report (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	submitter_name text NOT NULL,
	committee_id int NOT NULL REFERENCES committee(id),
	contents text NOT NULL,
	currency currency NOT NULL,
	total int NOT NULL
);

CREATE TABLE receipt_report_partial_sum (
	receipt_report_id int NOT NULL REFERENCES receipt_report(id),
	budget_entry_id int NOT NULL REFERENCES budget_entry(id),
	currency currency NOT NULL,
	amount int NOT NULL
);

COMMIT;
