package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaulter(t *testing.T) {
	assert := assert.New(t)

	var d Defaulter

	assert.NotNil(d.Register(nil))
	assert.NotNil(d.Set(nil))

	type Foo struct {
		A string
		B []string
	}
	defaultFoo := &Foo{
		A: "defaultA",
		B: []string{"default1"},
	}
	assert.Nil(d.Register(defaultFoo))

	// sets a default
	var arg Foo
	assert.Nil(d.Set(&arg))
	assert.Equal("defaultA", arg.A)
	assert.Equal([]string{"default1"}, arg.B)

	// doesn't set a default
	arg1 := Foo{
		A: "non-default-value",
	}
	assert.Nil(d.Set(&arg1))
	assert.Equal("non-default-value", arg1.A)

	// errors on an unregistered type
	type Bar struct {
		A string
	}
	assert.NotNil(d.Set(&Bar{}))
}
