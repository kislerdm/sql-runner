CREATE TABLE IF NOT EXISTS public.test_table AS 
SELECT 1     AS a
    , {col1} AS date
;

UNLOAD ('select * from public.test_table')
TO 's3://{bucket_path}'
FORMAT      AS CSV
DELIMITER   AS ','
;