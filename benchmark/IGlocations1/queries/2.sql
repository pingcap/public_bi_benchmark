SELECT "IGlocations1_1"."State" AS "State",   SUM(CAST("IGlocations1_1"."POPESTIMATE2014" AS SIGNED)) AS "TEMP(Calculation_8840730155303010)(3219150654)(0)" FROM "IGlocations1_1" WHERE ("IGlocations1_1"."State" NOT IN ('District of Columbia')) GROUP BY "IGlocations1_1"."State";
