SELECT "Generico_3"."Anunciante" AS "Anunciante" FROM "Generico_3" WHERE (("Generico_3"."Anunciante" IN ('BANTRAB/TODOTICKET', 'TODOTICKET', 'TODOTICKET.COM')) AND (CAST(EXTRACT(YEAR FROM "Generico_3"."FECHA") AS SIGNED) >= 2010) AND (CAST(EXTRACT(YEAR FROM "Generico_3"."FECHA") AS SIGNED) <= 2015)) GROUP BY "Generico_3"."Anunciante" ORDER BY "Anunciante" ASC ;
