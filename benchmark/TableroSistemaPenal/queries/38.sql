SELECT "TableroSistemaPenal_8"."PAÍS" AS "PAÍS",   "TableroSistemaPenal_8"."Tipo Sentencia" AS "Tipo Sentencia",   SUM(CAST("TableroSistemaPenal_8"."Number of Records" AS SIGNED)) AS "sum:Number of Records:ok",   CAST(EXTRACT(YEAR FROM "TableroSistemaPenal_8"."FECHA") AS SIGNED) AS "yr:FECHA:ok" FROM "TableroSistemaPenal_8" WHERE ("TableroSistemaPenal_8"."Tipo Sentencia" IN ('Sentencia Absolutoria', 'Sentencia Condenatoria')) GROUP BY "TableroSistemaPenal_8"."PAÍS",   "TableroSistemaPenal_8"."Tipo Sentencia", "yr:FECHA:ok";
