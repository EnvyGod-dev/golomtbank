-- TABLE "accounts"
BEGIN;

CREATE TABLE "accounts" (
    "Id" BIGINT NOT NULL,
    "Balance" BIGINT NOT NULL,
    "Owner" VARCHAR(100) DEFAULT('') NOT NULL,
    "BankName" VARCHAR(100) NOT NULL,
    "Currency" VARCHAR(5) DEFAULT('') NOT NULL,
    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT(NOW()),
    CONSTRAINT "ACCOUNTS_PK_Id" PRIMARY KEY ("Id"),
    CONSTRAINT "Accounts_Owner_FK" FOREIGN KEY ("Owner") REFERENCES "User" ("Username"),
    CONSTRAINT "Accounts_Owner_Currency" UNIQUE ("Owner", "Currency"),
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