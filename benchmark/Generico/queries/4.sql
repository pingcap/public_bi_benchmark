SELECT "Generico_5"."Anunciante" AS "Anunciante",   "Generico_5"."Anunciante" AS "Datos (copia)",   SUM("Generico_5"."InversionUS") AS "TEMP(TC_)(2622528870)(0)",   SUM("Generico_5"."InversionUS") AS "sum:Calculation_0061002123102817:ok",   CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) AS "yr:FECHA:ok" FROM "Generico_5" WHERE (("Generico_5"."Anunciante" IN ('BANTRAB/TODOTICKET', 'TODOTICKET', 'TODOTICKET.COM')) AND (CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) >= 2010) AND (CAST(EXTRACT(YEAR FROM "Generico_5"."FECHA") AS SIGNED) <= 2015)) GROUP BY "Generico_5"."Anunciante",   "yr:FECHA:ok",   "Generico_5"."Anunciante";
