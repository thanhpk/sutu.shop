package model

type IVarianceTypeRepo interface {
	Read(id string) *VarianceType
}

type VarianceTypeMgt struct {
	Repo IVarianceTypeRepo
}

func (me *VarianceTypeMgt) Read(id string) *VarianceType {
	return me.Repo.Read(id)
} 
