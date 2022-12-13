package services

import (
	"context"
	"passwordcheck/interfaces"
	"passwordcheck/internal/structJson"
	mock_interfaces "passwordcheck/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestServiceCheckPW(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	noMatchWant := []string{"MinDigit"}
	isCheckedWanted := false
	var noMatchTest []string = make([]string, 0)
	mockPWChecker := mock_interfaces.NewMockPWChecker(control)

	mockPWChecker.EXPECT().AddMatch(structJson.PSReceiveStructure{PW: "TesteSenhaForte!123&", Rules: map[string]int{
		"MinDigit": 4,
	}}, noMatchTest).Return([]string{"MinDigit"})

	testservice := PWService{[]interfaces.PWChecker{mockPWChecker}}
	IsCheckedReturn, noMatchReturn := testservice.CheckPW(context.Background(), structJson.PSReceiveStructure{PW: "TesteSenhaForte!123&", Rules: map[string]int{
		"MinDigit": 4,
	}})
	if !(reflect.DeepEqual(noMatchWant, noMatchReturn) && isCheckedWanted == IsCheckedReturn) {
		t.Errorf("got %s, want %s", noMatchReturn, noMatchWant)
		t.Errorf("got %t, want %t", IsCheckedReturn, isCheckedWanted)
	}
}
