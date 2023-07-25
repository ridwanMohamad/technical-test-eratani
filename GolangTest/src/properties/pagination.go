package properties

import (
	"math"
)

type Pagination struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
}

func (p *Pagination) CountTotalPage() {
	p.TotalPage = int64(math.Ceil(float64(p.TotalData) / float64(p.Limit)))
}
