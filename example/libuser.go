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

// ======================================================================
//  * usage * writing functions that return Recoverble
// ======================================================================

var err_dirtyset error = fmt.Errorf("dirty value - use forceSet(newvalue, true) to force set of dirty value")

// a value type api with explicit set and get
type Value interface {
	Set(v interface{}) error
	Get() (v interface{})
}

// supports Value interface
type value struct {
	v interface{}
}

func newValue() Value {
	return &value{}
}

// Set method doesn't allow dirty sets,
// and returns a recoverable error in that case, with signature:
//
// Recoverable()@Set => type func(v interface{}, f bool) (e error)
// - v: value to set
// - f: force set on dirty set
// - e: recovery attempt result.
func (self *value) Set(v interface{}) error {
	if self.v != nil {
		return self.newDirtySetError()
	}
	self.v = v
	return nil
}

func (self *value) newDirtySetError() error {
	// define recover function - here was just delegate to a method call.
	fn := func(in ...interface{}) interface{} {
		out := self.forceSet(in[0], in[1].(bool))
		return out
	}
	// return recoverable error
	return efface.NewRecoverableError(err_dirtyset, fn)
}

// Force Set of dirty value
func (self *value) forceSet(v interface{}, force bool) error {
	if self.v != nil && !force {
		return self.newDirtySetError()
	}
	self.v = nil
	e := self.Set(v)
	return e
}

func (self *value) Get() (v interface{}) {
	return self.v
}

// ======================================================================
//  * usage * using functions returning Recoverable
// ======================================================================

func main() {
	v := newValue()

	// set v 
	v.Set("Salaam!")

	// again calling set should raise an error.
	// v is dirty and we expect an error here
	// example of using Recoverable on the enduser / call site
	newvalue := "Salaam again!"
	e := v.Set(newvalue)
	fmt.Printf("* debug * main * v.Set() => '%s'\n", e)
	switch e.(type) {
	case efface.Recoverable:
		re, out := e.(efface.Recoverable).Recover(newvalue, true)
		if re != nil {
			fmt.Printf("* main * error * error on recover attempt * %s\n", re)
			return
		}
	case error:
		fmt.Printf("* main * unrecoverable error * %s\n", e)
		return
	default: // no error
	}
}
