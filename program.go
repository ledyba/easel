package easel

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Program ...
type Program struct {
	progID uint32
}

func newProgram(progID uint32) *Program {
	return &Program{
		progID: progID,
	}
}

// Destroy ...
func (p *Program) Destroy() error {
	gl.DeleteProgram(p.progID)
	p.progID = 0
	return checkGLError("Error while deleting program")
}

// Use ...
func (p *Program) Use() error {
	gl.UseProgram(p.progID)
	return checkGLError("Error while binding program")
}

// Unuse ...
func (p *Program) Unuse() {
	gl.UseProgram(0)
}

func (p *Program) attibLocation(name string) (int32, error) {
	idx := gl.GetAttribLocation(p.progID, gl.Str(name+"\x00"))
	err := checkGLError("error while get attrib location")
	if idx < 0 {
		return -1, fmt.Errorf("Attribute not found: %s", name)
	}
	return idx, err
}

func (p *Program) uniformLocation(name string) (int32, error) {
	idx := gl.GetUniformLocation(p.progID, gl.Str(name+"\x00"))
	err := checkGLError("error while get uniform location")
	if idx < 0 {
		return -1, fmt.Errorf("Attribute not found: %s", name)
	}
	return idx, err
}

// -----------------------------------------------------------------------------

func compileProgram(vertex, fragment string) (uint32, error) {
	var err error
	vsh, err := compileShader(vertex, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fsh, err := compileShader(fragment, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vsh)
	gl.AttachShader(prog, fsh)
	gl.LinkProgram(prog)

	var status int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(prog, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(prog, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vsh)
	gl.DeleteShader(fsh)
	return prog, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}
