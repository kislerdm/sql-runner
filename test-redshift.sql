CREATE TEMP TABLE mkt_test AS
WITH
    tmp AS (
        SELECT 1            AS col_1
             , CURRENT_DATE AS date
    )
SELECT *
FROM tmp
;

UNLOAD
('SELECT * FROM mkt_test')
TO '{destination}'
{autorization}
HEADER
FORMAT AS CSV
DELIMITER AS ','
NULL AS ''
ALLOWOVERWRITE
PARALLEL OFF
;
