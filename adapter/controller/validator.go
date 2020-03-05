package controller

import (
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"gopkg.in/validator.v2"
	"net/mail"
	"reflect"
	"strconv"
)

var (
	ErrRequired = validator.TextErr{Err: errors.New("required")}
	ErrUint     = validator.TextErr{Err: errors.New("invalid uint")}
	ErrEmail    = validator.TextErr{Err: errors.New("invalid email")}
	ErrUniq     = validator.TextErr{Err: errors.New("unique email")}
)

type ValidatorSetting struct {
	ArgName      string
	ValidateTags string
}

func initValidator() {
	validator.SetValidationFunc("required", requiredValidator)
	validator.SetValidationFunc("uint", uintValidator)
	validator.SetValidationFunc("email", emailValidator)
}

func Validate(params map[string]interface{}, settings []*ValidatorSetting) map[string]error {
	initValidator()

	errs := map[string]error{}
	for _, setting := range settings {
		err := validator.Valid(params[setting.ArgName], setting.ValidateTags)
		if err != nil {
			arr := err.(validator.ErrorArray)
			errs[setting.ArgName] = arr[0]
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func ValidateParams(params map[string]string, settings []*ValidatorSetting) map[string]error {
	p := map[string]interface{}{}
	for k, v := range params {
		p[k] = v
	}
	return Validate(p, settings)
}

func ValidateBody(body string, settings []*ValidatorSetting) map[string]error {
	var b map[string]interface{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return map[string]error{}
	}
	return Validate(b, settings)
}

func requiredValidator(v interface{}, param string) error {
	if v == nil {
		return ErrRequired
	}

	st := reflect.ValueOf(v)

	if st.String() == "" {
		return ErrRequired
	}

	return nil
}

func uintValidator(v interface{}, param string) error {
	if v == nil {
		return nil
	}

	st := reflect.ValueOf(v)

	if st.String() == "" {
		return nil
	}

	var n int

	switch st.Kind() {
	case reflect.String:
		n64, err := strconv.ParseInt(st.String(), 10, 64)
		if err != nil {
			glog.Warningf("%s:%s", param, err.Error())
			return validator.ErrUnsupported
		}
		n = int(n64)
	case reflect.Int:
		n = v.(int)
	case reflect.Float64:
		n = int(v.(float64))
	default:
		glog.Warningf("%s:%s", param, st.Kind())
		return validator.ErrUnsupported
	}

	if n < 0 {
		return ErrUint
	}

	return nil
}

func emailValidator(v interface{}, param string) error {
	if v == nil {
		return nil
	}

	st := reflect.ValueOf(v)

	if st.String() == "" {
		return nil
	}

	_, err := mail.ParseAddress(st.String())
	if err != nil {
		glog.Warningf("failed to parse %s: %s", st.String(), err.Error())
		return ErrEmail
	}

	return nil
}
