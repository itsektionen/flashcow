BEGIN;
CREATE TABLE committee (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	full_name text NOT NULL,
	short_name text NOT NULL
);

CREATE TABLE user_details (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	time timestamptz NOT NULL,
	full_name text NOT NULL,
	chapter_email_address text UNIQUE NOT NULL
);

CREATE TABLE receipt_report (
	id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	purchase_date date NOT NULL,
	submission_time timestamptz NOT NULL,
	submitter_id int NOT NULL REFERENCES user_details(id),
	committee_id int REFERENCES committee(id),
	contents text NOT NULL,
	food_total int NOT NULL,
	beer_total int NOT NULL,
	soda_total int NOT NULL,
	cider_total int NOT NULL,
	wine_total int NOT NULL,
	spirits_total int NOT NULL,
	material_total int NOT NULL,
	other_total int NOT NULL,
	internrep_totel int NOT NULL,
	comments text NOT NULL,
	image bytea NOT NULL
);

CREATE TABLE migrations (
	migration_id int PRIMARY KEY
);

INSERT INTO migrations VALUES (1);

COMMIT;
