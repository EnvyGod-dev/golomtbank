--TABLE "entries"
BEGIN;

CREATE TABLE "entries" (
    "Id" BIGSERIAL NOT NULL,
    "FromAccountId" BIGINT NOT NULL,
    "Amount" BIGINT NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT(now()),
    CONSTRAINT "Entries_Pk_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Entries_From_Account_Id" FOREIGN KEY ("FromAccountId") REFERENCES "accounts" ("Id")
) TABLESPACE pg_default;

COMMIT;