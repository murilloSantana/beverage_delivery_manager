package repository

import "beverage_delivery_manager/pdv/domain"

type PdvRepository interface {
	Save(pdv domain.Pdv) (domain.Pdv, error)
}
