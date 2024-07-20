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
