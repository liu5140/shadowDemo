package bizerr

import "errors"

type GlobalError struct {
	error
	Arg []interface{}
}

func GenErr(errMsg string, arg ...interface{}) (err error) {
	err = GlobalError{
		errors.New(errMsg),
		arg,
	}
	Log.Error(err)
	return err
}
