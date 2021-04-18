// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "beverage_delivery_manager/pdv/domain"

	mock "github.com/stretchr/testify/mock"
)

// PdvUseCase is an autogenerated mock type for the PdvUseCase type
type PdvUseCase struct {
	mock.Mock
}

// Save provides a mock function with given fields: pdv
func (_m *PdvUseCase) Save(pdv domain.Pdv) (domain.Pdv, error) {
	ret := _m.Called(pdv)

	var r0 domain.Pdv
	if rf, ok := ret.Get(0).(func(domain.Pdv) domain.Pdv); ok {
		r0 = rf(pdv)
	} else {
		r0 = ret.Get(0).(domain.Pdv)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Pdv) error); ok {
		r1 = rf(pdv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
