package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type PublicShopeeAPISignatureData struct {
	PartnerID int
	ApiPath   string
	Timestamp string
}

func NewPublicShopeeAPISignatureData(partnerID int, apiPath string) *PublicShopeeAPISignatureData {
	return &PublicShopeeAPISignatureData{
		PartnerID: partnerID,
		ApiPath:   apiPath,
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}
}
func SignShopeeSignature(partnerKey string, data *PublicShopeeAPISignatureData) string {
	tmpBaseString := fmt.Sprintf("%d%s%s", data.PartnerID, data.ApiPath, data.Timestamp)
	// Convert to bytes
	baseString := []byte(tmpBaseString)

	// Create HMAC-SHA256 hash
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	return sign
}
