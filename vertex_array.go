package easel

import "github.com/go-gl/gl/v4.1-core/gl"

// VertexArray ...
type VertexArray struct {
	id uint32
}

func newVertexArray() (*VertexArray, error) {
	var err error
	var vaID uint32
	gl.GenVertexArrays(1, &vaID)
	err = checkGLError("Error on glGenVertexArrays")
	if err != nil {
		return nil, err
	}
	return &VertexArray{
		id: vaID,
	}, nil
}

// Destroy ...
func (va *VertexArray) Destroy() {
	gl.DeleteVertexArrays(1, &va.id)
}

func (va *VertexArray) bind() error {
	gl.BindVertexArray(va.id)
	return checkGLError("Error while binding VertexArray")
}

func (va *VertexArray) unbind() {
	gl.BindVertexArray(0)
}
