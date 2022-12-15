/*
	Service that provides the function that checks the Password, has a list of constraints list

that it must follow, every constraint has a function that adds a string in a string array if Constrain IS NOT followed (AddMatch)

	Constrains parameters are defined on main.go
*/
package services

import (
	"context"

	"github.com/mariojose123/PasswordCheckJsonAPI/interfaces"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/services/passwordChecker/passwordCheckerRE"
	passwordcheckerrep "github.com/mariojose123/PasswordCheckJsonAPI/internal/services/passwordChecker/passwordCheckerRepetion"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/services/passwordChecker/pwCheckerLen"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"
)

/* Struct tha contains Contrains that are a array with password Checker*/
type PWService struct {
	constrainsPWRE []interfaces.PWChecker
}

/*
Init all constraints for PWService according to the specifications, Constraints parameters are defined on the main. go
the function could be generic for other password checkers but it has been implemented for the specification
The function add constrains checker for password on constrainsPWRE
*/
func NewPWService(UpperCaseString string, LowerCaseString string, DigitsString string, SpecialCharacters string,
	MinLenString string, RepeatedString string) (*PWService, error) {

	var checkers []interfaces.PWChecker

	pwCheckerUp, err := passwordCheckerRE.NewPWCheckREUpperCase(UpperCaseString)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerLower, err := passwordCheckerRE.NewPWCheckRELowedCase(LowerCaseString)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerDigits, err := passwordCheckerRE.NewPWCheckREDigits(DigitsString)
	if err != nil {
		return &PWService{}, err
	}
	pwCheckerSpecial, err := passwordCheckerRE.NewPWCheckRESpecialchar(SpecialCharacters)
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
Given that the PWService has a array of Password Checkers add string of
Password Checker to noMatch string if noMatch string len is bigger
than 0 a Constrain was not followed
return True if all password matched are followed and
false Otherwise and NoMatch string array with every Constrain for Password that was not followed
*/
func (serv *PWService) CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string) {

	var noMatch []string = make([]string, 0)
	/*Check every Password Checker of service according to specification constrains and add Match on noMatch
	  to create JSON noMatch Array if constrain is not followed*/
	for _, constrain := range serv.constrainsPWRE {
		noMatch = constrain.AddMatch(jsonStructure, noMatch)
	}
	/*Check if len of noMatch array acording to specification is empty*/
	if len(noMatch) > 0 {
		return false, noMatch
	}
	return true, noMatch
}
