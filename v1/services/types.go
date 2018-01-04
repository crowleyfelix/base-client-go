package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

//Credentials is the client identifier
type Credentials struct {
	SecretKey            string
	ClientApplicationKey string
	User                 string
}

func (c *Credentials) Signature(httpMethod string, url string) string {
	const divider int64 = 1000000000
	timestamp := time.Now().UnixNano() / divider
	content := fmt.Sprintf("cadu.%s.%s.%s.%d", c.ClientApplicationKey, httpMethod, url, timestamp)
	signature := hmac.New(sha256.New, []byte(c.SecretKey))
	signature.Write([]byte(content))
	mac := hex.EncodeToString(signature.Sum(nil))
	result := fmt.Sprintf(`CADU id=%q, ts="%d", mac=%q`, c.ClientApplicationKey, timestamp, mac)

	return result
}

type response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Paging  *models.Paging  `json:"paging"`
	Errors  []errorData     `json:"errors"`
}

func (r *response) UnmarshalData(data interface{}) errors.Error {
	err := json.Unmarshal(r.Data, data)

	if err != nil {
		return errors.NewSerializing(err.Error())
	}

	return nil
}

type errorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
