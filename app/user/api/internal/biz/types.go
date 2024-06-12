package biz

type UserRegInfo struct {
	Id        string
	UnionId   string
	Account   string
	Password  string
	Code      string
	Nickname  string
	Phone     string
	PhoneArea string
	Email     string
	Avatar    string
	SexEm     int64
}
type UserEditInfo struct {
	Id       string
	Code     *string
	Nickname *string
	Email    *string
	Avatar   *string
	SexEm    *int64
}
