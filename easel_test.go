package easel

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

const VertexShader = `
#version 410 core
layout(location = 0) in vec3 vert;

void main() {
	gl_Position = vec4(vert, 1);
}
`
const FragmentShader = `
#version 410
layout(location = 0) out vec4 color;

void main() {
	color = vec4(1,0,0,1);
}
`

func init() {
}

func startup() {
	runtime.LockOSThread()
	glfw.Init()
}

func shutdown() {
	glfw.Terminate()
	runtime.UnlockOSThread()
}

func TestCreateEasel(t *testing.T) {
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
}

func TestCreatePalette(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
	p, err := e.NewPalette()
	if err != nil {
		t.Fatalf("Could not creating palette: \n** Message **\n%v", err)
	}
	err = p.Bind()
	if err != nil {
		t.Fatalf("Could not bind palette: \n** Message **\n%v", err)
	}
	defer p.Unbind()
	defer p.Destroy()
}

func TestCompileProgram(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
	p, err := e.CompileProgram(VertexShader, FragmentShader)
	if err != nil {
		t.Fatalf("Failed to compile program. %v", err)
	}
	defer p.Destroy()
	err = p.Use()
	if err != nil {
		t.Fatalf("Failed to use program. %v", err)
	}
	defer p.Unuse()
}

func TestCompileInvalidProgram(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
	_, err = e.CompileProgram("", FragmentShader)
	if err == nil {
		t.Fatalf("Oops. We succeeded to compile empty shader: %v", err)
	}
	_, err = e.CompileProgram(VertexShader, "")
	if err == nil {
		t.Fatalf("Oops. We succeeded to compile empty shader: %v", err)
	}
}

func TestLoadingTexture(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()

	freader, err := os.Open("test-images/momiji.png")
	if err != nil {
		log.Fatal(err)
	}
	defer freader.Close()
	bytes, err := ioutil.ReadAll(freader)
	if err != nil {
		log.Fatal(err)
	}

	tex, _, err := e.LoadTexture2D(bytes)
	if err != nil {
		t.Fatalf("Could not create texure: \n** Message **\n%v", err)
	}
	defer tex.Destroy()
}

func TestFailingLoadingTexture(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()

	_, _, err = e.LoadTexture2D([]byte{})
	if err == nil {
		t.Fatalf("Oops. Succeeded to loading empty binary...")
	}
}

func TestRender(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
	p, err := e.NewPalette()
	if err != nil {
		t.Fatalf("Could not creating palette: \n** Message **\n%v", err)
	}
	err = p.Bind()
	if err != nil {
		t.Fatalf("Could not bind palette: \n** Message **\n%v", err)
	}
	defer p.Unbind()
	defer p.Destroy()
	// DO YOUR TEST

	prog, err := e.CompileProgram(VertexShader, FragmentShader)
	if err != nil {
		t.Fatalf("Could not compile shader: \n** Message **\n%v", err)
	}
	p.AttachProgram(prog)
	prog.Use()
	defer prog.Unuse()
	defer prog.Destroy()

	p.vertexArray.bind()
	indecies, err := p.AttachArrayIndexBuffer([]uint16{0, 1, 3, 2, 3, 0})
	if err != nil {
		t.Fatalf("Could not bind array indecies: \n** Message **\n%v", err)
	}
	buf, err := p.MakeArrayBuffer([]float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		1, 1, 0,
	})

	if err != nil {
		t.Fatalf("Could not create texure: \n** Message **\n%v", err)
	}
	err = p.BindArrayAttrib(buf, indecies, "vert", 3, 0, 0)
	if err != nil {
		t.Fatalf("Could not bind array attrib: \n** Message **\n%v", err)
	}

	img, err := p.Render(image.Rect(0, 0, 256, 256))
	if err != nil {
		t.Fatalf("Could not execute: \n** Message **\n%v", err)
	}

	r, g, b, a := img.At(0, 0).RGBA()
	if r == 65535 && g == 0 && b == 0 && a == 65535 {
	} else {
		t.Errorf("Failed to execute shader. Expected red pixel, but got (%d,%d,%d,%d)", r, g, b, a)
	}
}

func TestBindUniform(t *testing.T) {
	var err error
	startup()
	defer shutdown()
	e := NewEasel()
	e.MakeCurrent()
	defer e.DetachCurrent()
	defer e.Destroy()
	p, err := e.NewPalette()
	if err != nil {
		t.Fatalf("Could not creating palette: \n** Message **\n%v", err)
	}
	err = p.Bind()
	if err != nil {
		t.Fatalf("Could not bind palette: \n** Message **\n%v", err)
	}
	defer p.Unbind()
	defer p.Destroy()
	prog, err := e.CompileProgram(VertexShader, `
		#version 410
		layout(location = 0) out vec4 color;
		uniform vec4 c1;
		uniform ivec4 c2;

		void main() {
			color = c1 * vec4(c2);
		}
`)
	if err != nil {
		t.Fatalf("Failed to compile program. %v", err)
	}
	p.AttachProgram(prog)
	prog.Use()
	defer prog.Unuse()
	defer prog.Destroy()
	p.BindUniformf("c1", 4, []float32{1, 1, 0, 1})
	p.BindUniformi("c2", 4, []int32{0, 1, 0, 1})

	p.vertexArray.bind()
	indecies, err := p.AttachArrayIndexBuffer([]uint16{0, 1, 3, 2, 3, 0})
	if err != nil {
		t.Fatalf("Could not bind array indecies: \n** Message **\n%v", err)
	}
	buf, err := p.MakeArrayBuffer([]float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		1, 1, 0,
	})

	if err != nil {
		t.Fatalf("Could not create texure: \n** Message **\n%v", err)
	}
	err = p.BindArrayAttrib(buf, indecies, "vert", 3, 0, 0)
	if err != nil {
		t.Fatalf("Could not bind array attrib: \n** Message **\n%v", err)
	}

	img, err := p.Render(image.Rect(0, 0, 256, 256))
	if err != nil {
		t.Fatalf("Could not execute: \n** Message **\n%v", err)
	}

	r, g, b, a := img.At(0, 0).RGBA()
	if r == 0 && g == 65535 && b == 0 && a == 65535 {
	} else {
		t.Errorf("Failed to execute shader. Expected green pixel, but got (%d,%d,%d,%d)", r, g, b, a)
	}
}
