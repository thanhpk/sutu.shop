package model

import (
	"time"
)

type ISaleRepository interface {
	ListSaleAfterTime(time.Time) []Sale
}

type SaleMgt struct {
	Repo ISaleRepository
}

func (me *SaleMgt) ListRunningSales() []Sale {
	return me.Repo.ListSaleAfterTime(time.Now())
}
