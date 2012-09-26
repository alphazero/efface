//   Copyright 2012 Joubin Houshyar
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Example of a type returning efface.Recoverable in one of its methods.
// Here type value returns a Recoverable on Set() calls if the value is 
// already set. The recovery provided uses a force set.  
package main

import (
	"efface"
	"fmt"
)

type Value interface {
	Set(v interface{}) error
	Get() (v interface{})
}

type value struct {
	v interface{}
}

func newValue() Value {
	return &value{}
}

func (self *value) Set(v interface{}) error {
	if self.v != nil {
		// error to set on dirty value but recoverable
		return self.newDirtySetError() // use efface error type
	}
	self.v = v
	return nil
}

var set_error error = fmt.Errorf("dirty value - use SetF to force set of dirty value")

func (self *value) newDirtySetError() error {
	fn := func(in ...interface{}) (error, []interface{}) {

		// try recover
		e := self.forceSet(in[0])

		out := make([]interface{}, 1)
		out[0] = e

		return nil, out
	}
	return efface.NewRecoverableError(set_error, fn)
}

// Force Set of dirty value
func (self *value) forceSet(v interface{}) error {
	self.v = nil
	e := self.Set(v)
	return e
}

func (self *value) Get() (v interface{}) {
	return self.v
}
