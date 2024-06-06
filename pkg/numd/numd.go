package numd

import (
	"fmt"
	"sort"
)

// CombineInt64 按小到大用下划线拼接
func CombineInt64(aid, bid int64) string {
	ids := []int64{aid, bid}

	sort.Slice(ids, func(i, j int) bool {
		return i < j
	})
	return fmt.Sprintf("%d_%d", ids[0], ids[1])
}
