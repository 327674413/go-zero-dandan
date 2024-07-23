package svc

import (
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/mqClient"
	"go-zero-dandan/app/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mode   string
	modelMongo.ChatLogModel
	modelMongo.ConversationModel
	modelMongo.ConversationsModel
	modelMongo.SysMsgLogModel
	modelMongo.SysMsgStatModel
	mqClient.SysToUserTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		Mode:                    c.Mode,
		ChatLogModel:            modelMongo.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel:       modelMongo.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
		ConversationsModel:      modelMongo.MustConversationsModel(c.Mongo.Url, c.Mongo.Db),
		SysMsgLogModel:          modelMongo.MustSysMsgLogModel(c.Mongo.Url, c.Mongo.Db),
		SysMsgStatModel:         modelMongo.MustSysMsgStatModel(c.Mongo.Url, c.Mongo.Db),
		SysToUserTransferClient: mqClient.NewSysToUserTransferClient(c.SysToUserTransfer.Addrs, c.SysToUserTransfer.Topic),
	}
}
