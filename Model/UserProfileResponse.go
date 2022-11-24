package Model

type UserProfileResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	Address      string `json:"address"`
}
