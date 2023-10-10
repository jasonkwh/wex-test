DROP FUNCTION IF EXISTS wex.get_purchase;

CREATE FUNCTION wex.get_purchase(trans_id varchar)
RETURNS TABLE (transaction_date varchar, amount int, description varchar)
LANGUAGE plpgsql
AS $func$
BEGIN
    RETURN QUERY
    SELECT
        ts.transaction_date,
        ts.amount,
        ts.description
    FROM wex.transactions AS ts
    WHERE ts.id = trans_id AND ts.transaction_type = 'PURCHASE';
END;
$func$;

GRANT EXECUTE ON FUNCTION wex.get_purchase(trans_id varchar) TO ${service_user};