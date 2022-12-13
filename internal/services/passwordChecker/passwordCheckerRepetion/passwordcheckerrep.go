/*package for Password Checker of repetition  */
package passwordcheckerrep

import "passwordcheck/internal/structJson"

/*A struct that Check repetition*/
type PWCheckerRepetion struct {
	nameConst string
}

/*Create Password Checker that check repetion*/
func NewPWCheck(constname string) PWCheckerRepetion {
	return PWCheckerRepetion{nameConst: constname}
}

/* Add Match if PW has repeated characters*/
func (PWChecker PWCheckerRepetion) AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string {
	PW := jsonStruct.PW
	isRepeated := checkRepetionPW(PW)
	if isRepeated {
		return append(noMatch, PWChecker.nameConst)
	}
	return noMatch
}

/*Check if a string has repeated chars*/
func checkRepetionPW(PW string) bool {
	for indexChar := 0; indexChar < len(PW)-1; indexChar++ {
		if PW[indexChar] == PW[indexChar+1] {
			return true
		}
	}
	return false
}
