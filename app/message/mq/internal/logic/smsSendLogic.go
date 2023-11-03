package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/mq/internal/svc"
	"strings"
)

type SmsSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsSendLogic {
	return &SmsSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SmsSendLogic) Consume(key, val string) error {
	fmt.Printf("get key: %s val: %s\n", key, val)
	err := l.SmsLog()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("es成功")
	}
	return nil
}
func (l *SmsSendLogic) SmsLog() error {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.Es.Client,
		Index:  "sms-index",
	})
	if err != nil {
		return err
	}
	type sms struct {
		Phone   string
		Content string
	}
	d := sms{
		Phone:   "15111111111",
		Content: "kl阿迪舅舅看了多少分看了是打发了空间圣诞快乐sad副经理看的说法",
	}
	v, err := json.Marshal(d)
	if err != nil {
		return err
	}

	payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
	err = bi.Add(l.ctx, esutil.BulkIndexerItem{
		Action:     "update",
		DocumentID: d.Phone,
		Body:       strings.NewReader(payload),
		OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
		},
		OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
		},
	})

	if err != nil {
		return err
	}

	return bi.Close(l.ctx)
}
