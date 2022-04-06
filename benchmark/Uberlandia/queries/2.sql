SELECT "Uberlandia_1"."Calculation_838513978982854656" AS "Calculation_838513978982854656",   COUNT(DISTINCT "Uberlandia_1"."Calculation_838513981462429699") AS "ctd:Calculation_838513981462429699:ok",   "Uberlandia_1"."municipio_do_local_da_oferta" AS "municipio_do_local_da_oferta",   "Uberlandia_1"."nome_curso_catalogo_guia" AS "nome_curso_catalogo_guia",   "Uberlandia_1"."nome_da_ue" AS "nome_da_ue",   "Uberlandia_1"."subtipo_curso" AS "subtipo_curso",   "Uberlandia_1"."uf_do_local_da_oferta" AS "uf_do_local_da_oferta",   CAST(EXTRACT(YEAR FROM "Uberlandia_1"."data_de_inicio") AS SIGNED) AS "yr:data_de_inicio:ok" FROM "Uberlandia_1" WHERE (("Uberlandia_1"."municipio_do_local_da_oferta" = 'Uberlândia') AND ("Uberlandia_1"."subtipo_curso" = 'FIC') AND ("Uberlandia_1"."uf_do_local_da_oferta" = 'MG') AND (CAST(EXTRACT(YEAR FROM "Uberlandia_1"."data_de_inicio") AS SIGNED) IN (2011, 2012, 2013, 2014)) AND (NOT (("Uberlandia_1"."nome da sit matricula (situacao detalhada)" IN ('CANC_DESISTENTE', 'CANC_MAT_PRIM_OPCAO', 'CANC_SANÇÃO', 'CANC_SEM_FREQ_INICIAL', 'CANC_TURMA', 'DOC_INSUFIC', 'ESCOL_INSUFIC', 'INC _ITINERARIO', 'INSC_CANC', 'Não Matriculado', 'NÃO_COMPARECEU', 'TURMA_CANC', 'VAGAS_INSUFIC')) OR ("Uberlandia_1"."nome da sit matricula (situacao detalhada)" IS NULL))) AND (NOT ("Uberlandia_1"."situacao_da_turma" IN ('CANCELADA', 'CRIADA', 'PUBLICADA')))) GROUP BY "Calculation_838513978982854656",   "municipio_do_local_da_oferta",   "nome_curso_catalogo_guia",   "nome_da_ue",   "subtipo_curso",   "uf_do_local_da_oferta",   "yr:data_de_inicio:ok";
