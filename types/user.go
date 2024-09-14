package types

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
