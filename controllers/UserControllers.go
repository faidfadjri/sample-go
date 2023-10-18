package controllers

import (
	"encoding/json"
	"example/go-api/models"
	"net/http"
)

var data = []models.Student{
    {ID: "E001", Name: "ethan",Grade:  21},
    {ID: "W001", Name: "wick", Grade: 22},
    {ID: "B001", Name: "bourne", Grade: 23},
    {ID: "B002", Name: "bond", Grade: 23},
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

    if req.Method == "GET" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(res, err.Error(), http.StatusInternalServerError)
            return
        }

        res.Write(result)
        return
    }

    http.Error(res, "", http.StatusBadRequest)
}