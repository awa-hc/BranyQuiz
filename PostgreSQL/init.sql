-- Verificar si la base de datos ya existe
SELECT CASE WHEN EXISTS (
    SELECT 1
    FROM pg_catalog.pg_database
    WHERE datname = 'brainy'
) THEN 0 ELSE 1 END AS db_exists \gset

-- Si la base de datos no existe, entonces crearla
\if :db_exists
    CREATE DATABASE brainy;
\endif
