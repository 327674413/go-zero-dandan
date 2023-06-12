package util

import (
	"fmt"
	"github.com/sony/sonyflake"
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

func MakeId() uint64 {
	ret, err := sf.NextID()
	if err != nil {
		ret = uint64(Int63nRange(1926425572, 1926425572223607))
	}
	return ret
}
