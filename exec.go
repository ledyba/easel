package easel

import "github.com/go-gl/gl/v4.1-core/gl"

// Exec ...
type Exec struct {
	program     *Program
	vertexArray *VertexArray
	textureName string
}

func newExec() *Exec {
	return &Exec{
		program:     nil,
		vertexArray: newVertexArray(),
	}
}

func (e *Exec) attachProgram(p *Program) {
	e.program = p
}

func (e *Exec) attachBuffer(buff *VertexBuffer) {
}

// Run ...
func (e *Exec) Run(tex *Texture2D) error {
	var err error
	if err = e.program.use(); err != nil {
		return err
	}
	defer e.program.unuse()
	e.vertexArray.Bind()
	if err = e.vertexArray.Bind(); err != nil {
		return err
	}
	defer e.vertexArray.Unbind()
	gl.ActiveTexture(gl.TEXTURE0)
	if err = checkGLError("Error while activating texture 0"); err != nil {
		return err
	}
	tex.Bind()
	defer tex.Unbind()
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

	return nil
}
