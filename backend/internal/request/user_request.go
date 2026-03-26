package request

type CreateUserRequest struct{
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type UpdateUserRequest struct{
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}