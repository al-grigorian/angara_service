package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)


type WeightRequest struct {
	AccessKey int64  `json:"access_key"`
	Weight int `json:"weight"`
}

type Request struct {
	ApplicationId int64 `json:"application_id"`
}


func (h *Handler) issueWeight(c *gin.Context) {
	var input Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueWeight:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(4 * time.Second)
		sendWeightRequest(input)
	}()
}

func sendWeightRequest(request Request) {

	var weight = -1
	if rand.Intn(10) % 10 >= 2 {
	 weight = 500 + rand.Intn(1000)
	}

	answer := WeightRequest{
		AccessKey: 123,
		Weight: weight,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/applications/%d/update_weight/", request.ApplicationId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Фактический вес: ", weight)
	fmt.Println("PUT Request Status:", response.Status)
}
