-- CONNECT TO THE DATABASE
\c :_db;

SET myvars._db TO :'_db';
SET myvars._dbpass TO :'_dbpass';
SET myvars._dbuser TO :'_dbuser';
SET myvars._user TO :'_user';

-- TEST IF THE WEX SCHEMA ALREADY EXISTS
DO $$
DECLARE
  _db TEXT := current_setting('myvars._db', true);
  _dbpass TEXT := current_setting('myvars._dbpass', true);
  _dbuser TEXT := current_setting('myvars._dbuser', true);
  _user TEXT := current_setting('myvars._user', true);
BEGIN
  IF NOT EXISTS(
      SELECT schema_name
        FROM information_schema.schemata
        WHERE schema_name = 'wex'
    )
  THEN
    CREATE SCHEMA IF NOT EXISTS wex;
     IF NOT EXISTS (
       SELECT
       FROM   pg_catalog.pg_roles
       WHERE  rolname = _dbuser)
     THEN
       EXECUTE 'CREATE USER ' || _dbuser || ' PASSWORD ''' || _dbpass || '''';
     END IF;
    EXECUTE 'ALTER SCHEMA wex OWNER TO ' ||  _user;
    EXECUTE 'GRANT USAGE ON SCHEMA wex TO ' || _dbuser || ';';
  END IF;
END
$$;