package classis

type Login struct {
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	ID           string `json:"_id,omitempty"`
	Token        string `json:"token,omitempty"`
	TokenExpires string `json:"tokenExpires,omitempty"`
}
