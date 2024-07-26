package window

import (
	"github.com/jc-lab/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (s *Window) Run() {
	sciter.GetApi().SciterExec(sciter.SCITER_APP_LOOP, 0, 0)
}

func (s *Window) Show() {
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_SET_STATE, sciter.UINT_PTR(sciter.SCITER_WINDOW_STATE_SHOWN), 0)
}

func (s *Window) Collapse() {
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_SET_STATE, sciter.UINT_PTR(sciter.SCITER_WINDOW_STATE_MINIMIZED), 0)
}

func (s *Window) Expand(maximize bool) {
	state := sciter.SCITER_WINDOW_STATE_SHOWN
	if maximize {
		state = sciter.SCITER_WINDOW_STATE_MAXIMIZED
	}
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_SET_STATE, sciter.UINT_PTR(state), 0)
}

func (s *Window) RequestClose() {
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_SET_STATE, sciter.UINT_PTR(sciter.SCITER_WINDOW_STATE_CLOSED), sciter.UINT_PTR(sciter.FALSE))
}

func (s *Window) Close() {
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_SET_STATE, sciter.UINT_PTR(sciter.SCITER_WINDOW_STATE_CLOSED), sciter.UINT_PTR(sciter.TRUE))
}

func (s *Window) Activate(bringToFront bool) {
	p1 := sciter.UINT_PTR(sciter.FALSE)
	if bringToFront {
		p1 = sciter.UINT_PTR(sciter.TRUE)
	}
	sciter.GetApi().SciterWindowExec(s.GetHwnd(), sciter.SCITER_WINDOW_ACTIVATE, p1, 0)
}
