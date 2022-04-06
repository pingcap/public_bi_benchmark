SELECT SUM(CAST("HashTags_1"."Number of Records" AS SIGNED)) AS "sum:Number of Records:ok",   "HashTags_1"."twitter#user#screen_name" AS "twitter#user#screen_name" FROM "HashTags_1" GROUP BY "HashTags_1"."twitter#user#screen_name" HAVING (SUM(1) > 200);
