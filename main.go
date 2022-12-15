/*
The main Program Logic is that a service with a list of password Checkers for every rule  will be created and
it will execute a function on handlerPWCheck, the struct that deals with HTTP requests
*/
package main

import (
	"log"

	"github.com/mariojose123/PasswordCheckJsonAPI/interfaces"
	envVar "github.com/mariojose123/PasswordCheckJsonAPI/internal/envVarStruct"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/handler"
	"github.com/mariojose123/PasswordCheckJsonAPI/internal/pwCheckerBuilder"
)

/*
	true if all variables are setted according to test and specifications  for this api restful api according to the requirements

false if all variables are setted acording to dockerfile or enviroment variables
*/

/*
init variables Enviroment Variables Struct ,Service Struct with password checkers for handler HandlerPWCheck with http functions for API
*/
var env envVar.EnvVar
var PWService *interfaces.PWService
var handlerPWCheck *handler.HandlerPWCheck

/*
Create every struct required for the server a Builder will be created and then
this function will be a shorter function
*/
func init() {
	env = envVar.NoEnvVariables()

	var err error
	handlerPWCheck, err = pwCheckerBuilder.NewBuilderpwCheckHandler(
		env.UpperCaseString,
		env.LowerCaseString,
		env.DigitsString,
		env.SpecialCharacters,
		env.MinLenString,
		env.RepeatedString,
	)
	if err != nil {
		log.Print(err)
		return
	}
}

/* main with the http server start function */
func main() {
	handlerPWCheck.PWRouter()
}
