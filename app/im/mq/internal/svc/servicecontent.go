package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/internal/config"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/social/rpc/social"
	"go-zero-dandan/common/interceptor"
	"net/http"
)

type ServiceContext struct {
	config.Config
	WsClient websocketd.Client
	modelMongo.ChatLogModel
	modelMongo.ConversationModel
	modelMongo.SysMsgLogModel
	modelMongo.SysMsgStatModel
	SocialRpc social.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	socialRpc := social.NewSocial(zrpc.MustNewClient(c.SocialRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	svc := &ServiceContext{
		Config:            c,
		ChatLogModel:      modelMongo.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: modelMongo.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
		SysMsgLogModel:    modelMongo.MustSysMsgLogModel(c.Mongo.Url, c.Mongo.Db),
		SysMsgStatModel:   modelMongo.MustSysMsgStatModel(c.Mongo.Url, c.Mongo.Db),
		SocialRpc:         socialRpc,
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
