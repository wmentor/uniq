package uniq

import (
	"fmt"
	"hash/crc32"
	"os"
	"sync/atomic"
	"time"
)

const (
	version string = "a"
)

var (
	nextId uint64 = uint64(time.Now().UnixNano() % 0xffff)
	pid    int    = os.Getpid()
	hcode  uint32
)

func init() {
	host, _ := os.Hostname()
	hcode = crc32.ChecksumIEEE([]byte(host))
}

func New() string {

	now := time.Now().UnixNano()

	nid := atomic.AddUint64(&nextId, 1) & 0xffff

	return fmt.Sprintf("%s%08x%04x%016x%04x", version, hcode, pid, now, nid)
}
