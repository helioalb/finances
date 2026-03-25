CREATE TABLE IF NOT EXISTS transactions (
    uuid UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),

    account_id BIGINT NOT NULL,

    amount INT NOT NULL,
    description VARCHAR(255),
    type VARCHAR(10) NOT NULL, 

    CONSTRAINT ck_transactions_type
        CHECK (type IN ('EXPENSE', 'INCOME', 'TRANSFER')),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_transactions_account
        FOREIGN KEY (account_id)
        REFERENCES accounts(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_transactions_account_id
    ON transactions(account_id);

CREATE INDEX IF NOT EXISTS idx_transactions_account_id_created_at_desc
    ON transactions(account_id, created_at DESC);