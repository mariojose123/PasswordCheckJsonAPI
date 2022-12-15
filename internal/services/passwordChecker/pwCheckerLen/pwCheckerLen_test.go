package pwCheckerLen

import (
	"log"
	"reflect"
	"testing"

	"github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"
)

func TestCheckMinLen(t *testing.T) {
	mapWithString4 := map[string]int{
		"MinSize": 4,
	}
	mapWithString0 := map[string]int{
		"MinSize": 0,
	}
	mapWithoutString := map[string]int{}
	tests := []struct {
		name          string
		jsonStructMap structJson.PSReceiveStructure
		noMatch       []string
		Checker       PWCheckerLen
		want          []string
	}{
		{"Min Len inCorrect and Parameter inside map", structJson.PSReceiveStructure{PW: "111", Rules: mapWithString4}, []string{}, PWCheckerLen{"MinSize"}, []string{"MinSize"}},
		{"MinLenCorrect and Parameter is not inside  map", structJson.PSReceiveStructure{PW: "11111111", Rules: mapWithoutString}, []string{}, PWCheckerLen{"MinSize"}, []string{}},
		{"MinLen Correct 4 and Parameter inside map", structJson.PSReceiveStructure{PW: "11111111", Rules: mapWithString4}, []string{}, PWCheckerLen{"MinSize"}, []string{}},
		{"MinLen Correct 0 and Parameter inside map", structJson.PSReceiveStructure{PW: "", Rules: mapWithString0}, []string{}, PWCheckerLen{"MinSize"}, []string{}},
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
