package passwordCheckerRE

import (
	"passwordcheck/internal/structJson"
	"regexp"
)

/*Struct that checks repetion of regular Expression in the PW */
type PWCheckerMinRE struct {
	regexCompiled *regexp.Regexp
	nameConst     string
}

func NewPWCheckRE(name string, regextext string) (PWCheckerMinRE, error) {
	regexCompiled, err := regexp.Compile(regextext)
	if err != nil {
		return PWCheckerMinRE{}, err
	}
	return PWCheckerMinRE{nameConst: name, regexCompiled: regexCompiled}, nil
}

/*Add Match if  PW has matched Minimun of n regexp  */
func (PWChecker PWCheckerMinRE) AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string {
	isMin := isMinRegexExpression(jsonStruct, PWChecker.regexCompiled, PWChecker.nameConst)
	if !isMin {
		return append(noMatch, PWChecker.nameConst)
	}
	return noMatch
}

/* Check if string has matched Minimun of n regexp   */
func isMinRegexExpression(jsonStructure structJson.PSReceiveStructure, reExp *regexp.Regexp, RuleName string) bool {
	PW := jsonStructure.PW
	minRe, ok := jsonStructure.Rules[RuleName]
	if ok {
		return checkRegexLenPWMin(PW, reExp, minRe)
	} else {
		return true
	}
}

/* Check if regex match minRe Times or more*/
func checkRegexLenPWMin(PW string, reExp *regexp.Regexp, minRe int) bool {
	return (len(reExp.FindAllString(PW, -1)) >= minRe)
}
