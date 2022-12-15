/*Interfaces for Password Checker API*/
package interfaces

import "github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"

/*Password Checker Add Match is used to add a string for NoMatch if Password Checker
Do not follow its Constraints */
type PWChecker interface {
	AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string
}
