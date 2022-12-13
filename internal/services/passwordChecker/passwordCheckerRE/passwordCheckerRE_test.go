package passwordCheckerRE

import (
	"log"
	"passwordcheck/internal/structJson"
	"reflect"
	"regexp"
	"testing"
)

func TestCheckMinMatchRegularExpression(t *testing.T) {

	REminMatchString := "ReMin"

	RegexSpecial, _ := regexp.Compile("[" + regexp.QuoteMeta(`!@#$%^&*()-+\/{}[]`) + "]")

	mapWithString16 := map[string]int{
		REminMatchString: 16,
	}
	mapWithString0 := map[string]int{
		REminMatchString: 0,
	}
	mapEmpty := map[string]int{}

	CheckerSpecial := PWCheckerMinRE{RegexSpecial, REminMatchString}

	tests := []struct {
		name          string
		jsonStructMap structJson.PSReceiveStructure
		noMatch       []string
		Checker       PWCheckerMinRE
		want          []string
	}{
		{"All special Characters must have Min of 16 char", structJson.PSReceiveStructure{PW: `!@#$%^&*()-+\/{}[]`, Rules: mapWithString16}, []string{}, CheckerSpecial, []string{}},
		{"0 match then any string is allowed for this checker", structJson.PSReceiveStructure{PW: "", Rules: mapWithString0}, []string{}, CheckerSpecial, []string{}},
		{"Repeated Digits beetween not repeated Digits", structJson.PSReceiveStructure{PW: `!@#$%^&*`, Rules: mapWithString16}, []string{}, CheckerSpecial, []string{REminMatchString}},
		{"Repeated Digits beetween not repeated Digits", structJson.PSReceiveStructure{PW: `!@#$%^&*`, Rules: mapEmpty}, []string{}, CheckerSpecial, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Print(tt.jsonStructMap)
			ret := tt.Checker.AddMatch(tt.jsonStructMap, tt.noMatch)
			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("got %s, want %s", ret, tt.want)
			}
		})
	}
}
