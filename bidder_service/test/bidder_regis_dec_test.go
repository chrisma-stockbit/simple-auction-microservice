package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	requestdecoder "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/codec/request"
	dto "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
	"github.com/stretchr/testify/assert"
)

func TestDecodeRegisterBidderRequest(t *testing.T) {
	fullName := "Test Name"
	accountNumber := "123"
	bankName := "BTEST"
	phoneNumber := "081212345"

	jsonString := fmt.Sprintf(`{"fullName": "%s", "accountNumber": "%s", "bankName": "%s", "phoneNumber": "%s"}`,
		fullName,
		accountNumber,
		bankName,
		phoneNumber,
	)
	body := []byte(jsonString)
	req, _ := http.NewRequest("", "", bytes.NewBuffer(body))
	resp, _ := requestdecoder.DecodeRegisterBidderRequest(nil, req)

	bidderRegisReq, ok := resp.(dto.BidderRegistrationRequest)
	if !ok {
		t.Errorf("Fail to decode bidder registration request")
	}
	assert.Equal(t, fullName, bidderRegisReq.FullName, "wrong full name")
	assert.Equal(t, accountNumber, bidderRegisReq.AccountNumber, "wrong account number")
	assert.Equal(t, bankName, bidderRegisReq.BankName, "wrong bank name")
	assert.Equal(t, phoneNumber, bidderRegisReq.PhoneNumber, "wrong phone number")
}
