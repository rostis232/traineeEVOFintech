CREATE TABLE transactions
(
    transaction_id integer primary key unique,
    request_id integer,
    terminal_id integer,
    partner_object_id integer,
    amount_total decimal(8,2),
    amount_original decimal(8,2),
    commission_ps decimal(8,2),
    commission_client decimal(8,2),
    commission_provider decimal(8,2),
    date_input timestamp,
    date_post timestamp,
    status varchar(8),
    payment_type varchar(4),
    payment_number varchar(10),
    service_id integer,
    service varchar(20),
    payee_id integer,
    payee_name varchar(10),
    payee_bank_mfo integer,
    payee_bank_account varchar(20),
    payment_narrative varchar(250)
);

-- CREATE INDEX