/*Env struct and functions that set Received Json Variables Names*/
package envVar

import (
	"os"
)

/*Structure with all env variables*/
type EnvVar struct {
	UpperCaseString   string
	LowerCaseString   string
	DigitsString      string
	SpecialCharacters string
	MinLenString      string
	RepeatedString    string
}

/* Get Enviroment Variables of OS*/
func GetEnvVariables() EnvVar {
	return EnvVar{os.Getenv("UpperCaseString"),
		os.Getenv("LOWERCASESTRING"),
		os.Getenv("DIGITSSTRING"),
		os.Getenv("SPECIALCHARACTERS"),
		os.Getenv("MINLENSTRING"),
		os.Getenv("REPEATEDSTRING")}
}

/* Get Default Json Names*/
func NoEnvVariables() EnvVar {
	return EnvVar{"minUppercase",
		"minLowercase",
		"minDigit",
		"minSpecialChars",
		"minSize",
		"noRepeted"}
}
