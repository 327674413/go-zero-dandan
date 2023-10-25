package authd

import (
	"context"
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	policy := [][]string{
		[]string{"p", "张三", "文章", "查询"},
		[]string{"p", "张三", "文章", "修改"},
		[]string{"p", "组1", "文章", "删除"},
		[]string{"g", "张三", "组1"},
	}

	auth, err := NewAuth(context.Background(), ModelRBAC, PolicyToLine(policy))
	if err != nil {
		t.Error(err)
		return
	}
	sub := "张三"
	obj := "文章"
	act := "删除"
	res, err := auth.Check(sub, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}
