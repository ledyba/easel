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
	defer glfw.DetachCurrentContext()
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Easel Created.")
	log.Infof("  ** OpenGL Info **")
	log.Infof("    OpenGL Version: %s", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Infof("    GLSL Version:   %s", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	log.Infof("    OpenGL Vendor:  %s", gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Infof("    Renderer:       %s", gl.GoStr(gl.GetString(gl.RENDERER)))
	log.Infof("    ** Extensions **")
	for i := uint32(0); i < gl.NUM_EXTENSIONS; i++ {
		str := gl.GetStringi(gl.EXTENSIONS, i)
		code := gl.GetError()
		if code == gl.INVALID_VALUE {
			break
		}
		if str != nil {
			log.Infof("      - %s", gl.GoStr(str))
		}
	}

	log.Debug("Easel Created.")
	return &Easel{
		window: w,
	}
}

// Destroy ...
func (e *Easel) Destroy() {
	e.window.Destroy()
}

// MakeCurrent ...
func (e *Easel) MakeCurrent() {
	e.window.MakeContextCurrent()
}

// DetachCurrent ...
func (e *Easel) DetachCurrent() {
	glfw.DetachCurrentContext()
}

// NewPalette ...
func (e *Easel) NewPalette() (*Palette, error) {
	var err error
	va, err := newVertexArray()
	if err != nil {
		return nil, err
	}
	var fb uint32
	gl.GenFramebuffers(1, &fb)
	err = checkGLError("Error while generating framebuffer")
	if err != nil {
		return nil, err
	}
	p := &Palette{
		easel:         e,
		program:       nil,
		vertexArray:   va,
		frameBufferID: fb,
	}
	return p, nil
}

// SwapBuffers ...
func (e *Easel) SwapBuffers() {
	e.window.SwapBuffers()
}

// CompileProgram ...
func (e *Easel) CompileProgram(vertex, fragment string) (*Program, error) {
	progID, err := compileProgram(vertex, fragment)
	if err != nil {
		return nil, err
	}
	return newProgram(progID), nil
}

// LoadTexture2D ...
func (e *Easel) LoadTexture2D(data []byte) (*Texture2D, error) {
	return newTexture2D(data)
}
