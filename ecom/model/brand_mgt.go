package model

type IBrandRepository interface {
	Read(id string) *Brand
}

type BrandMgt struct {
	Repo IBrandRepository
}

func (me *BrandMgt) Read(id string) *Brand {
	return me.Repo.Read(id)
}
