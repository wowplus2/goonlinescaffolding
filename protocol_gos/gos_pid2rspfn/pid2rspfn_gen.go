// Code generated by "genprotocol -ver=68269c8bfeecd6e461aa862d64007f60a1aeccd64229d523ae7b99a446255112 -basedir=. -prefix=gos -statstype=int"

package gos_pid2rspfn

import (
	"fmt"
	"sync"

	"github.com/kasworld/goonlinescaffolding/protocol_gos/gos_packet"
)

type HandleRspFn func(gos_packet.Header, interface{}) error
type PID2RspFn struct {
	mutex      sync.Mutex
	pid2recvfn map[uint32]HandleRspFn
	pid        uint32
}

func New() *PID2RspFn {
	rtn := &PID2RspFn{
		pid2recvfn: make(map[uint32]HandleRspFn),
	}
	return rtn
}
func (p2r *PID2RspFn) NewPID(fn HandleRspFn) uint32 {
	p2r.mutex.Lock()
	defer p2r.mutex.Unlock()
	p2r.pid++
	p2r.pid2recvfn[p2r.pid] = fn
	return p2r.pid
}
func (p2r *PID2RspFn) HandleRsp(header gos_packet.Header, body interface{}) error {
	p2r.mutex.Lock()
	if recvfn, exist := p2r.pid2recvfn[header.ID]; exist {
		delete(p2r.pid2recvfn, header.ID)
		p2r.mutex.Unlock()
		return recvfn(header, body)
	}
	p2r.mutex.Unlock()
	return fmt.Errorf("pid not found")
}