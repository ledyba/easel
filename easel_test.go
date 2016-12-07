package easel

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

const VertexShader = `
#version 410 core
layout(location = 0) in vec3 vert;
out vec2 uv;

void main() {
	uv = (vert.xy+vec2(1,1))/2.0;
	gl_Position = vec4(vert, 1);
}
`
const FragmentShader = `
#version 410
uniform sampler2D tex;
uniform vec2 srcSize;
uniform vec2 scale;
in vec2 uv;
layout(location = 0) out vec4 color;
const float kPI = 3.14159265358979323846264338327950288;
const int kLobe = 10;

double L(float x) {
	if (abs(x) <= 0.00000001) {
		return 1;
	}
	float px = kPI * x;
	double r = double(kLobe) / (double(px) * double(px));
	r *= sin(px);
	r *= sin(px / float(kLobe));
	return r;
}

void main() {
  dvec4 sum = vec4(0,0,0,0);
  vec2 srcPt = srcSize * uv;
  vec2 base = floor(srcPt) + vec2(0.5, 0.5);

  vec2 pt;
  dvec4 c;
  double w;
  double sumw = 0;
  vec2 nscale = max(vec2(1,1), 1 / scale);
  vec2 support = kLobe * nscale;
  nscale = 1 / nscale;
  vec2 start = floor(base - support + 0.5);
  ivec2 contributes = ivec2(support * 2);

	double alpha;
  for(int dx = 0; dx < contributes.x; dx++) {
    for(int dy = 0; dy < contributes.y; dy++) {
      pt = start + vec2(dx, dy);
      c = dvec4(texture(tex, pt / srcSize));
      w = L(nscale.x * (pt.x - srcPt.x + 0.5)) * L(nscale.y * (pt.y - srcPt.y + 0.5));
      sumw += w;
      sum += c * w;
    }
  }
  sum /= sumw;
	color = vec4(clamp(sum, dvec4(0,0,0,0), dvec4(1,1,1,1)));
}

`

func init() {
}

func TestRender(t *testing.T) {
	var err error
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	glfw.Init()
	defer glfw.Terminate()
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

	/**** Render Image ****/
	freader, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer freader.Close()
	bytes, err := ioutil.ReadAll(freader)
	if err != nil {
		log.Fatal(err)
	}

	tex, src, err := e.LoadTexture2D(bytes)
	if err != nil {
		t.Fatalf("Could not create texure: \n** Message **\n%v", err)
	}
	defer tex.Destroy()

	p.BindUniformf("srcSize", 2, []float32{
		float32(src.Bounds().Dx()), float32(src.Bounds().Dy())})
	p.BindUniformf("scale", 2, []float32{
		float32(math.Max(1.0, float64(512.0)/float64(src.Bounds().Dx()))),
		float32(math.Max(1.0, float64(512.0)/float64(src.Bounds().Dy()))),
	})

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

	p.BindTexture("tex", tex)

	img, err := p.Render(image.Rect(0, 0, 512, 512))
	if err != nil {
		t.Fatalf("Could not execute: \n** Message **\n%v", err)
	}
	file, err := os.Create("test.out.png")
	if err != nil {
		t.Error(err)
	}

	png.Encode(file, img)

}
