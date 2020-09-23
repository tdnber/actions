package hello_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
"actions/hello"
)

func Test_Result(t *testing.T) {

	assert.Equal(t, 1, hello.GetResult())
}
