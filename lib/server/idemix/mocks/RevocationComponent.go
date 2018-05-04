/*
Copyright IBM Corp. 2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0

package mocks

import idemix "github.com/hyperledger/fabric-ca/lib/server/idemix"
import mock "github.com/stretchr/testify/mock"

// RevocationComponent is an autogenerated mock type for the RevocationComponent type
type RevocationComponent struct {
	mock.Mock
}

// GetNewRevocationHandle provides a mock function with given fields:
func (_m *RevocationComponent) GetNewRevocationHandle() (*idemix.RevocationHandle, error) {
	ret := _m.Called()

	var r0 *idemix.RevocationHandle
	if rf, ok := ret.Get(0).(func() *idemix.RevocationHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.RevocationHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}