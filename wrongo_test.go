package wrongo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wrongo"
)

func TestAdd(t *testing.T) {
	err := wrongo.New("error1").Add("error2").Add("error3")
	assert.Equal(t,"error3\nerror2\nerror1", err.Error())
}

func TestOr(t *testing.T) {
	var err1 error
	err2 := wrongo.New("error2")
	err3 := wrongo.Or(err1, err2)
	assert.Equal(t, "error2", err3.Error())
}
