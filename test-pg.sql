CREATE TABLE IF NOT EXISTS public.test_table AS 
SELECT 1            AS a
    , CURRENTE_DATE AS date
;

COPY public.test_table 
TO '{destination}' 
DELIMITER ',' 
CSV 
HEADER
;