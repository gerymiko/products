package models

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type ResponseUser struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token omitempty"`
}
