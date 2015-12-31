//
//   date  : 2014-06-04
//   author: xjdrew
//

package tunnel

// tunnel read/write timeout
const (
	PacketSize = 524288
)

var (
	Timeout  int64 = 3
	LogLevel uint  = 1
	mpool          = NewMPool(PacketSize)
)

type Service interface {
	Start() error
	Status()
}
