// Code generated by goctl. DO NOT EDIT.
package types

type AccountLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AccountLoginResp struct {
	Token string `json:"token"`
}
