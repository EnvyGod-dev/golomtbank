--TABLE "entries"
BEGIN;

CREATE TABLE "entries" (
    "Id" BIGSERIAL NOT NULL,
    "FromAccountId" BIGINT NOT NULL,
    "BankName" VARCHAR(100) NOT NULL,
    "Amount" BIGINT NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT(NOW()),
    CONSTRAINT "Entries_Pk_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Entries_From_Account_Id" FOREIGN KEY ("FromAccountId") REFERENCES "accounts" ("Id"),
    CONSTRAINT "BankName_Check" CHECK (
        "BankName" IN (
            'Голомт Банк',
            'Хаан банк',
            'Mbank',
            'Төрийн банк',
            'Худалдаа хөгжлийн банк',
            'Богд Банк'
        )
    )
) TABLESPACE pg_default;

COMMIT;