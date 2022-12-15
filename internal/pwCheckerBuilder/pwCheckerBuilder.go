/*package with builders */
package pwCheckerBuilder

import (
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/handler"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/services"
)

/*
Build Handler (handler.HandlerPWCheck) for API that checks the password
Input is the string of every Json according to specification
for example for the specification DigitsString="minDigit"
this is to check JSON and to be able to change JSON string if necessary
*/
func NewBuilderpwCheckHandler(UpperCaseString string, LowerCaseString string, DigitsString string, SpecialCharacters string,
	MinLenString string, RepeatedString string) (*handler.HandlerPWCheck, error) {
	PWService, err := services.NewPWService(
		UpperCaseString,
		LowerCaseString,
		DigitsString,
		SpecialCharacters,
		MinLenString,
		RepeatedString,
	)
	if err != nil {
		return &handler.HandlerPWCheck{}, err
	}
	return handler.NewHandler(PWService), nil

}
