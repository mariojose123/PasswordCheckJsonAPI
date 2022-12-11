package structJson

type JsonResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}

/* Structure that is created from json received from Verify POST API call  for PW Check */
type PSReceiveStructure struct {
	PW    string
	Rules map[string]int
}

/* Raw structure received by json*/
type PSReceiveStructureRaw struct {
	PW    string
	Rules []Rule
}
type Rule struct {
	Rule  string
	Value int
}

/* Transform PSReceiveStructureRaw To P5ReceiveStructure*/
func NewPSReceiveStructure(raw PSReceiveStructureRaw) PSReceiveStructure {
	newPSReceiveStructure := PSReceiveStructure{raw.PW, map[string]int{}}
	for _, rule := range raw.Rules {
		newPSReceiveStructure.Rules[rule.Rule] = rule.Value
	}
	return newPSReceiveStructure
}

/*Transform PW Checker returned structure into a Json*/
func PStoJsonResponse(isCorrectPass bool, matchedRules []string) JsonResponse {
	var jsonStruct JsonResponse = JsonResponse{Verify: isCorrectPass, NoMatch: matchedRules}
	return jsonStruct
}
