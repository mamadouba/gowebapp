package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Struct(obj interface{}) (bool, []string) {
	var errs []string
	errs = validateStruct(errs, obj)
	return len(errs) == 0, errs
}

func Var(obj interface{}, tag string) (bool, []string) {
	var errs []string
	errs = validateVar(errs, obj, tag)
	return len(errs) == 0, errs
}

func validateVar(errs []string, obj interface{}, tag string) []string {
	return errs
}

func validateStruct(errs []string, obj interface{}) []string {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" || tag == "-" {
			continue
		}

		fieldVal := val.Field(i)
		fieldValue := fieldVal.Interface()
		zero := reflect.Zero(field.Type).Interface()

		if field.Type.Kind() == reflect.Struct {
			errs = validateStruct(errs, fieldValue)
		}
		errs = validateField(errs, zero, field, fieldVal, fieldValue)
	}
	return errs
}

func validateField(errs []string, zero interface{}, field reflect.StructField, fieldVal reflect.Value, fieldValue interface{}) []string {
	if fieldVal.Kind() == reflect.Slice {
		for i := 0; i < fieldVal.Len(); i++ {
			sliceVal := fieldVal.Index(i)
			if sliceVal.Kind() == reflect.Ptr {
				sliceVal = sliceVal.Elem()
			}
			sliceValue := sliceVal.Interface()
			zero := reflect.Zero(sliceVal.Type()).Interface()
			if sliceVal.Kind() == reflect.Struct {
				errs = validateStruct(errs, sliceValue)
			} else {
				errs = validateField(errs, zero, field, sliceVal, sliceValue)
			}
		}
	}
VALIDATE_ARGS:
	for _, arg := range strings.Split(field.Tag.Get("validate"), ";") {
		if len(arg) == 0 {
			continue
		}
		switch {
		case arg == "omitempty":
			if reflect.DeepEqual(zero, fieldValue) {
				break VALIDATE_ARGS
			}
		case arg == "required":
			v := reflect.ValueOf(fieldValue)
			if v.Kind() == reflect.Slice {
				if v.Len() == 0 {
					errs = append(errs, fmt.Sprintf("%s is required", field.Name))
				}
				continue
			}
			if reflect.DeepEqual(zero, fieldValue) {
				errs = append(errs, fmt.Sprintf("%s is required", field.Name))
				break VALIDATE_ARGS
			}
		case strings.HasPrefix(arg, "min="):
			min, _ := strconv.Atoi(arg[4:])
			if val, ok := fieldValue.(int); ok && val < min {
				errs = append(errs, fmt.Sprintf("%s must be greater than %d", field.Name, min))
				break VALIDATE_ARGS
			}
		case strings.HasPrefix(arg, "max="):
			max, _ := strconv.Atoi(arg[4:])
			if val, ok := fieldValue.(int); ok && val > max {
				errs = append(errs, fmt.Sprintf("%s must be lesser than %d", field.Name, max))
				break VALIDATE_ARGS
			}
		case strings.HasPrefix(arg, "length="):
			length, _ := strconv.Atoi(arg[7:])
			if val, ok := fieldValue.(string); ok && len(val) < length {
				errs = append(errs, fmt.Sprintf("%s must be %d characters at least", field.Name, length))
				break VALIDATE_ARGS
			}
		case strings.HasPrefix(arg, "regexp="):
			re := regexp.MustCompile(arg[7:])
			if !re.MatchString(fieldValue.(string)) {
				errs = append(errs, fmt.Sprintf("%s does not match to expected format", field.Name))
				break VALIDATE_ARGS
			}
		case arg == "email":
			re := regexp.MustCompile(`^[a-z0-9._\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
			if !re.MatchString(fieldValue.(string)) {
				errs = append(errs, fmt.Sprintf("%s is not valid", field.Name))
				break VALIDATE_ARGS
			}
		default:
		}
	}
	return errs
}
