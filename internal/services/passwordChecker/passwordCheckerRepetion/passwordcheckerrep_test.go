package passwordcheckerrep

import (
	"log"
	"passwordcheck/internal/structJson"
	"reflect"
	"testing"
)

func TestCheckRepeated(t *testing.T) {
	RepeatedString := "noRepeted"
	mapWithStringTrue := map[string]int{
		RepeatedString: 0,
	}
	tests := []struct {
		name          string
		jsonStructMap structJson.PSReceiveStructure
		noMatch       []string
		Checker       PWCheckerRepetion
		want          []string
	}{
		{"Only repeated digits", structJson.PSReceiveStructure{PW: "111", Rules: mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{RepeatedString}},
		{"No Repeated Digits", structJson.PSReceiveStructure{PW: "1234567", Rules: mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{}},
		{"Repeated Digits beetween not repeated Digits", structJson.PSReceiveStructure{PW: "34111567", Rules: mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{RepeatedString}},
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
