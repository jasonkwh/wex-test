DROP FUNCTION IF EXISTS wex.save_purchase;

CREATE FUNCTION wex.save_purchase(
    trans_id varchar, trans_date varchar, trans_amount int, trans_desc varchar
) RETURNS void
LANGUAGE plpgsql
AS $func$
BEGIN
    INSERT INTO wex.transactions(id,transaction_date,amount,description,transaction_type) 
    VALUES (trans_id, trans_date, trans_amount, trans_desc, "PURCHASE") 
    ON CONFLICT (id)
    DO UPDATE SET 
    transaction_date = EXCLUDED.transaction_date,
    amount = EXCLUDED.amount,
    description = EXCLUDED.description,
    transaction_type = EXCLUDED.transaction_type;
END;
$func$;

GRANT EXECUTE ON FUNCTION wex.save_purchase(trans_id varchar, trans_date varchar, trans_amount int, trans_desc varchar) TO ${service_user};