ALTER TABLE accounts
    DROP CONSTRAINT IF EXISTS uq_accounts_user_name;

ALTER TABLE accounts
    ADD CONSTRAINT uq_accounts_user_name UNIQUE (user_id, name);