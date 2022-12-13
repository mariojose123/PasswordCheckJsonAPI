package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"passwordcheck/internal/structJson"
	mock_interfaces "passwordcheck/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandlerJson(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	mockPWService := mock_interfaces.NewMockPWService(control)
	PSReceiveStructure := structJson.PSReceiveStructure{PW: "TesteSenhaForte!123&", Rules: map[string]int{"minSize": 8, "minDigit": 4}}
	mockPWService.EXPECT().CheckPW(context.Background(), PSReceiveStructure).Return(false, []string{"minDigit"})
	PostStruct := map[string]interface{}{
		"password": "TesteSenhaForte!123&",
		"rules": append([]map[string]interface{}{},
			map[string]interface{}{"rule": "minSize", "value": 8},
			map[string]interface{}{"rule": "minDigit", "value": 4},
		),
	}
	var bufRequest bytes.Buffer
	err := json.NewEncoder(&bufRequest).Encode(PostStruct)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/verify", &bufRequest)
	w := httptest.NewRecorder()
	NewHandler(mockPWService).postVerify(w, req)
	res := w.Result()

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	wantstring := "{\"verify\":false,\"noMatch\":[\"minDigit\"]}\n"
	retstring := string(data)
	if !reflect.DeepEqual(retstring, wantstring) {
		t.Errorf("got %s, want %s", retstring, wantstring)
	}
}
