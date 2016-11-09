package easel

import (
	"encoding/base64"
	_ "image/gif"
	"runtime"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

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
#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
		fragTexCoord = vertTexCoord;
		gl_Position = projection * camera * model * vec4(vert, 1);
}
`
const FragmentShader = `
#version 330
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
		outputColor = texture(tex, fragTexCoord);
}
`

func setup() {
	runtime.LockOSThread()
	glfw.Init()
}
func end() {
	glfw.Terminate()
}

func TestCreate(t *testing.T) {
	setup()
	defer end()
	e := NewEasel()
	e.Destroy()
}

func TestCompileShader(t *testing.T) {
	setup()
	defer end()
	e := NewEasel()
	_, err := e.CompileProgram(VertexShader, FragmentShader)
	if err != nil {
		t.Errorf("Could not compile shader: \n** Message **\n%v", err)
	}
	e.Destroy()
}

func TestCreateTexture(t *testing.T) {
	setup()
	defer end()
	// DO YOUR TEST
	data, _ := base64.StdEncoding.DecodeString(ICON)
	e := NewEasel()
	_, err := e.CreateTexture2D(data)
	if err != nil {
		t.Errorf("Could not create texure: \n** Message **\n%v", err)
	}
	e.Destroy()
}

func TestDrawTexture(t *testing.T) {
	setup()
	defer end()
	// DO YOUR TEST
	data, _ := base64.StdEncoding.DecodeString(ICON)
	e := NewEasel()
	_, err := e.CreateTexture2D(data)
	if err != nil {
		t.Errorf("Could not create texure: \n** Message **\n%v", err)
	}
	e.Destroy()
}
