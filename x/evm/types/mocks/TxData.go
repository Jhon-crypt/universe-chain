// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/evmos/v18/x/evm/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxData is an autogenerated mock type for the TxData type
type TxData struct {
	mock.Mock
}

// AsEthereumData provides a mock function with given fields:
func (_m *TxData) AsEthereumData() types.TxData {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AsEthereumData")
	}

	var r0 types.TxData
	if rf, ok := ret.Get(0).(func() types.TxData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TxData)
		}
	}

	return r0
}

// Copy provides a mock function with given fields:
func (_m *TxData) Copy() evmtypes.TxData {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Copy")
	}

	var r0 evmtypes.TxData
	if rf, ok := ret.Get(0).(func() evmtypes.TxData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evmtypes.TxData)
		}
	}

	return r0
}

// Cost provides a mock function with given fields:
func (_m *TxData) Cost() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Cost")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// EffectiveCost provides a mock function with given fields: baseFee
func (_m *TxData) EffectiveCost(baseFee *big.Int) *big.Int {
	ret := _m.Called(baseFee)

	if len(ret) == 0 {
		panic("no return value specified for EffectiveCost")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(baseFee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// EffectiveFee provides a mock function with given fields: baseFee
func (_m *TxData) EffectiveFee(baseFee *big.Int) *big.Int {
	ret := _m.Called(baseFee)

	if len(ret) == 0 {
		panic("no return value specified for EffectiveFee")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(baseFee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// EffectiveGasPrice provides a mock function with given fields: baseFee
func (_m *TxData) EffectiveGasPrice(baseFee *big.Int) *big.Int {
	ret := _m.Called(baseFee)

	if len(ret) == 0 {
		panic("no return value specified for EffectiveGasPrice")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(baseFee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// Fee provides a mock function with given fields:
func (_m *TxData) Fee() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fee")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetAccessList provides a mock function with given fields:
func (_m *TxData) GetAccessList() types.AccessList {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccessList")
	}

	var r0 types.AccessList
	if rf, ok := ret.Get(0).(func() types.AccessList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccessList)
		}
	}

	return r0
}

// GetChainID provides a mock function with given fields:
func (_m *TxData) GetChainID() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetChainID")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetData provides a mock function with given fields:
func (_m *TxData) GetData() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GetGas provides a mock function with given fields:
func (_m *TxData) GetGas() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGas")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetGasFeeCap provides a mock function with given fields:
func (_m *TxData) GetGasFeeCap() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGasFeeCap")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetGasPrice provides a mock function with given fields:
func (_m *TxData) GetGasPrice() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrice")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetGasTipCap provides a mock function with given fields:
func (_m *TxData) GetGasTipCap() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGasTipCap")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetNonce provides a mock function with given fields:
func (_m *TxData) GetNonce() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNonce")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetRawSignatureValues provides a mock function with given fields:
func (_m *TxData) GetRawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRawSignatureValues")
	}

	var r0 *big.Int
	var r1 *big.Int
	var r2 *big.Int
	if rf, ok := ret.Get(0).(func() (*big.Int, *big.Int, *big.Int)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func() *big.Int); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func() *big.Int); ok {
		r2 = rf()
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	return r0, r1, r2
}

// GetTo provides a mock function with given fields:
func (_m *TxData) GetTo() *common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTo")
	}

	var r0 *common.Address
	if rf, ok := ret.Get(0).(func() *common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *TxData) GetValue() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetValue")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// SetSignatureValues provides a mock function with given fields: chainID, v, r, s
func (_m *TxData) SetSignatureValues(chainID *big.Int, v *big.Int, r *big.Int, s *big.Int) {
	_m.Called(chainID, v, r, s)
}

// TxType provides a mock function with given fields:
func (_m *TxData) TxType() byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxType")
	}

	var r0 byte
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *TxData) Validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTxData creates a new instance of TxData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxData(t interface {
	mock.TestingT
	Cleanup(func())
},
) *TxData {
	mock := &TxData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
