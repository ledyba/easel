package easel

import "github.com/go-gl/gl/v4.1-core/gl"

// Easel ...
type Easel struct {
	program     *Program
	vertexArray *VertexArray
	textureName string
}

func newEasel() *Easel {
	return &Easel{
		program:     nil,
		vertexArray: newVertexArray(),
	}
}

func (e *Easel) attachProgram(p *Program) {
	e.program = p
}

func (e *Easel) bindArrayAttrib(vb *VertexBuffer, name string, size, stride, offset int32) error {
	idx, err := e.program.attibLocation(name)
	if err != nil {
		return err
	}
	gl.EnableVertexAttribArray(idx)
	vb.bind()
	defer vb.unbind()
	gl.VertexAttribPointer(idx, size, gl.FLOAT, false, stride, gl.PtrOffset(int(offset)))
	return checkGLError("Error while binding array attrib.")
}

func (e *Easel) attachArrayBuffer(data []float32) (*VertexBuffer, error) {
	var err error
	buff := newVertexArrayBuffer()
	buff.bind()
	defer buff.unbind()
	err = buff.loadDataf(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func (e *Easel) attachArrayIndexBuffer(data []uint32) (*VertexBuffer, error) {
	var err error
	buff := newVertexIndexArrayBuffer()
	buff.bind()
	defer buff.unbind()
	err = buff.loadDatai(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

// Run ...
func (e *Easel) Run(tex *Texture2D) error {
	var err error
	if err = e.program.use(); err != nil {
		return err
	}
	defer e.program.unuse()
	if err = e.vertexArray.bind(); err != nil {
		return err
	}
	defer e.vertexArray.unbind()
	gl.ActiveTexture(gl.TEXTURE0)
	if err = checkGLError("Error while activating texture 0"); err != nil {
		return err
	}
	tex.bind()
	defer tex.unbind()
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

	return nil
}
