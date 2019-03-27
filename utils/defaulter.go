package utils

import (
	"reflect"

	"github.com/hbagdi/defaulter"
	"github.com/pkg/errors"
)

// Defaulter registers types and fills in struct fields with
// default values.
type Defaulter struct {
	r map[string]interface{}
}

func (d *Defaulter) once() {
	if d.r == nil {
		d.r = make(map[string]interface{})
	}
}

// Register registers a type and it's default value.
// The default value is passed in and the type is inferred from the
// default value.
func (d *Defaulter) Register(def interface{}) error {
	d.once()
	v := reflect.ValueOf(def)
	if !v.IsValid() {
		return errors.New("invalid value")
	}
	v = reflect.Indirect(v)
	d.r[v.Type().String()] = def
	return nil
}

// Set fills in default values in a struct of a registered type.
func (d *Defaulter) Set(arg interface{}) error {
	d.once()
	v := reflect.ValueOf(arg)
	if !v.IsValid() {
		return errors.New("invalid value")
	}
	v = reflect.Indirect(v)
	defValue, ok := d.r[v.Type().String()]
	if !ok {
		return errors.New("type not registered: " + reflect.TypeOf(arg).String())
	}
	return defaulter.Set(arg, defValue)
}
