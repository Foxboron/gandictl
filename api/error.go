package api

import (
	"encoding/json"
)

type ErrorReponse struct {
	Status  int
	Cause   string
	Code    int
	Message string
	Object  string
}

func NewErrorResponse(resp []byte) *ErrorReponse {
	var err ErrorReponse
	json.Unmarshal(resp, &err)
	return &err
}

type SuccsessResponse struct {
	Message string `json:"message"`
}

func NewSuccsessResponse(resp []byte) *SuccsessResponse {
	var msg SuccsessResponse
	json.Unmarshal(resp, &msg)
	return &msg
}
