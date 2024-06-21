package ctxd

import "context"

const (
	KeyPlatId     = "platId"
	KeyPlatClasEm = "platClasEm"
	KeyUserId     = "userId"
)

func PlatId(ctx context.Context) string {
	id, _ := ctx.Value(KeyPlatId).(string)
	return id
}

func PlatClasEm(ctx context.Context) int64 {
	id, _ := ctx.Value(KeyPlatClasEm).(int64)
	return id
}
func UserId(ctx context.Context) string {
	id, _ := ctx.Value(KeyUserId).(string)
	return id
}
