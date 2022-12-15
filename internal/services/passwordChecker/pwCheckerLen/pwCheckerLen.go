/* Password Checker of minimun len x*/
package pwCheckerLen

import "github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"

/*A struct that check length of PW */
type PWCheckerLen struct {
	nameConst string
}

func NewPWCheck(constname string) PWCheckerLen {
	return PWCheckerLen{nameConst: constname}
}

/*Add Match if  PW has reached minimum length  */
func (PWChecker PWCheckerLen) AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string {
	isMin := PWChecker.isMinLen(jsonStruct)
	if !isMin {
		return append(noMatch, PWChecker.nameConst)
	}
	return noMatch
}

/*Check if PW has Minimun Lenght */
func (checker PWCheckerLen) isMinLen(jsonStructure structJson.PSReceiveStructure) bool {
	PW := jsonStructure.PW
	minLen, ok := jsonStructure.Rules[checker.nameConst]
	if ok {
		return (len(PW) >= minLen)
	} else {
		return true
	}
}
