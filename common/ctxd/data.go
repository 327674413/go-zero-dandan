package ctxd

import "context"

const (
	KeyPlatId     = "platId"
	KeyPlatClasEm = "platClasEm"
	KeyUserId     = "userId"
)

func GetPlatId(ctx context.Context) string {
	id, _ := ctx.Value(KeyPlatId).(string)
	return id
}

func GetPlatClasEm(ctx context.Context) int64 {
	id, _ := ctx.Value(KeyPlatClasEm).(int64)
	return id
}
func GetUserId(ctx context.Context) string {
	id, _ := ctx.Value(KeyUserId).(string)
	return id
}
