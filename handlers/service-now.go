package handlers

import (
	"fmt"
	"os"
)

func CreateServiceNowTicket() error {
	val, err := RequestAccessToken(os.Getenv("AZ_TENANT_ID"))

	if err != nil {
		fmt.Printf("Error with token processing %s\n", err)
		return err
	}

	fmt.Printf("value from token %s\n", val)
	return nil
}
