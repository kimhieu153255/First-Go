CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL UNIQUE,
  "full_name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "role" varchar NOT NULL DEFAULT 'user'
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'usd',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bills" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'usd',
  "name" varchar NOT NULL,
  "tag" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "due_date" timestamptz NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'usd',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigserial NOT NULL,
  "to_account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'usd',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "bills" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "accounts" ("user_id");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "bills" ("account_id");

CREATE INDEX ON "bills" ("tag");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';