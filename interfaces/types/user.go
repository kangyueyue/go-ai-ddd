package types

type (
	RegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Captcha  string `json:"captcha"`
		Password string `json:"password" binding:"gt=0"`
	}

	RegisterResponse struct {
		Response
		Token string `json:"token,omitempty"`
	}

	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"gt=0"`
	}

	LoginResponse struct {
		Response
		Token string `json:"token,omitempty"`
	}

	CaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}

	CaptchaResponse struct {
		Response
	}
)
