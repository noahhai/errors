package wrongo_test

import (
	"fmt"
	"testing"
	"wrongo"
)

func TestFormat(t *testing.T) {
	err := wrongo.New("error1").Add("error2").Add("error3")
	fmt.Println(err)
}
