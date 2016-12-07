package easel

import (
	"errors"
	"fmt"
	"image"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Palette ...
type Palette struct {
	easel         *Easel
	program       *Program
	vertexArray   *VertexArray
	indecies      *VertexBuffer
	frameBufferID uint32
	textureUnits  [gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS - 1]struct {
		name string
		tex  *Texture2D
	}
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
func (p *Palette) AttachProgram(prog *Program) {
	p.program = prog
}

// Program ...
func (p *Palette) Program() *Program {
	return p.program
}

// BindArrayAttrib ...
func (p *Palette) BindArrayAttrib(vertexbuffer *VertexBuffer, indecies *VertexBuffer, name string, size, stride, offset int32) error {
	var err error
	idx, err := p.program.attibLocation(name)
	if err != nil {
		return err
	}
	err = vertexbuffer.Bind()
	if err != nil {
		return err
	}
	err = indecies.Bind()
	if err != nil {
		return err
	}
	gl.EnableVertexAttribArray(uint32(idx))
	err = checkGLError(fmt.Sprintf("Error on enabling vertex attrib array (location: %d)", idx))
	if err != nil {
		return err
	}
	gl.VertexAttribPointer(uint32(idx), size, gl.FLOAT, false, stride, gl.PtrOffset(int(offset)))
	return checkGLError("Error on binding array attrib.")
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

// BindUniformi ...
func (p *Palette) BindUniformi(name string, vecDim int, data []int32) error {
	var err error
	loc, err := p.program.uniformLocation(name)
	if err != nil {
		return err
	}
	switch vecDim {
	case 1:
		gl.Uniform1iv(loc, int32(len(data)), &data[0])
		return checkGLError("Error on glUniform1fv")
	case 2:
		gl.Uniform2iv(loc, int32(len(data)/2), &data[0])
		return checkGLError("Error on glUniform2fv")
	case 3:
		gl.Uniform3iv(loc, int32(len(data)/3), &data[0])
		return checkGLError("Error on glUniform3fv")
	case 4:
		gl.Uniform4iv(loc, int32(len(data)/4), &data[0])
		return checkGLError("Error on glUniform4fv")
	default:
		return fmt.Errorf("Unsupported vector dimension: %d", vecDim)
	}
}

// BindTexture ...
func (p *Palette) BindTexture(name string, tex *Texture2D) (*Texture2D, error) {
	var err error
	idx := 0
	for i, unit := range p.textureUnits {
		idx = i
		if unit.name == name || len(unit.name) == 0 {
			break
		}
	}
	if idx >= len(p.textureUnits) {
		return nil, fmt.Errorf("Texture units limit succeeded: %d", len(p.textureUnits))
	}
	log.Debugf("%s assigned to TextureUnit %d", name, idx+1)
	gl.ActiveTexture(gl.TEXTURE1 + uint32(idx))
	if err = checkGLError(fmt.Sprintf("Error on activate texture unit %d", idx+1)); err != nil {
		return nil, err
	}
	err = tex.bind()
	if err != nil {
		return nil, err
	}
	old := p.textureUnits[idx].tex
	p.textureUnits[idx].name = name
	p.textureUnits[idx].tex = tex

	loc, err := p.program.uniformLocation(name)
	if err != nil {
		return old, err
	}
	gl.Uniform1i(loc, int32(idx+1))
	return old, nil
}

// MakeArrayBuffer ...
func (p *Palette) MakeArrayBuffer(data []float32) (*VertexBuffer, error) {
	var err error
	buff := newVertexArrayBuffer()
	err = buff.Bind()
	if err != nil {
		return nil, err
	}
	defer buff.Unbind()
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
	err = buff.Bind()
	if err != nil {
		return nil, err
	}
	err = buff.loadDatai(data)
	if err != nil {
		return nil, err
	}
	p.indecies = buff
	return buff, nil
}

// Render ...
func (p *Palette) Render(size image.Rectangle) (image.Image, error) {
	var err error
	var texID uint32
	gl.BindFramebuffer(gl.FRAMEBUFFER, p.frameBufferID)
	err = checkGLError("Error on binding framebuffer")
	if err != nil {
		return nil, err
	}

	if err = checkGLError("Error on binding framebuffer"); err != nil {
		return nil, err
	}
	/* Setup Texture for FrameBuffer */
	gl.ActiveTexture(gl.TEXTURE0)
	if err = checkGLError("Error on activate texture unit 0"); err != nil {
		return nil, err
	}
	gl.GenTextures(1, &texID)
	if err = checkGLError("Error on generating framebuffer texture"); err != nil {
		return nil, err
	}
	gl.BindTexture(gl.TEXTURE_2D, texID)
	if err = checkGLError(fmt.Sprintf("Error on binding framebuffer texture (%d)", texID)); err != nil {
		return nil, err
	}
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nil))
	if err = checkGLError("Error on creating empty framebuffer texture"); err != nil {
		return nil, err
	}
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	if err = checkGLError("Error on setting framebuffer texture parameter"); err != nil {
		return nil, err
	}
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, texID, 0)
	if err = checkGLError("Error on attaching framebuffer texture"); err != nil {
		return nil, err
	}
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return nil, errors.New("Invalid Framebuffer Status")
	}
	gl.Viewport(0, 0, int32(size.Dx()), int32(size.Dy()))
	if err = checkGLError("Error on setting viewport"); err != nil {
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
	defer p.vertexArray.unbind()

	gl.Enable(gl.BLEND)
	if err = checkGLError("Error on enabling glBlend"); err != nil {
		return nil, err
	}
	gl.BlendFunc(gl.ONE, gl.ZERO)
	if err = checkGLError("Error on set blend func"); err != nil {
		return nil, err
	}

	gl.DrawElements(gl.TRIANGLES, int32(p.indecies.length), gl.UNSIGNED_SHORT, gl.Ptr(nil))
	if err = checkGLError("Error on DrawArrays"); err != nil {
		return nil, err
	}
	//e.easel.SwapBuffers()

	/* Readback the output */
	gl.ActiveTexture(gl.TEXTURE0)
	if err = checkGLError("Error on activate texture unit 0"); err != nil {
		return nil, err
	}
	gl.BindTexture(gl.TEXTURE_2D, texID)
	if err = checkGLError(fmt.Sprintf("Error on binding framebuffer texture (%d)", texID)); err != nil {
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
