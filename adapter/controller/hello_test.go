package controller

import (
	"clean-serverless-book-sample/mocks"
	"encoding/json"

	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestPosetHello_400(t *testing.T) {
	tables := mocks.SetupDB(t)
	defer tables.Cleanup()

	body := map[string]interface{}{
		"name": "",
	}

	bodyStr, err := json.Marshal(body)
	assert.NoError(t, err)

	req := events.APIGatewayProxyRequest{
		Body: string(bodyStr),
	}

	res := PostHello(req)

	var resBody map[string]interface{}
	err = json.Unmarshal([]byte(res.Body), &resBody)
	assert.NoError(t, err)
	assert.Equal(t, 400, res.StatusCode)

	errs := resBody["errors"].(map[string]interface{})
	assert.Equal(t, "名前を入力してください", errs["name"])
}
