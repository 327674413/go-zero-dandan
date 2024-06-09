package svc

import (
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mode   string
	modelMongo.ChatLogModel
	modelMongo.ConversationModel
	modelMongo.ConversationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		Mode:               c.Mode,
		ChatLogModel:       modelMongo.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel:  modelMongo.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
		ConversationsModel: modelMongo.MustConversationsModel(c.Mongo.Url, c.Mongo.Db),
	}
}
