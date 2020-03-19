package request

type BidderRegistrationRequest struct {
	FullName string `json:"fullName"`
	AccountNumber string `json:"accountNumber"`
	BankName string `json:"bankName"`
	PhoneNumber string `json:"phoneNumber"`
}