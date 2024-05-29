-- TABLE "transfer"
BEGIN;

CREATE TABLE "transfer" (
    "Id" BIGSERIAL NOT NULL,
    "FromAccountId" BIGINT NOT NULL,
    "BankName" VARCHAR(100) NOT NULL,
    "ToAccountId" BIGINT NOT NULL,
    "Amount" BIGINT NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "transfers_PK_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Transfers_From_Account_Id" FOREIGN KEY ("FromAccountId") REFERENCES "accounts" ("Id"),
    CONSTRAINT "Transfers_To_Account_Id" FOREIGN KEY ("ToAccountId") REFERENCES "accounts" ("Id"),
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