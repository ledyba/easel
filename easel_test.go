package easel

import (
	"encoding/base64"
	"image"
	_ "image/gif"
	"image/png"
	"math"
	"os"
	"runtime"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// in GIF format.
const ICON = `
R0lGODlhIAAeAOfPAL8ARMsgPcwiPNMyONc+NtlFNtpFNN5MMuBRMeNXMOBaMeVaLuRcL+RdL+BgMepf
LOlhLOVkLulkLO5jKuhmLOhmLeVpLupqLOhrLexrKvFqJ+lsK+ZvLepwK+pyK+V2LvZyJPJ2Ju95KPJ5
Ju56Ket7LP93H+t8K/t5Iv15IP55IOp+K+9+KO1/Kv97H/58IPl+Iu6CKfSBJf9/H/CDKP2BIO6FKf+D
H/iGI/uGIviHI/2GIP2GIe2KKvqHIv+GH/yIIf+IHv+KH/+LH/uNIv6NH/+OH/+PH/+QH/aUJP+SH/+T
Hv+TH/uUIf+UH/OXJv6VIP+VHv+WH/eYJPmYIv+XH/6YIP+YH/6ZIPabJP+ZHvebI/+aHvicI/+bH/qd
Iv+cH/+dH/mfIv+eHv+eH/2fIP+fHv+fH/+gHviiI/miI/+iH/ujIfujIvekJPWlJf+jH/ykIP+kHv+k
H/+lH/enJP+mHv+mH/ynIf+nH/2oIP+oH/ypIfiqI/+pH/+qH/+rHvatJf+rH/+sH/yuIfqvIv+uH/uv
If+vHv+vH/ixI/azJPyyIfqzIv+yHv+yH/W1Jf+zHv+zH/61H/+1Hv+1H/+2Hv23If+3H/64IP+4Hv+4
H/+5Hv+6H/q8Iv+7Hv+7H/m9I/+8H/29IP+9Hv+9H/++H/jAJP+/H//AH/vBIv/BHv/CH//DH//EH//G
H//HH//JH/vKIf/KHv/KH//LHv/LH//MH//OH/zPIf/PH//QH/7SIP/SH//TH//UH/7WIP/WH//XH//Z
H//aH//bH/7cH//cH//dH//eH//fH//iH//jH//mH//qH///////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
/yH5BAEKAP8ALAAAAAAgAB4AAAj+AP8JHEiwoMGDCAkySTgQFkOETKQkNNXK4UODTKB4OQiKosWLBJ3o
8HKmICZQqUZ9BClQCREsZ9YMfIRJVB9XK0EeUYLDSkw6/xJJwsToUkWWA40wwdEk5p1/hoa68fjQREGl
VRLIWZMHqiROT6gyNOFixg2BO7kg2JrnjyFKJ1ZRfHVRRYobP4T8U8LlwBqubtUQ6sgKpIsXIIIMMaLk
igwRgCO1mNRxIK5ewoolY9bMGcEZKDQIOcKkCpkKXAV1sXHSI63Lwo4Z41zQbI0JS5RUMdOAqx8LmlpX
pKXrV65TypgdxAvjge4xFNbsSbKFZsdWtTwtgiQLWbJlCYfQAIGgJQwDOIgKOBoKqhSrD4F2YS6GjOEQ
JD4WgNkwJ0KcRJaQskIPb8xyC2zDIGTWDzwUwYQOElxQhgeDiBGDA5mswgosr/XyEF475BDCCBkMYAAG
CkzBxybXvUILUkMcEQUXaBAQAAlftGEJJ9dxqIuHSAkkQAdU+MEGByUUgspwlwUTJAD/2MHCIAB+ckgd
jSjiiSq+ZBakQIDQoAd7prASiy27hMILMMqAh1QeguCRRSWdzNVhMMR8WVAaoOj5D5R+BioooIIeRGih
iF4UEAA7
`
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
uniform vec2 offset[441];
uniform vec4 kernel[441];
in vec2 uv;
layout(location = 0) out vec4 color;
void main() {
	vec4 sum = vec4(0,0,0,0);
	for (int i = 0; i < 441; i++) {
			sum += texture(tex, uv.st + offset[i]) * kernel[i];
	}
	color = sum;
}
`

func init() {
}

func TestRender(t *testing.T) {
	runtime.LockOSThread()
	glfw.Init()
	defer glfw.Terminate()
	e := NewEasel()
	e.MakeCurrent()
	defer e.Destroy()
	p := e.NewPalette()
	p.Bind()
	defer p.Unbind()
	defer p.Destroy()
	// DO YOUR TEST

	prog, err := e.CompileProgram(VertexShader, FragmentShader)
	if err != nil {
		t.Errorf("Could not compile shader: \n** Message **\n%v", err)
	}
	p.AttachProgram(prog, "tex")
	prog.Use()
	defer prog.Unuse()
	defer prog.Destroy()

	data, _ := base64.StdEncoding.DecodeString(ICON)
	tex, err := e.LoadTexture2D(data)
	if err != nil {
		t.Errorf("Could not create texure: \n** Message **\n%v", err)
	}
	defer tex.Destroy()

	p.vertexArray.bind()
	indecies, err := p.AttachArrayIndexBuffer([]uint16{0, 1, 3, 2, 3, 0})
	if err != nil {
		t.Errorf("Could not bind array indecies: \n** Message **\n%v", err)
	}
	_, err = p.AttachArrayBuffer([]float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		1, 1, 0,
	})

	if err != nil {
		t.Errorf("Could not create texure: \n** Message **\n%v", err)
	}
	err = p.BindArrayAttrib(indecies, "vert", 3, 0, 0)
	if err != nil {
		t.Errorf("Could not bind array attrib: \n** Message **\n%v", err)
	}

	offset := make([]float32, 441*2)
	kernel := make([]float32, 441*4)
	idx := 0
	sumk := float32(0.0)
	for x := -10; x <= 10; x++ {
		for y := -10; y <= 10; y++ {
			offset[idx*2+0] = float32(x) / float32(32)
			offset[idx*2+1] = float32(y) / float32(32)
			fx := float64(x) * math.Pi
			fy := float64(y) * math.Pi
			kx := float64(1)
			if x != 0 {
				kx = 10 * math.Sin(fx) * math.Sin(fx/10.0) / (fx * fx)
			}
			ky := float64(1)
			if y != 0 {
				ky = 10 * math.Sin(fy) * math.Sin(fy/10.0) / (fy * fy)
			}
			k := float32(kx * ky)
			sumk += k
			kernel[idx*4+0] = k
			kernel[idx*4+1] = k
			kernel[idx*4+2] = k
			kernel[idx*4+3] = k
			idx++
		}
	}
	if sumk != 1.0 {
		t.Errorf("sumk should be 1, but %f", sumk)
	}

	err = p.BindUniformf("offset", 2, offset)
	if err != nil {
		t.Error(err)
	}
	err = p.BindUniformf("kernel", 4, kernel)
	if err != nil {
		t.Error(err)
	}

	img, err := p.Render(indecies, tex, image.Rect(0, 0, 256, 256))
	if err != nil {
		t.Errorf("Could not execute: \n** Message **\n%v", err)
	}
	file, err := os.Create("test.png")
	if err != nil {
		t.Error(err)
	}

	png.Encode(file, img)

}
