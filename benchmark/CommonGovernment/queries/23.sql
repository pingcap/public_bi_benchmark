SELECT CAST(EXTRACT(YEAR FROM (cast("CommonGovernment_4"."signeddate" as DATE) + INTERVAL '3' MONTH)) AS SIGNED) AS "yr:signeddate:ok" FROM "CommonGovernment_4" GROUP BY "yr:signeddate:ok";
