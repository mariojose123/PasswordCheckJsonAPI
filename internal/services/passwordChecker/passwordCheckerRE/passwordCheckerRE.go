/* package for password checker that check if password Regular Expression is Matched x times in the password string*/
package passwordCheckerRE

import (
	"regexp"

	"github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"
)

/*Struct that checks repetition of regular Expression in the PW */
type PWCheckerMinRE struct {
	regexCompiled *regexp.Regexp
	nameConst     string
}

/*Create Password Checker(PWCheckerMinRE) that checks if a regular Expression happens x times */
func NewPWCheckRE(name string, regextext string) (PWCheckerMinRE, error) {
	regexCompiled, err := regexp.Compile(regextext)
	if err != nil {
		return PWCheckerMinRE{}, err
	}
	return PWCheckerMinRE{nameConst: name, regexCompiled: regexCompiled}, nil
}

/* Create Password Checker(PWCheckerMinRE) FOR Upper Case characters Min number of times*/
func NewPWCheckREUpperCase(name string) (PWCheckerMinRE, error) {
	return NewPWCheckRE(name, `[A-Z]`)
}

/* Create Password Checker(PWCheckerMinRE) FOR Digits characters Min number of times*/
func NewPWCheckREDigits(name string) (PWCheckerMinRE, error) {
	return NewPWCheckRE(name, `[0-9]`)
}

/* Create Password Checker(PWCheckerMinRE) FOR Special  characters Min number of times ACCORDING TO SPECIFICATION*/
func NewPWCheckRESpecialchar(name string) (PWCheckerMinRE, error) {
	return NewPWCheckRE(name, "["+regexp.QuoteMeta(`!@#$%^&*()-+\/{}[]`)+"]")
}

/* Create Password Checker(PWCheckerMinRE) FOR Lower Case characters Min number of times*/
func NewPWCheckRELowedCase(name string) (PWCheckerMinRE, error) {
	return NewPWCheckRE(name, `[a-z]`)
}

/*Add Match if  PW has matched a Minimum of n regexp matchs  */
func (PWChecker PWCheckerMinRE) AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string {
	isMin := isMinRegexExpression(jsonStruct, PWChecker.regexCompiled, PWChecker.nameConst)
	if !isMin {
		return append(noMatch, PWChecker.nameConst)
	}
	return noMatch
}

/* Check if string has matched Minimum of n regexp */
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
