package stream

import (
	"container/list"
	"rtmpRecorder/serv/event"
	"sync"
)

type StreamsMngr struct {
	Mtx     sync.Mutex
	Streams map[string]Stream

	Ract event.VideoReactor
}

type Stream struct {
	SrcUrl  string
	DstPath string
	Videos  list.List
	Status  int

	SigChan chan int
}

type Video struct {
	Name   string
	Size   int64
	Len    int
	Format string
}
