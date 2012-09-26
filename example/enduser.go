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

// example of call site of a function/method returning Recoverable errors.
//
package main

import (
	"efface"
	"errors"
	"fmt"
)

func main() {
	v := newValue()

	v.Set("Salaam!")

	// v is dirty and we expect an error here
	// example of using Recoverable on the enduser / call site
	//
	newvalue := "Salaam again!"
	e = v.Set(errors.New(newvalue))
	switch e.(type) {
	case efface.Recoverable:
		re, _ := e.(efface.Recoverable).Recover(newvalue)
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
