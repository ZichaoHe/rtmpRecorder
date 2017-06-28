package stream

var GlbSm *StreamsMngr

func init() {
	GlbSm = new(StreamsMngr)
}

func (p *StreamsMngr) StartRecord(src string, dst string, t string) {
	s := newDefaultStream()
	p.addToMngr(s)
	go s.writeRoutine()
}

func (p *StreamsMngr) addToMngr(s *Stream) {

}

func (p *StreamsMngr) StopRecord(src string) {
	s := newDefaultStream()
	p.delFromMngr(s)
	s.stopRoutine()
}

func (p *StreamsMngr) delFromMngr(s *Stream) {

}
