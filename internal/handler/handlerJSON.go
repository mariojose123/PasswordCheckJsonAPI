/*This package handler its the server and handler function and  handle everything related to http request on password check project*/
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"passwordcheck/interfaces"
	"passwordcheck/internal/structJson"
)

type HandlerPWCheck struct {
	passService interfaces.PWService
}

/*Create a new handler with its service */
func NewHandler(PWService interfaces.PWService) *HandlerPWCheck {
	return &HandlerPWCheck{passService: PWService}
}

/*Set the URL for the API */
func (h *HandlerPWCheck) PWRouter() {
	http.HandleFunc("/verify", h.verify)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*This function receives the JSON according to the specification on HTTP.request and send on w the JSON with the answer*/
func (h *HandlerPWCheck) verify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.postVerify(w, r)
	} else {
		WrongMethodVerify(w)
	}

}

/* POST  part of verifying func and the part required according to the specification because all API calls will be valid*/
func (h *HandlerPWCheck) postVerify(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	jsonPW, err := DecodeJsonVerify(r)
	if err != nil {
		log.Print(err)
		BadRequestVerify(w)
	}

	isCorrectPass, matchedRules := h.passService.CheckPW(ctx, jsonPW)

	jsonstruct := structJson.PStoJsonResponse(isCorrectPass, matchedRules)
	SendJsonVerify(w, jsonstruct)
}

/*A method is call for every wrong Json on an API call for this Restful API*/
func WrongMethodVerify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func BadRequestVerify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

/*Send Struct of service check PW as JSON to client*/
func SendJsonVerify(w http.ResponseWriter, jsonstruct structJson.JsonResponse) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonstruct)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

/* Decode the JSON for verify, it is used on CheckPW service function*/
func DecodeJsonVerify(r *http.Request) (structJson.PSReceiveStructure, error) {
	var jsonPWraw structJson.PSReceiveStructureRaw
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&jsonPWraw)
	if err != nil {
		return structJson.PSReceiveStructure{}, err
	}
	jsonPW := structJson.NewPSReceiveStructure(jsonPWraw)
	return jsonPW, nil
}
