package biz

import (
	"encoding/json"
	"time"
)

type ApplyRecord struct {
	UserId   string `json:"userId"`
	TimeAt   int64  `json:"timeAt"`
	Text     string `json:"text"`
	TypeEm   int64  `json:"typeEm"`
	SourceEm int64  `json:"sourceEm"`
}

func AddApplyRecord(oldContent string, userId string, applyMsg string, sourceEm int64, typeEm int64) string {
	msgList := make([]*ApplyRecord, 0)
	if oldContent != "" {
		_ = json.Unmarshal([]byte(oldContent), &msgList)
	}
	msgList = append([]*ApplyRecord{
		{
			UserId:   userId,
			TimeAt:   time.Now().Unix(),
			Text:     applyMsg,
			SourceEm: sourceEm,
			TypeEm:   typeEm,
		},
	}, msgList...)
	newContent, _ := json.Marshal(msgList)
	return string(newContent)
}
