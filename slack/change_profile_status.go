package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.com/fhuitelec/slack-status/config"
)

type slackProfileResponse struct {
	StatusCode int
	OK         bool   `json:"ok"`
	Error      string `json:"error"`
}

func (resp slackProfileResponse) IsOk() bool {
	return resp.OK || resp.StatusCode != http.StatusOK
}

func ChangeProfileStatus(status string, emoji string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	rawResponse, err := client.Do(getRequest(status, emoji))

	if err != nil || nil != processResponse(rawResponse) {
		log.Fatal(err)
	}

	log.Println(
		fmt.Sprintf("Successfully changed profile status with emoji \"%s\" and status \"%s\"", emoji, status),
	)
}

func getRequest(status string, emoji string) *http.Request {
	var payloadBuffer = fmt.Sprintf(`{
		"status_text": "%s",
		"status_emoji": "%s"
	}`, status, emoji)

	form := url.Values{}
	form.Add("profile", payloadBuffer)
	form.Add("token", config.GetToken())

	request, err := http.NewRequest("POST", "https://slack.com/api/users.profile.set", strings.NewReader(form.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request
}

func processResponse(rawResponse *http.Response) error {
	defer rawResponse.Body.Close()

	response := slackProfileResponse{StatusCode: rawResponse.StatusCode}
	json.NewDecoder(rawResponse.Body).Decode(&response)

	if !response.IsOk() {
		return fmt.Errorf("Unsucessful request to Slack Profile API. Error: %s", response.Error)
	}

	return nil
}
