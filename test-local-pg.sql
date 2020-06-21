CREATE TABLE IF NOT EXISTS public.test_table AS 
SELECT 1     AS a
    , {col1} AS date
;

COPY public.test_table TO '/data/test.csv' DELIMITER ',' CSV HEADER
;