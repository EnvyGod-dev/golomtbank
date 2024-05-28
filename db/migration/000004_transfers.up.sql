-- TABLE "transfer"

BEGIN;

CREATE TABLE "transfer" (
    "Id" BIGSERIAL NOT NULL,
    "FromAccountId" BIGINT NOT NULL,
    "ToAccountId" BIGINT NOT NULL,
    "Amount" BIGINT NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT(now()),
    CONSTRAINT "transfers_PK_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Transfers_From_Account_Id" FOREIGN KEY ("FromAccountId") REFERENCES "accounts" ("Id"),
    CONSTRAINT "Transfers_To_Account_Id" FOREIGN KEY ("ToAccountId") REFERENCES "accounts" ("Id")
) TABLESPACE pg_default;

COMMIT;