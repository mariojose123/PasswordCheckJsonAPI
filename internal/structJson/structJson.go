/*Package that contains every Json Structure and Transformations of Json Structures  and Json Related Functions */
package structJson

/*Json that is the awnser for the API call*/
type JsonResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}

/* Structure that is created from JSON received from Verify POST API call  for Password Checker Service */
type PSReceiveStructure struct {
	PW    string         `json:"password"`
	Rules map[string]int `json:"rules"`
}

/* The raw structure received by JSON*/
type PSReceiveStructureRaw struct {
	PW    string `json:"password"`
	Rules []Rule `json:"rules"`
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
