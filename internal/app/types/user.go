package types

type (
	User struct {
		UserName string `json:"username,omitempty" bson:"username,omitempty"`
		Password string `json:"password,omitempty" bson:"password,omitempty"`
	}
)

func (user *User) Strip() *User {
	stripedUser := User(*user)
	stripedUser.Password = ""
	return &stripedUser
}
