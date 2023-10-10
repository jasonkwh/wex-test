CREATE TABLE wex.transactions (
    id varchar(64),
    transaction_date varchar(10),
    amount int,
    description varchar(50),
    transaction_type varchar(20)
);

-- permission
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA wex TO ${service_user};
GRANT CREATE ON SCHEMA wex TO ${service_user};
ALTER TABLE wex.transactions OWNER TO ${service_user};