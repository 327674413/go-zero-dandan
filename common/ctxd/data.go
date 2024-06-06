package ctxd

import "context"

const (
	KeyPlatId     = "platId"
	KeyPlatClasEm = "platClasEm"
	KeyUserId     = "userId"
)

func GetPlatId(ctx context.Context) int64 {
	id, _ := ctx.Value(KeyPlatId).(int64)
	return id
}

func GetPlatClasEm(ctx context.Context) int64 {
	id, _ := ctx.Value(KeyPlatClasEm).(int64)
	return id
}
func GetUserId(ctx context.Context) int64 {
	id, _ := ctx.Value(KeyUserId).(int64)
	return id
}
