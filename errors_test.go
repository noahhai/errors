package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
	sysErrors "errors"
)

func TestNew(t *testing.T) {
	err := New("error1")
	assert.Equal(t,"error1", err.Error())
}

func TestNewF(t *testing.T) {
	err := NewF("error%d = %s", 1, "failure")
	assert.Equal(t,"error1 = failure", err.Error())
}
func TestAdd(t *testing.T) {
	err := New("error1").Add("error2").Add("error3")
	assert.Equal(t,"error3\nerror2\nerror1", err.Error())
}

func TestOr(t *testing.T) {
	var err1 error
	err2 := sysErrors.New("error2")
	err3 := Or(err1, err2)
	assert.Equal(t, "error2", err3.Error())
}

func TestFrom(t *testing.T) {
	f := func()error {return sysErrors.New("error during 'f()'")}
	f2 := func()error {return From(f()).Add("error during 'f2()'")}
	e := f2()
	assert.Equal(t, "error during 'f2()'\nerror during 'f()'", e.Error())
}

func TestFromTuple(t *testing.T) {
	f := func()error {return sysErrors.New("error during 'f()'")}
	f2 := func()error {return From(f()).Add("error during 'f2()'")}
	e := f2()
	assert.Equal(t, "error during 'f2()'\nerror during 'f()'", e.Error())
}
