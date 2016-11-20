package easel

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Studio ...
type Studio struct {
	window *glfw.Window
}

// NewStudio ...
func NewStudio() *Studio {
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w, err := glfw.CreateWindow(640, 480, "Studio", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	w.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Studio Created.")
	log.Infof("  ** OpenGL Info **")
	log.Infof("    OpenGL Version: %s", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Infof("    GLSL Version:   %s", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	log.Infof("    OpenGL Vendor:  %s", gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Infof("    Renderer:       %s", gl.GoStr(gl.GetString(gl.RENDERER)))
	log.Infof("    ** Extensions **")
	for i := uint32(0); i < gl.NUM_EXTENSIONS; i++ {
		str := gl.GetStringi(gl.EXTENSIONS, i)
		if str != nil {
			log.Infof("      - %s", gl.GoStr(str))
		}
	}

	log.Debug("Studio Created.")
	return &Studio{
		window: w,
	}
}

// Destroy ...
func (s *Studio) Destroy() {
	s.window.Destroy()
}

// MakeCurrent ...
func (s *Studio) MakeCurrent() {
	s.window.MakeContextCurrent()
}

// MakeEasel ...
func (s *Studio) MakeEasel() *Easel {
	return newEasel(s)
}

// SwapBuffers ...
func (s *Studio) SwapBuffers() {
	s.window.SwapBuffers()
}

// CompileProgram ...
func (s *Studio) CompileProgram(vertex, fragment string) (*Program, error) {
	progID, err := compileProgram(vertex, fragment)
	if err != nil {
		return nil, err
	}
	return newProgram(progID), nil
}

// LoadTexture2D ...
func (s *Studio) LoadTexture2D(data []byte) (*Texture2D, error) {
	return newTexture2D(data)
}

// Destroy ...
func (p *Program) Destroy() {
	gl.DeleteProgram(p.progID)
}
