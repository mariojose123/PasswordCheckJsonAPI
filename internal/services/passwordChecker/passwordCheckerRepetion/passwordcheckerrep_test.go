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
		RepeatedString: 1,
	}
	mapWithStringFalse := map[string]int{
		RepeatedString: 0,
	}
	mapWithoutString := map[string]int{}
	tests := []struct {
		name          string
		jsonStructMap structJson.PSReceiveStructure
		noMatch       []string
		Checker       PWCheckerRepetion
		want          []string
	}{
		{"Only repeated digits", structJson.PSReceiveStructure{"111", mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{RepeatedString}},
		{"No Repeated Digits", structJson.PSReceiveStructure{"1234567", mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{}},
		{"Repeated Digits beetween not repeated Digits", structJson.PSReceiveStructure{"34111567", mapWithStringTrue}, []string{}, PWCheckerRepetion{RepeatedString}, []string{RepeatedString}},
		{"Dont Check Repeated Digits", structJson.PSReceiveStructure{"22111222", mapWithStringFalse}, []string{}, PWCheckerRepetion{RepeatedString}, []string{}},
		{"Json Map empty", structJson.PSReceiveStructure{"22111222", mapWithoutString}, []string{}, PWCheckerRepetion{RepeatedString}, []string{}},
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
