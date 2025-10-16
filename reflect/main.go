package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct {
		return errors.New("hanya menerima struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tags := field.Tag

		if tags.Get("required") == "true" {
			if isEmpty(value) {
				return fmt.Errorf("field %s wajib diisi", field.Name)
			}
		}

		if value.Kind() == reflect.Int {
			if minTag := tags.Get("min"); minTag != "" {
				min := parseInt(minTag)
				if int(value.Int()) < min {
					return fmt.Errorf("field %s harus >= %d", field.Name, min)
				}
			}
			if maxTag := tags.Get("max"); maxTag != "" {
				max := parseInt(maxTag)
				if int(value.Int()) > max {
					return fmt.Errorf("field %s harus <= %d", field.Name, max)
				}
			}
		}

		if value.Kind() == reflect.String {
			if minLenTag := tags.Get("minLen"); minLenTag != "" {
				minLen := parseInt(minLenTag)
				if len(value.String()) < minLen {
					return fmt.Errorf("field %s minimal %d karakter", field.Name, minLen)
				}
			}
			if maxLenTag := tags.Get("maxLen"); maxLenTag != "" {
				maxLen := parseInt(maxLenTag)
				if len(value.String()) > maxLen {
					return fmt.Errorf("field %s maksimal %d karakter", field.Name, maxLen)
				}
			}

			// Handle email
			if tags.Get("email") == "true" {
				if !isValidEmail(value.String()) {
					return fmt.Errorf("field %s bukan email valid", field.Name)
				}
			}
		}
	}

	return nil
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int:
		return v.Int() == 0
	default:
		return v.IsZero()
	}
}

func parseInt(s string) int {
	var val int
	fmt.Sscan(s, &val)
	return val
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

type CriminalRecord struct {
	Name      string `required:"true" minLen:"3" maxLen:"30"`
	Email     string `required:"true" email:"true"`
	ThreatLvl int    `required:"true" min:"1" max:"10"`
}

func main() {
	// Data sample valid
	record := CriminalRecord{
		Name:      "Loki Laufeyson",
		Email:     "loki@asgard.org",
		ThreatLvl: 9,
	}

	// Data sample invalid (contoh uji)
	recordInvalid := CriminalRecord{
		Name:      "Lo",
		Email:     "loki-asgard",
		ThreatLvl: 11,
	}

	fmt.Println("=== VALIDASI RECORD VALID ===")
	if err := ValidateStruct(record); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validasi berhasil ✅")
	}

	fmt.Println("\n=== VALIDASI RECORD INVALID ===")
	if err := ValidateStruct(recordInvalid); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validasi berhasil ✅")
	}
}
