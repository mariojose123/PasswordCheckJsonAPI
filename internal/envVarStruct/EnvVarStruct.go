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

func GetEnvVariables() EnvVar {
	return EnvVar{os.Getenv("UpperCaseString"),
		os.Getenv("LOWERCASESTRING"),
		os.Getenv("DIGITSSTRING"),
		os.Getenv("SPECIALCHARACTERS"),
		os.Getenv("MINLENSTRING"),
		os.Getenv("REPEATEDSTRING")}
}

func DebugEnvVariables() EnvVar {
	return EnvVar{"minUppercase",
		"minLowercase",
		"minDigit",
		"minSpecialChars",
		"minSize",
		"noRepeted"}
}
