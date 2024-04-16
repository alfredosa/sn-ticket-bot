package handlers

import "os"

func CreateServiceNowTicket() {
	RequestAccessToken(os.Getenv("AZ_TENANT_ID"))
}
