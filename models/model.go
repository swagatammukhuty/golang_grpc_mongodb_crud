package models

type User struct {
	Age      int32  `bson:"age,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	UserId   string `bson:"user_id,omitempty"`
}
