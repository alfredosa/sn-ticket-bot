package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Payload struct {
	Variables interface{} `json:"variables"`
}

type OnboardingVariables struct {
	RequestFor       string `json:"requested_for"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	PictureIDCard    string `json:"attach_picture_for_the_id_card"`
}

func CreateServiceNowTicket() error {
	val, err := RequestAccessToken(os.Getenv("AZ_TENANT_ID"))
	if err != nil {
		fmt.Printf("Error with token processing %s\n", err)
		return err
	}

	image, err := readPicture()
	if err != nil {
		fmt.Println("Why is life so hard, just read the pic")
	}
	description := "This comes from a golang application, and its intended as a POC"
	shortDescription := "golang service now ticket"
	requestedFor := "00860289"

	onboardingVars := Payload{
		Variables: OnboardingVariables{
			RequestFor:       requestedFor,
			Description:      description,
			ShortDescription: shortDescription,
			PictureIDCard:    image,
		},
	}

	err = PostSNCatalogItem(&onboardingVars, val)
	if err != nil {
		slog.Error("Unable to post RITM ticket", "error", err)
		return err
	}

	return nil
}

func PostSNCatalogItem(payload interface{}, token string) error {
	json, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Printf("Sending Json %v", len(json))

	return nil
}

func readPicture() (string, error) {
	base64string, err := imgToBase64("assets/smile.jpg")
	if err != nil {
		slog.Error("Error reading base64 string", "error", err)
		return "", err
	}

	return base64string, nil
}

func imgToBase64(filename string) (string, error) {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "." // default to current directory if not set
	}
	filepath := filepath.Join(basePath, filename)

	img, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading the file %s\n", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(img)
	return base64Str, nil
}
