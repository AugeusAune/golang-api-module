ALTER TABLE
    users
ADD
    COLUMN company_id UUID NOT NULL;

ALTER TABLE
    users
ADD
    CONSTRAINT fk_users_company FOREIGN KEY (company_id) REFERENCES companies(id)
