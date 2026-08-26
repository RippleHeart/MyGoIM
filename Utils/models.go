package Utils

type LoginForm struct {
	Username string `gorm:"column:name" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
type JWToken struct {
	Token string
}
