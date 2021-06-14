package shortcodeutil

import (
	"testing"
)

func TestGenerateShortcode(t *testing.T) {
    t.Run("success", func(t *testing.T){
    	length := 6
		shortcode := GenerateShortcode(length)

		if len(shortcode) != length {
			t.Fatalf(`got length %v, want length of %v`, len(shortcode), length)
		}
    })
}

func TestValidateShortcode(t *testing.T) {
	t.Run("valid shortcode", func(t *testing.T){
    	shortcode := "asd123"
		isValid := ValidateShortcode(shortcode)

		if isValid == false {
			t.Fatalf(`got %v, want %v for %v`, isValid, true, shortcode)
		}
    })

    t.Run("too long shortcode", func(t *testing.T){
    	shortcode := "1234567"
		isValid := ValidateShortcode(shortcode)

		if isValid == true {
			t.Fatalf(`got %v, want %v for %v`, isValid, false, shortcode)
		}
    })

    t.Run("too short shortcode", func(t *testing.T){
    	shortcode := "12345"
		isValid := ValidateShortcode(shortcode)

		if isValid == true {
			t.Fatalf(`got %v, want %v for %v`, isValid, false, shortcode)
		}
    })

    t.Run("shortcode not alphanumeric", func(t *testing.T){
    	shortcode := "as12@$"
		isValid := ValidateShortcode(shortcode)

		if isValid == true {
			t.Fatalf(`got %v, want %v for %v`, isValid, false, shortcode)
		}
    })
}

