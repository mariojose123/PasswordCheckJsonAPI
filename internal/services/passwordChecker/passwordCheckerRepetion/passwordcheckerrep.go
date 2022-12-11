package passwordcheckerrep

import "passwordcheck/internal/structJson"

/*Struct that Check repetetion*/
type PWCheckerRepetion struct {
	nameConst string
}

func NewPWCheck(constname string) PWCheckerRepetion {
	return PWCheckerRepetion{nameConst: constname}
}

/* Add Match if PW has repeated characters*/
func (PWChecker PWCheckerRepetion) AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string {
	PW := jsonStruct.PW
	isRepeated := isRepeated(PW, PWChecker.nameConst, jsonStruct)
	if isRepeated {
		return append(noMatch, PWChecker.nameConst)
	}
	return noMatch
}

/* Check if there is repeated words in the string*/
func isRepeated(PW string, repeationString string, jsonStructure structJson.PSReceiveStructure) bool {
	CheckRepeationInt, ok := jsonStructure.Rules[repeationString]
	CheckRepeation := CheckRepeationInt == 1
	if ok {
		if CheckRepeation {
			return checkRepetionPW(PW)
		}
		return false
	} else {
		return false
	}
}

func checkRepetionPW(PW string) bool {
	for indexChar := 0; indexChar < len(PW)-1; indexChar++ {
		if PW[indexChar] == PW[indexChar+1] {
			return true
		}
	}
	return false
}
