SELECT
	"credit_card_type",
	count(credit_card_type) AS total
FROM
	"public"."users"
GROUP BY
	"credit_card_type"
ORDER BY
	total desc
Limit 1