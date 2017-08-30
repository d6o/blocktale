// Code generated by mockery v1.0.0
package mocks

import banana "github.com/makarski/mybanana/db/banana"
import mock "github.com/stretchr/testify/mock"

// BananaFinder is an autogenerated mock type for the BananaFinder type
type BananaFinder struct {
	mock.Mock
}

// Find provides a mock function with given fields: ID
func (_m *BananaFinder) Find(ID uint64) (*banana.Banana, error) {
	ret := _m.Called(ID)

	var r0 *banana.Banana
	if rf, ok := ret.Get(0).(func(uint64) *banana.Banana); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*banana.Banana)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
