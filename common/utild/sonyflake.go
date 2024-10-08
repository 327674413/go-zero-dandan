package utild

import (
	"fmt"
	"github.com/sony/sonyflake"
	"sort"
	"strconv"
)
import (
	"math/rand"
	"sync"
	"time"
)

var sf *sonyflake.Sonyflake

var rg = struct {
	sync.Mutex
	rand *rand.Rand
}{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

func init() {
	// todo::如何自动设置不重复的machineID
	var f sonyflake.Settings
	f.MachineID = func() (uint16, error) {
		return 1111, nil
	}
	f.StartTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	sf = sonyflake.NewSonyflake(f)
	if sf == nil {
		panic("sonyflake init error.")
	}
	fmt.Println("--------------------sonyflake Init ------------------")
}

func Int63nRange(min, max int64) int64 {
	rg.Lock()
	defer rg.Unlock()
	return rg.rand.Int63n(max-min) + min
}

func MakeId() string {
	ret, err := sf.NextID()
	if err != nil {
		return fmt.Sprintf("%d", Int63nRange(1926425572, 9223372036854775806))
	}
	//Note: Sonyflake currently does not use the most significant bit of IDs, so you can convert Sonyflake IDs from uint64 to int64 safely.
	return fmt.Sprintf("%d", ret)
}

func CombineId(aid, bid string) string {
	ids := []string{aid, bid}

	sort.Slice(ids, func(i, j int) bool {
		a, _ := strconv.ParseUint(ids[i], 0, 64)
		b, _ := strconv.ParseUint(ids[j], 0, 64)
		return a < b
	})

	return fmt.Sprintf("%s_%s", ids[0], ids[1])
}
