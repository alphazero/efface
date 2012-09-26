# efface

The `efface` package defines the semantics for creating and using errors that can be optionally recovered.
It is intended for use by either library providers or app developers.

# usage

Both caller and reciever sides must import "efface".

## function/method provider

If a function or method returns Go's 'error' type, it can choose to return efface.Recoverable instead.  To do this,
it must provide a function satisfying efface.RecoeryFn and return a new effece.Recoverable via efface.NewRecoverableError() function.  This returned function has generalized nary (e.g. ...) input arg signature, so the most important
requirement for method/func developers is to document the expected arguments for the recovery functions.

## function/method caller

The end users of such methods and functions can test if the 'error' returned is a Recoverable error via efface.IsRecoverable().  If yes, they can call Recoverable.Recover(args...) to attempt to recover from the error.  The returned
results are in form of a generic []interface{}, which is a bit tedious but unavoidable.


