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

func newVertexElementArrayBuffer() *VertexBuffer {
	vb := &VertexBuffer{}
	gl.GenBuffers(1, &vb.id)
	vb.target = gl.ELEMENT_ARRAY_BUFFER
	return vb
}

// Bind ...
func (vb *VertexBuffer) Bind() {
	gl.BindBuffer(vb.target, vb.id)
}

// Unbind ...
func (vb *VertexBuffer) Unbind() {
	gl.BindBuffer(vb.target, 0)
}

// LoadData ...
func (vb *VertexBuffer) LoadData(data []float32) {
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
}
