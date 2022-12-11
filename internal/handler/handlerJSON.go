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

func NewHandler(PWService interfaces.PWService) *HandlerPWCheck {
	return &HandlerPWCheck{passService: PWService}
}

func (h *HandlerPWCheck) PWRouter() {
	http.HandleFunc("/verify", h.verify)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h *HandlerPWCheck) verify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.postVerify(w, r)
	} else {
		WrongMethodVerify(w)
	}

}

/* Post part of verify func*/
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

func WrongMethodVerify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func BadRequestVerify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

/*Send Struct of service check PW as json to client*/
func SendJsonVerify(w http.ResponseWriter, jsonstruct structJson.JsonResponse) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonstruct)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

/* Decode the json for verify, it is used on CheckPW service function*/
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
