package easel

import "github.com/go-gl/gl/v4.1-core/gl"

// VertexBuffer ...
type VertexBuffer struct {
	id     uint32
	target uint32
}

func newVertexArrayBuffer() *VertexBuffer {
	vb := &VertexBuffer{}
	gl.GenBuffers(1, &vb.id)
	vb.target = gl.ARRAY_BUFFER
	return vb
}

func newVertexIndexArrayBuffer() *VertexBuffer {
	vb := &VertexBuffer{}
	gl.GenBuffers(1, &vb.id)
	vb.target = gl.ELEMENT_ARRAY_BUFFER
	return vb
}

func (vb *VertexBuffer) bind() error {
	gl.BindBuffer(vb.target, vb.id)
	return checkGLError("Error while binding vertex buffer")
}

func (vb *VertexBuffer) unbind() {
	gl.BindBuffer(vb.target, 0)
}

func (vb *VertexBuffer) loadDataf(data []float32) error {
	gl.BufferData(vb.target, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	return checkGLError("Error while loading data(int)")
}

func (vb *VertexBuffer) loadDatai(data []uint32) error {
	gl.BufferData(vb.target, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	return checkGLError("Error while loading data(float)")
}
