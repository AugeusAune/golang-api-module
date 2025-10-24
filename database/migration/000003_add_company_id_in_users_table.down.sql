ALTER TABLE
    users DROP CONSTRAINT IF EXISTS fk_users_company;

ALTER TABLE
    users DROP COLUMN IF EXISTS company_id;
