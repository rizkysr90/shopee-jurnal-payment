package main

import (
	"log"
	"net/http"
	"shopee-jurnal-payment/service/shopee"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	PartnerID           = 00000
	PartnerKey          = ""
	DomainOpenAPIShopee = "https://partner.test-stable.shopeemobile.com"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/callback/shopee", func(ctx *gin.Context) {
		log.Println(ctx.Request.URL)
		code := ctx.Query("code")
		shopID := ctx.Query("shop_id")
		shopIDInt, err := strconv.Atoi(shopID)
		if err != nil {
			ctx.Error(err)
		}
		client := shopee.NewAccessTokenClient(PartnerKey, DomainOpenAPIShopee)
		request := &shopee.AccessTokenRequest{Code: code, PartnerID: PartnerID, ShopID: shopIDInt}
		response, err := client.GetAccessToken(request)
		if err != nil {
			ctx.Error(err)
		}

		ctx.JSON(200, gin.H{"response": response})
	})
	// Run the server on 127.0.0.1:8080
	r.Run("127.0.0.1:8080")

}
