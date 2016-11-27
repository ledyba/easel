package easel

import (
	"errors"
	"fmt"
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Palette ...
type Palette struct {
	easel         *Easel
	program       *Program
	vertexArray   *VertexArray
	frameBufferID uint32
	textureName   string
}

func newPalette(e *Easel) *Palette {
	e.MakeCurrent()
	p := &Palette{
		easel:       e,
		program:     nil,
		vertexArray: newVertexArray(),
	}
	gl.GenFramebuffers(1, &p.frameBufferID)
	return p
}

// Bind ...
func (p *Palette) Bind() error {
	var err error
	err = p.vertexArray.bind()
	if err != nil {
		return err
	}
	return nil
}

// Unbind ...
func (p *Palette) Unbind() {
	p.vertexArray.unbind()
}

// Destroy ...
func (p *Palette) Destroy() {
	p.vertexArray.Destroy()
	gl.DeleteFramebuffers(1, &p.frameBufferID)
}

// AttachProgram ...
func (p *Palette) AttachProgram(prog *Program, textureName string) {
	p.program = prog
	p.textureName = textureName
}

// Program ...
func (p *Palette) Program() *Program {
	return p.program
}

// BindArrayAttrib ...
func (p *Palette) BindArrayAttrib(vb *VertexBuffer, name string, size, stride, offset int32) error {
	var err error
	idx, err := p.program.attibLocation(name)
	if err != nil {
		return err
	}
	err = vb.bind()
	if err != nil {
		return err
	}
	gl.EnableVertexAttribArray(uint32(idx))
	err = checkGLError(fmt.Sprintf("Error while enabling vertex attrib array (location: %d)", idx))
	if err != nil {
		return err
	}
	gl.VertexAttribPointer(uint32(idx), size, gl.FLOAT, false, stride, gl.PtrOffset(int(offset)))
	return checkGLError("Error while binding array attrib.")
}

// BindUniformf ...
func (p *Palette) BindUniformf(name string, vecDim int, data []float32) error {
	var err error
	loc, err := p.program.uniformLocation(name)
	if err != nil {
		return err
	}
	switch vecDim {
	case 1:
		gl.Uniform1fv(loc, int32(len(data)), &data[0])
		return checkGLError("Error on glUniform1fv")
	case 2:
		gl.Uniform2fv(loc, int32(len(data)/2), &data[0])
		return checkGLError("Error on glUniform2fv")
	case 3:
		gl.Uniform3fv(loc, int32(len(data)/3), &data[0])
		return checkGLError("Error on glUniform3fv")
	case 4:
		gl.Uniform4fv(loc, int32(len(data)/4), &data[0])
		return checkGLError("Error on glUniform4fv")
	default:
		return fmt.Errorf("Unsupported vector dimension: %d", vecDim)
	}
}

// AttachArrayBuffer ...
func (p *Palette) AttachArrayBuffer(data []float32) (*VertexBuffer, error) {
	var err error
	buff := newVertexArrayBuffer()
	err = buff.bind()
	if err != nil {
		return nil, err
	}
	err = buff.loadDataf(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

// AttachArrayIndexBuffer ...
func (p *Palette) AttachArrayIndexBuffer(data []uint16) (*VertexBuffer, error) {
	var err error
	buff := newVertexIndexArrayBuffer()
	err = buff.bind()
	if err != nil {
		return nil, err
	}
	err = buff.loadDatai(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

// Render ...
func (p *Palette) Render(indecies *VertexBuffer, tex *Texture2D, size image.Rectangle) (image.Image, error) {
	var err error
	var texID uint32
	gl.BindFramebuffer(gl.FRAMEBUFFER, p.frameBufferID)
	err = checkGLError("Error while binding framebuffer")
	if err != nil {
		return nil, err
	}

	if err = checkGLError("Error while binding framebuffer"); err != nil {
		return nil, err
	}
	/* Setup Texture for FrameBuffer */
	gl.GenTextures(1, &texID)
	if err = checkGLError("Error while generating framebuffer texture"); err != nil {
		return nil, err
	}
	gl.BindTexture(gl.TEXTURE_2D, texID)
	if err = checkGLError("Error while binding framebuffer texture"); err != nil {
		return nil, err
	}
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nil))
	if err = checkGLError("Error while creating empty framebuffer texture"); err != nil {
		return nil, err
	}
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	if err = checkGLError("Error while setting framebuffer texture parameter"); err != nil {
		return nil, err
	}
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, texID, 0)
	if err = checkGLError("Error while attaching framebuffer texture"); err != nil {
		return nil, err
	}
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return nil, errors.New("Invalid Framebuffer Status")
	}
	gl.Viewport(0, 0, int32(size.Dx()), int32(size.Dy()))
	if err = checkGLError("Error while set viewport"); err != nil {
		return nil, err
	}
	/* Start rendering */
	if err = p.program.Use(); err != nil {
		return nil, err
	}
	defer p.program.Unuse()
	if err = p.vertexArray.bind(); err != nil {
		return nil, err
	}
	gl.ActiveTexture(gl.TEXTURE0)
	if err = tex.bind(); err != nil {
		return nil, err
	}
	defer tex.unbind()
	textureLoc, err := p.program.uniformLocation(p.textureName)
	if err != nil {
		return nil, err
	}
	gl.Uniform1i(int32(textureLoc), 0) // We use texture 0

	p.vertexArray.bind()
	defer p.vertexArray.unbind()

	gl.DrawElements(gl.TRIANGLES, int32(indecies.length), gl.UNSIGNED_SHORT, gl.Ptr(nil))

	if err = checkGLError("Error on DrawArrays"); err != nil {
		return nil, err
	}
	//e.easel.SwapBuffers()

	/* Readback the output */
	gl.BindTexture(gl.TEXTURE_2D, texID)
	if err = checkGLError("Error on bind framebuffer texture"); err != nil {
		return nil, err
	}
	out := image.NewRGBA(size)

	gl.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(out.Pix))
	if err = checkGLError("Error on GetTexImage"); err != nil {
		return nil, err
	}
	gl.DeleteTextures(1, &texID)
	if err = checkGLError("Error on DeleteTextures"); err != nil {
		return nil, err
	}

	return out, nil
}
