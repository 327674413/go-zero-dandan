package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/config"
	"go-zero-dandan/app/im/ws/websocketd"
	"net/http"
)

type ServiceContext struct {
	config.Config
	WsClient websocketd.Client
	modelMongo.ChatLogModel
	modelMongo.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:            c,
		ChatLogModel:      modelMongo.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: modelMongo.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
	}
	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}
	if token == "" {
		logx.Error("ws 客户端的token为空，key为chat_system_root_token")
	}
	header := http.Header{}
	header.Set("token", token)
	svc.WsClient = websocketd.NewClient(c.Ws.Host, websocketd.WithClientHeader(header))
	return svc
}
func (t *ServiceContext) GetSystemToken() (string, error) {
	return t.Ws.SysToken, nil
}
