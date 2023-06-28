package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todolist/config"
	"go-todolist/helper"
)

func SignatureMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//1. Validate signature dari header yang dikirim
		//2. Jika signature valid, maka lanjut ke handler
		//3. Jika signature tidak valid, maka return error

		//validation signature
		signature := c.Get("X-Gic-Signature")
		url := string(c.Context().Path())
		JwtToken := c.Get("Authorization")
		method := c.Method()
		timeStamp := "2023-10-10 12:12:12"

		//create json payload
		jsonPayload := `{"path":"` + url + `","jwt":"` + JwtToken + `","method":"` + method + `","timestamp":"` + timeStamp + `"}`
		fmt.Println(jsonPayload)

		//create new signature from backend
		signatureBackend := ValidateSignature(signature, jsonPayload)
		fmt.Println(signatureBackend)

		if signatureBackend {
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Signature failed", "Signature failed", nil))
		}
	}
}

func ValidateSignature(signature string, signatureCompare string) bool {
	//make compare signature and signature from backend
	//signature from backend
	secretKey := config.SignatureSecret
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signatureCompare))
	signatureBackend := fmt.Sprintf("%x", h.Sum(nil))              //balikan ini adalah hex
	return hmac.Equal([]byte(signature), []byte(signatureBackend)) //compare signature from header and signature from backend
}
