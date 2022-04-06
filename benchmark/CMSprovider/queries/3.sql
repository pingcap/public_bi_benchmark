SELECT "CMSprovider_2"."HCPCS_DESCRIPTION" AS "HCPCS_DESCRIPTION",   CAST("CMSprovider_2"."NPPES_PROVIDER_ZIP" AS SIGNED) AS "NPPES_PROVIDER_ZIP",   SUM("CMSprovider_2"."AVERAGE_MEDICARE_PAYMENT_AMT") AS "sum:AVERAGE_MEDICARE_PAYMENT_AMT:ok" FROM "CMSprovider_2" WHERE ("CMSprovider_2"."NPPES_PROVIDER_STATE" = 'CA') GROUP BY "CMSprovider_2"."HCPCS_DESCRIPTION",   "CMSprovider_2"."NPPES_PROVIDER_ZIP",   "CMSprovider_2"."NPPES_PROVIDER_ZIP";
