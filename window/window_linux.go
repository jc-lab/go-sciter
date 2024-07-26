package window

import (
	"fmt"
	"github.com/jc-lab/go-sciter"
)

// Linux/gtk3 must (at least) use sciter.DefaultWindowCreationFlag to create the main window correctly
func New(creationFlags sciter.WindowCreationFlag, rect *sciter.Rect) (*Window, error) {
	w := new(Window)
	w.creationFlags = creationFlags
	// create window
	hwnd := sciter.CreateWindow(
		creationFlags,
		rect,
		0,
		0,
		sciter.BAD_HWINDOW)

	if hwnd == sciter.BAD_HWINDOW {
		return nil, fmt.Errorf("Sciter CreateWindow failed")
	}

	w.Sciter = sciter.Wrap(hwnd)
	return w, nil
}

func (s *Window) SetTitle(title string) {
	//	w := C.gwindow((*C.GtkWidget)(unsafe.Pointer(s.GetHwnd())))
	//	t := C.CString(title)
	//	C.gtk_window_set_title(w, t)
	//	C.free(unsafe.Pointer(t))
}

func (s *Window) AddQuitMenu() {
	// Define behaviour for linux
}
