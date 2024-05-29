-- TABLE "accounts"
BEGIN;

CREATE TABLE "accounts" (
    "Id" BIGSERIAL NOT NULL,
    "Balance" BIGINT NOT NULL DEFAULT 5000,
    "Owner" VARCHAR(100) NOT NULL,
    "BankName" VARCHAR(100) NOT NULL,
    "Currency" VARCHAR(5) NOT NULL DEFAULT '',
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "Accounts_Pk_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Accounts_Owner_Fk" FOREIGN KEY ("Owner") REFERENCES "User" ("Username"),
    CONSTRAINT "Accounts_BankName_Check" CHECK (
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