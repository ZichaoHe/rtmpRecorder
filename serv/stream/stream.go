package stream

import (
	"rtmpRecorder/serv/lib/gortmp"
)

func init() {

}

func newDefaultStream() *Stream {
	s := new(Stream)
	return s
}

func (p *Stream) writeRoutine() {
	gortmp.Dial("", nil, 100)
}

func (p *Stream) stopRoutine() {

}

type IWrite interface {
	write() error
}
