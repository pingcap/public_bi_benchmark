SELECT "Motos_2"."Marca" AS "Datos (copia)",   MAX("Motos_2"."Vehiculo") AS "TEMP(attr:Vehiculo:nk)(1662645443)(0)",   MIN("Motos_2"."Vehiculo") AS "TEMP(attr:Vehiculo:nk)(536654816)(0)",   SUM((CAST("Motos_2"."Cols" AS SIGNED) * CAST("Motos_2"."Plgs" AS SIGNED))) AS "sum:Calculation_1450626233922327:ok",   SUM(CAST("Motos_2"."NumAnuncios" AS SIGNED)) AS "sum:NumAnuncios:ok",   CAST(EXTRACT(YEAR FROM "Motos_2"."FECHA") AS SIGNED) AS "yr:FECHA:ok" FROM "Motos_2" WHERE ((CAST(EXTRACT(YEAR FROM "Motos_2"."FECHA") AS SIGNED) = 2015) AND ("Motos_2"."Categoria" = 'MOTOCICLETAS') AND ("Motos_2"."Medio" = 'PRENSA')) GROUP BY "Motos_2"."Marca",   "yr:FECHA:ok";
