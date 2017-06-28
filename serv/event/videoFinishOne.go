package event

type VideoReactor struct {
	ch chan interface{}
}

func init() {

}

func newDeafultVideoReactor() *VideoReactor {
	r := new(VideoReactor)
	r.ch = make(chan interface{}, 1024)
	return r
}

func (p *VideoReactor) VideoReactorRoutine() {
	for {
		_ = <-p.ch
		switch {

		}
	}
}
