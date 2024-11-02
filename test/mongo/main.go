package main

import (
	"context"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/pkg/bitmapd"
)

func main() {
	db := modelMongo.MustChatLogModel("mongodb://root:8a7yNLrsThjw3jra@127.0.0.1:27017", "chat")
	ctx := context.Background()
	chatLog, err := db.FindOne(ctx, "66c06549ebeab36e97929403")
	if err != nil {
		panic(err)
	}
	fmtd.Info(chatLog.MsgReads)
	readRecords := bitmapd.Load(chatLog.MsgReads)
	fmtd.Info(readRecords.IsSetId("1"))
	fmtd.Json(chatLog)
}
