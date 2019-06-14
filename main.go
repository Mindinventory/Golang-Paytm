package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"paytm/paytm"
)

type PaytmOrderResponse struct {
	MID              string `json:"MID"`
	ORDER_ID         string `json:"ORDER_ID"`
	CUST_ID          string `json:"CUST_ID"`
	INDUSTRY_TYPE_ID string `json:"INDUSTRY_TYPE_ID"`
	CHANNEL_ID       string `json:"CHANNEL_ID"`
	TXN_AMOUNT       string `json:"TXN_AMOUNT"`
	WEBSITE          string `json:"WEBSITE"`
	CALLBACK_URL     string `json:"CALLBACK_URL"`
	EMAIL            string `json:"EMAIL"`
	MOBILE_NO        string `json:"MOBILE_NO"`
	PAYMENT_TYPE_ID  string `json:"PAYMENT_TYPE_ID"`
	CHECKSUMHASH     string `json:"CHECKSUMHASH"`
}

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func paytmOrderResponseFun(orderId string, custId string, mobileId string, email string, txnAmount string, website string, callbackUrl string, checksum string, channelId string, verifiedBy string, isUserVerified string, paymentModeOnly string, authMode string, paymentTypeId string) PaytmOrderResponse {
	res := PaytmOrderResponse{
		MID:              os.Getenv("PAYTM_MID"),
		ORDER_ID:         orderId,
		CUST_ID:          custId,
		INDUSTRY_TYPE_ID: os.Getenv("PAYTM_INDUSTRY_TYPE_ID"),
		CHANNEL_ID:       channelId,
		TXN_AMOUNT:       txnAmount,
		WEBSITE:          website,
		CALLBACK_URL:     callbackUrl,
		EMAIL:            email,
		MOBILE_NO:        mobileId,
		CHECKSUMHASH:     checksum,
		PAYMENT_TYPE_ID:  paymentTypeId,
	}
	return res
}

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	var ORDER_ID = "234234234"
	var MOBILE_NO = "9586311403"
	var EMAIL = "xyz@gmail.com"
	var TXN_AMOUNT = "45.00"
	var VERIFIED_BY = "MOBILE"
	var IS_USER_VERIFIED = "YES"
	var PAYMENT_MODE_ONLY = "YES"
	var AUTH_MODE = "USRPWD"
	var PAYMENT_TYPE_ID = "PPI"

	var CUST_ID = "5647478899090"
	var channelID = os.Getenv("PAYTM_CHANNEL_ID")
	var website = os.Getenv("PAYTM_WEBSITE")

	//channelID = os.Getenv("PAYTM_CHANNEL_ID")
	//website = os.Getenv("PAYTM_WEBSITE")

	callbackUrl := os.Getenv("PAYMENT_MIDDLEWARE_URL")
	paramList := map[string]string{
		"MID":              os.Getenv("PAYTM_MID"),
		"ORDER_ID":         ORDER_ID,
		"CUST_ID":          CUST_ID,
		"INDUSTRY_TYPE_ID": os.Getenv("PAYTM_INDUSTRY_TYPE_ID"),
		"CHANNEL_ID":       channelID,
		"TXN_AMOUNT":       TXN_AMOUNT,
		"WEBSITE":          website,
		"CALLBACK_URL":     callbackUrl,
		"EMAIL":            EMAIL,
		"MOBILE_NO":        MOBILE_NO,
		"PAYMENT_TYPE_ID":  PAYMENT_TYPE_ID,
	}

	checksum, err := paytm.GetChecksumFromArray(paramList)

	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("checksum", checksum)
	}

	orderResponse := paytmOrderResponseFun(ORDER_ID, CUST_ID, MOBILE_NO, EMAIL, TXN_AMOUNT, website, callbackUrl, checksum, channelID, VERIFIED_BY, IS_USER_VERIFIED, PAYMENT_MODE_ONLY, AUTH_MODE, PAYMENT_TYPE_ID)

	fmt.Println("Response", orderResponse)
}
