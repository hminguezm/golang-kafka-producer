package service

import (
  "fmt"
)

func GetProductToSendSql(lastExec string) string {
  slq := fmt.Sprintf(
    `SELECT product_sku,
       product_sku_parent,
       variant_id,
       category_code,
       tittle,
       short_description,
       long_description,
       brand,
       seller_id,
       offer_id,
       part_number,
       '{' || LISTAGG(attributes, ', ') WITHIN GROUP (ORDER BY product_sku) || '}' AS "ATTRIBUTES",
       media,
       created_at,
       updated_at
FROM (
         SELECT DISTINCT cp.PRODUCT_SKU                                                  AS product_sku,
                         cp.PRODUCT_SKU_PARENT                                           AS product_sku_parent,
                         cp.VARIANT_ID                                                   AS variant_id,
                         cp.CATEGORY_CODE                                                AS category_code,
                         cp.NAME                                                         AS tittle,
                         cp.SHORTDESCRIPTION                                             AS short_description,
                         NULLIF(DBMS_LOB.SUBSTR(cp.LONGDESCRIPTION, 250, 1), 'NULL')     AS long_description,
                         cp.BRAND                                                        AS brand,
                         cp.SELLER_ID                                                    AS seller_id,
                         co.OFFERS_ID                                                    AS offer_id,
                         cpa.PARTNUMBER                                                  AS part_number,
                         cp.CREATE_DATE                                                  AS created_at,
                         cp.LAST_DATE                                                    AS updated_at,
                         COALESCE('"' || cpa.NAME || '": "' || cpa.VALUE || '"', 'NULL') AS attributes,
                         '[{"thumbnail": "' || cp.THUMBNAIL || '" }, { "full_image": "' || cp.FULLIMAGE ||
                         '"}]'                                                           AS media
         FROM CON_PRODUCTS cp
                  INNER JOIN CON_PRODUCT_ATTRIBUTE cpa ON cpa.PRODUCT_SKU = cp.PRODUCT_SKU
                  LEFT OUTER JOIN CON_OFFERS co ON co.CON_PRODUCT_PRODUCT_SKU = cp.PRODUCT_SKU
         WHERE (cp.SYNC = 1)
           AND (cp.CREATE_DATE >= to_date('%s', 'YYYY-MM-DD"T"HH24:MI:SS'))
           AND (cp.CREATE_DATE < SYSDATE)
         GROUP BY cp.PRODUCT_SKU,
                  cp.PRODUCT_SKU_PARENT,
                  cp.VARIANT_ID,
                  cp.CATEGORY_CODE,
                  cp.NAME,
                  cp.SHORTDESCRIPTION,
                  DBMS_LOB.SUBSTR(cp.LONGDESCRIPTION, 250, 1),
                  cp.THUMBNAIL,
                  cp.FULLIMAGE,
                  cp.BRAND,
                  cp.SELLER_ID,
                  cp.LAST_DATE,
                  co.OFFERS_ID,
                  cpa.NAME,
                  cpa.PARTNUMBER,
                  cpa.VALUE,
                  cp.CREATE_DATE
     )
GROUP BY product_sku,
         product_sku_parent,
         variant_id,
         category_code,
         tittle,
         short_description,
         long_description,
         brand,
         seller_id,
         offer_id,
         part_number,
         created_at,
         updated_at,
         media`, lastExec)

  return slq
}
