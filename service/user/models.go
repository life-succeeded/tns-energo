package user

type RegisterRequest struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Surname    string  `json:"surname"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic"`
	Position   *string `json:"position"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
