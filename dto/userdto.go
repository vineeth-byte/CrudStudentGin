package dto

type User struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	EmailId      string `json:"emailId,omitempty" bson:"emailId,omitempty"`
	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	PasswordHash []byte `json:"passwordHash,omitempty" bson:"passwordHash,omitempty"`
}
