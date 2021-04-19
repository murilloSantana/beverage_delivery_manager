// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "beverage_delivery_manager/pdv/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PdvRepository is an autogenerated mock type for the PdvRepository type
type PdvRepository struct {
	mock.Mock
}

// FindByAddress provides a mock function with given fields: point
func (_m *PdvRepository) FindByAddress(point domain.Point) (domain.Pdv, error) {
	ret := _m.Called(point)

	var r0 domain.Pdv
	if rf, ok := ret.Get(0).(func(domain.Point) domain.Pdv); ok {
		r0 = rf(point)
	} else {
		r0 = ret.Get(0).(domain.Pdv)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Point) error); ok {
		r1 = rf(point)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ID
func (_m *PdvRepository) FindByID(ID string) (domain.Pdv, error) {
	ret := _m.Called(ID)

	var r0 domain.Pdv
	if rf, ok := ret.Get(0).(func(string) domain.Pdv); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Pdv)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateNewID provides a mock function with given fields:
func (_m *PdvRepository) GenerateNewID() func() string {
	ret := _m.Called()

	var r0 func() string
	if rf, ok := ret.Get(0).(func() func() string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() string)
		}
	}

	return r0
}

// HasDocument provides a mock function with given fields: document
func (_m *PdvRepository) HasDocument(document string) (bool, error) {
	ret := _m.Called(document)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(document)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, pdv, generateNewID
func (_m *PdvRepository) Save(ctx context.Context, pdv domain.Pdv, generateNewID func() string) (domain.Pdv, error) {
	ret := _m.Called(ctx, pdv, generateNewID)

	var r0 domain.Pdv
	if rf, ok := ret.Get(0).(func(context.Context, domain.Pdv, func() string) domain.Pdv); ok {
		r0 = rf(ctx, pdv, generateNewID)
	} else {
		r0 = ret.Get(0).(domain.Pdv)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Pdv, func() string) error); ok {
		r1 = rf(ctx, pdv, generateNewID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}