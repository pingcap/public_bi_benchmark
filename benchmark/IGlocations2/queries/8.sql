SELECT "IGlocations2_1"."media_type" AS "media_type",   "IGlocations2_1"."state" AS "state",   SUM(CAST("IGlocations2_1"."Number of Records" AS SIGNED)) AS "sum:Number of Records:ok" FROM "IGlocations2_1" WHERE ("IGlocations2_1"."state" IN ('Alabama', 'Alaska', 'Arizona', 'Arkansas', 'California', 'Colorado', 'Connecticut', 'Delaware', 'Florida', 'Georgia', 'Hawaii', 'Idaho', 'Illinois', 'Indiana', 'Iowa', 'Kansas')) GROUP BY "IGlocations2_1"."media_type",   "IGlocations2_1"."state";
