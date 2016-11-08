package easel

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Easel ...
type Easel struct {
	window *glfw.Window
}

// NewEasel ...
func NewEasel() *Easel {
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w, err := glfw.CreateWindow(640, 480, "Easel", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	w.MakeContextCurrent()
	log.Debug("Easel Created.")
	return &Easel{
		window: w,
	}
}

// Destroy ...
func (e *Easel) Destroy() {
	e.window.Destroy()
}

// CompileProgram ...
func (e *Easel) CompileProgram(vertex, fragment string) (*Program, error) {
	progID, err := compileProgram(vertex, fragment)
	if err != nil {
		return nil, err
	}
	return newProgram(e, progID), nil
}

// Destroy ...
func (p *Program) Destroy() {
	gl.DeleteProgram(p.progID)
}
