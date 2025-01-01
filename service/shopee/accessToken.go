package shopee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"shopee-jurnal-payment/helper"
)

type AccessTokenRequest struct {
	Code      string `json:"code"`
	PartnerID int    `json:"partner_id"`
	ShopID    int    `json:"shop_id"` // input 1 only.
	// MainAccountID int    `json:"main_account_id"` // input 1 only.
}

type AccessTokenResponse struct {
	RequestID      string `json:"request_id"`
	Error          string `json:"error"`
	RefreshToken   string `json:"refresh_token"`
	AccessToken    string `json:"access_token"`
	ExpireIn       int    `json:"expire_in"`
	Message        string `json:"message"`
	MerchantIDList []int  `json:"merchant_id_list"`
	ShopIDList     []int  `json:"shop_id_list"`
}
type AccessToken interface {
	GetAccessToken(request *AccessTokenRequest) (*AccessTokenResponse, error)
}

type AccessTokenClient struct {
	PartnerKey string
	Host       string
}

func NewAccessTokenClient(partnerKey, host string) *AccessTokenClient {
	return &AccessTokenClient{PartnerKey: partnerKey, Host: host}
}
func (a *AccessTokenClient) GetAccessToken(request *AccessTokenRequest) (*AccessTokenResponse, error) {

	signatureData := helper.NewPublicShopeeAPISignatureData(request.PartnerID, PathAccessToken)
	sign := helper.SignShopeeSignature(a.PartnerKey, signatureData)
	// Build URL with query parameters

	url := fmt.Sprintf("%s%s?partner_id=%d&timestamp=%s&sign=%s",
		a.Host,
		PathAccessToken,
		request.PartnerID,
		signatureData.Timestamp,
		sign,
	)
	log.Println("HERE : ", url)
	log.Println("HERE : ", signatureData.Timestamp)
	log.Println("HERE : ", a.PartnerKey)

	// Create request body
	jsonBody, err := json.Marshal(request)
	if err != nil {
		// Handle error
		return nil, err
	}
	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		// Handle error
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		return nil, err
	}

	// Parse response into your struct
	var response AccessTokenResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
