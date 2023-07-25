SELECT
	"country",
	COALESCE(SUM ( belanjas.total_buy ),0) AS total_spend 
FROM
	"public"."users"
	LEFT JOIN PUBLIC.belanjas ON users.ID = belanjas.id_user 
GROUP BY
	"country"
ORDER BY
	total_spend desc