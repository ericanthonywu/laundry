package Model

type UserQuotaProfileResponse struct {
}

type UpdateProfile struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	Address      string `json:"address"`
}

func (req UpdateProfile) IsValid() bool {
	return req.Name != "" && req.PhoneNumber != "" && req.Address != "" && req.EmailAddress != ""
}
