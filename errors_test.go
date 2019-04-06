package errors

import (
	sysErrors "errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("error1")
	assert.Equal(t, "error1", err.Error())
}

func TestNewF(t *testing.T) {
	err := NewF("error%d = %s", 1, "failure")
	assert.Equal(t, "error1 = failure", err.Error())
}
func TestAdd(t *testing.T) {
	err := New("error1").Add("error2").Add("error3")
	assert.Equal(t, "error3\nerror2\nerror1", err.Error())
}

func TestOr(t *testing.T) {
	var err1 error
	err2 := sysErrors.New("error2")
	err3 := Or(err1, err2)
	assert.Equal(t, "error2", err3.Error())
}

func TestFrom(t *testing.T) {
	f := func() error { return sysErrors.New("error during 'f()'") }
	f2 := func() error { return From(f()).Add("error during 'f2()'") }
	e := f2()
	assert.Equal(t, "error during 'f2()'\nerror during 'f()'", e.Error())
}

func TestFromTuple(t *testing.T) {
	f := func() (int, error) { return 5, sysErrors.New("error during 'f()'") }
	f2 := func() (interface{}, error) { return GetFromTupleAdd("error during 'f2()'")(f()) }
	v, e := f2()
	assert.Equal(t, "error during 'f2()'\nerror during 'f()'", e.Error())
	assert.Equal(t, 5, v.(int))
}

func TestExample(t *testing.T) {
	mockApiFunc := func() error {
		return sysErrors.New("unexpected column found")
	}
	modelFunc := func() *Error {
		return From(mockApiFunc()).Add("failed to upsert value")
	}
	businessFunc := func() *Error {
		return modelFunc().AddF("failed to perform daily sync of customer '%d'", 35)
	}

	// consume
	handlerFunc := func() {
		if err := businessFunc(); err != nil {
			fmt.Println("Error handling request")
			fmt.Println("Error: ")
			fmt.Println(err)
			fmt.Println()
			fmt.Println("Cause:")
			fmt.Println(err.Cause())
			fmt.Println()
			fmt.Println("Symptom:")
			fmt.Println(err.Symptom())
		}
	}
	handlerFunc()

	// return as built-in Error, e.g. to AWS API Gateway
	extHandlerFunc := func() error {
		return businessFunc()
	}
	_ = extHandlerFunc()
}