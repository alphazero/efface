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

// Efface embodies a recoverable error pattern.  
//
// Recoverable errors are error types that further support the efface.Recoverable
// interface.  This interface provides a Recover() function with generalized call in/out
// signature semantics. 
//
// usage model:
//
// Packages that wish to return Recoverables as errors provide function(s) of type RecoverFn
// and return efface.NewRecoverable(error, RecoverFn) as the error result.
//
// End users of such packages can test if efface.IsRecoverable(error) and if yes, can
// attempt e.(efface.Recoverable).Recover(recovery, method, args, go, here) 
//
// A e.Recovery() will return an error if the recovery process itself faulted.
//
// A e.Recovery() will return as generalized marshalled out params a []interface{}
// result set of the package provided RecoverFn (whatever they may be).
//
package efface 

import (
	"fmt"
)

type RecoverFn func(in_args ...interface{}) (error, []interface{})
type Recoverable interface {
	error
	Recover(in_args ...interface{}) (error, []interface{})
}

func IsRecoverable(ref interface{}) bool {
	switch ref.(type) {
	case Recoverable:
		return true
	}
	return false
}

type reocverableError struct {
	cause       error
	recoverFunc RecoverFn
}

func NewRecoverableError(cause error, fn RecoverFn) Recoverable {
	return reocverableError{cause, fn}
}

func (e reocverableError) Error() string {
	return e.cause.Error()
}

func (e reocverableError) Recover(in ...interface{}) (re error, rout []interface{}) {
	defer func() {
		p := recover()
		if p != nil {
			switch p.(type) {
			case error:
				re = p.(error)
			default:
				re = fmt.Errorf("error recovering %s - %s", e.Error(), p)
			}
			return
		}
		return
	}()
	return e.recoverFunc(in...)
}
