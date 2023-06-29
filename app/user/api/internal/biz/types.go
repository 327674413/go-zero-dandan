package biz

type UserRegInfo struct {
	Id        int64
	UnionId   int64
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
	Id       int64
	Code     *string
	Nickname *string
	Email    *string
	Avatar   *string
	SexEm    *int64
}
