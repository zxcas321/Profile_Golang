package response

type UserResponse struct{
	ID        string    `json:"id"`
	Email     string    `json:"email"`
}

func NewUserResponse(id, email string) UserResponse {
	return UserResponse{
		ID:        id,
		Email:     email,
	}
}