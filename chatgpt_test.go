package chatgpt

import (
	"fmt"
	"os"
	"testing"
)

// TestSendMessage tests the send method of the client
func TestSendMessage(t *testing.T) {
	// Prepare test data
	fmt.Println(os.Getenv("BEARER_TOKEN"))
	fmt.Println(os.Getenv("SESSION_TOKEN"))
	chatgpt := NewChatGpt(NewClient(&Credentials{
		BearerToken:  os.Getenv("BEARER_TOKEN"),
		SessionToken: os.Getenv("SESSION_TOKEN"),
	}))
	mockRequest := "hello"

	// Run test
	res, err := chatgpt.SendMessage(mockRequest)
	if err != nil {
		t.Errorf("error sending request, %v", err)
	}

	// Handle response
	fmt.Println(res)
}
