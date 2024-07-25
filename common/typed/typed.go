package typed

type ReqRpcErr struct {
	Code  int      `json:"code"`
	Msg   string   `json:"msg"`
	Temps []string `json:"temps"`
}
type ReqMeta struct {
	UserId     string     `json:"userId"`
	ErrUser    *ReqRpcErr `json:"errUser"`
	Lang       string     `json:"lang"`
	PlatId     string     `json:"platId"`
	PlatClasEm int64      `json:"platClasEm"`
}
