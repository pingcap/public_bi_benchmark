SELECT "Generico_5"."Medio" AS "Medio" FROM "Generico_5" WHERE (("Generico_5"."Anunciante" IN ('BANTRAB/TODOTICKET', 'TODOTICKET', 'TODOTICKET.COM')) AND (CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) >= 2010) AND (CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) <= 2015) AND (CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) IN (2014, 2015)) AND (((CAST(EXTRACT(MONTH FROM "Generico_5"."FECHA") AS SIGNED) >= 1) AND (CAST(EXTRACT(MONTH FROM "Generico_5"."FECHA") AS SIGNED) <= 7)) OR (CAST(EXTRACT(MONTH FROM "Generico_5"."FECHA") AS SIGNED) IS NULL))) GROUP BY "Generico_5"."Medio";
