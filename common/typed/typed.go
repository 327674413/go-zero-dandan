package typed

type ReqMeta struct {
	UserId     string `json:"userId"`
	UserErr    string `json:"userErr"`
	Lang       string `json:"lang"`
	PlatId     string `json:"platId"`
	PlatClasIm int64  `json:"platClasIm"`
}
