package controllers

import (
	"atp-go-rest/models"
	"atp-go-rest/utils"
	"encoding/json"
	"net/http"
)

func AddCoach(w http.ResponseWriter, r *http.Request) {
	var coach models.Coach
	err := json.NewDecoder(r.Body).Decode(&coach)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Message: "wrong format",
			Data: err,
		})
	}
	insertID, err := coach.Insert()
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Message: "insert failed",
			Data: err,
		})
	}
	utils.ResponseWithJson(w, http.StatusOK, utils.Response{
		StatusCode: http.StatusOK,
		Message:    "Success!",
		Data:       insertID,
	})
}
