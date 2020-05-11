// Code generated by "genprotocol -ver=68269c8bfeecd6e461aa862d64007f60a1aeccd64229d523ae7b99a446255112 -basedir=. -prefix=gos -statstype=int"

package gos_const

import "time"

const (
	// MaxBodyLen set to max body len, affect send/recv buffer size
	MaxBodyLen = 0xfffff
	// PacketBufferPoolSize max size of pool packet buffer
	PacketBufferPoolSize = 100

	// ServerAPICallTimeOutDur api call watchdog timer
	ServerAPICallTimeOutDur = time.Second * 2
)