package services

import (
	"context"
	"passwordcheck/interfaces"
	"passwordcheck/internal/services/passwordChecker/passwordCheckerRE"
	passwordcheckerrep "passwordcheck/internal/services/passwordChecker/passwordCheckerRepetion"
	"passwordcheck/internal/services/passwordChecker/pwCheckerLen"
	"passwordcheck/internal/structJson"
)

/*
	Service that provide the fuction that check the PW for bussness logic,it has a list of constrains

that it must follows ,every constrain has a function that add a string in a string array if Constrain IS NOT followed

	Constrains parameters are defined on main.go
*/
type PWService struct {
	ConstrainsPWRE []interfaces.PWChecker
}

/*Init all constrains for PW Chekcer , Constrains parameters are defined on main.go*/
func NewPWService(UpperCaseString string, LowerCaseString string, DigitsString string, SpecialCharacters string,
	MinLenString string, RepeatedString string, UpperCaseRE string, LowerCaseRE string, DigitsRE string, SPRE string) (*PWService, error) {

	var checkers []interfaces.PWChecker

	pwCheckerUp, err := passwordCheckerRE.NewPWCheckRE(UpperCaseString, UpperCaseRE)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerLower, err := passwordCheckerRE.NewPWCheckRE(LowerCaseString, LowerCaseRE)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerDigits, err := passwordCheckerRE.NewPWCheckRE(DigitsString, DigitsRE)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerSpecial, err := passwordCheckerRE.NewPWCheckRE(SpecialCharacters, SPRE)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerLen := pwCheckerLen.NewPWCheck(MinLenString)
	pwCheckerRepeated := passwordcheckerrep.NewPWCheck(RepeatedString)

	checkers = append(checkers, pwCheckerUp, pwCheckerLower,
		pwCheckerDigits, pwCheckerSpecial, pwCheckerLen, pwCheckerRepeated)
	return &PWService{checkers}, nil
}

/*
Receives context and jsonStructure with map for every parameter for the Checkers and execute all the checkers method that add a Match if
Checker is not followed ,in the end check if noMatch is empty to see if no Checker added a Match because constrain was not followed
*/
func (serv *PWService) CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string) {

	var noMatch []string
	for _, checkerRE := range serv.ConstrainsPWRE {
		noMatch = checkerRE.AddMatch(jsonStructure, noMatch)
	}
	if len(noMatch) > 0 {
		return false, noMatch
	}
	return true, noMatch
}
