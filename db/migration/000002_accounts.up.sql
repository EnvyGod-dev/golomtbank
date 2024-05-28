-- TABLE "accounts"
BEGIN;

CREATE TABLE "accounts" (
    "Id" BIGSERIAL NOT NULL,
    "Balance" BIGINT NOT NULL,
    "Owner" VARCHAR(100) DEFAULT('') NOT NULL,
    "Currency" VARCHAR(5) DEFAULT('') NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT(NOW()),
    CONSTRAINT "ACCOUNTS_PK_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Accounts_Owner_FK" FOREIGN KEY ("Owner") REFERENCES "User" ("Username"),
    CONSTRAINT "Accounts_Owner_Currency" UNIQUE ("Owner", "Currency")
) TABLESPACE pg_default;

COMMIT;