package main

import (
	"log"
	"regexp"

	"passwordcheck/interfaces"
	"passwordcheck/internal/handler"
	"passwordcheck/internal/services"

	envVar "passwordcheck/internal/envVarStruct"
)

/* true if a variables are setted according to test for this api restful api */
var debug bool = true

/*
init vars
*/
var env envVar.EnvVar
var PWService *interfaces.PWService
var handlerPWCheck *handler.HandlerPWCheck

/* */
func init() {
	if debug {
		env = envVar.DebugEnvVariables()
	} else {
		env = envVar.GetEnvVariables()
	}
	PWService, err := services.NewPWService(env.UpperCaseString,
		env.LowerCaseString,
		env.DigitsString,
		env.SpecialCharacters,
		env.MinLenString,
		env.RepeatedString,
		`[A-Z]`,
		`[a-z]`,
		`[0-9]`,
		"["+regexp.QuoteMeta(`!@#$%^&*()-+\/{}[]`)+"]",
	)
	if err != nil {
		log.Print(err)
		return
	}
	handlerPWCheck = handler.NewHandler(PWService)
}

/* */
func main() {
	handlerPWCheck.PWRouter()
}
