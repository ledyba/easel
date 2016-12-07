package filters

import (
	"image"

	log "github.com/Sirupsen/logrus"

	"github.com/ledyba/easel/proto"
	"golang.org/x/net/context"
)

const (
	// LanczosFilter ...
	LanczosFilter = "lanczos"
)

// UpdateLanczos ...
func UpdateLanczos(serv proto.EaselServiceClient, easelID, paletteID string, lobes int) error {
	var err error
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

  uniform int lobes;
  layout(location = 0) out vec4 color;
  const float kPI = 3.14159265358979323846264338327950288;

  double L(float x) {
  	if (abs(x) <= 0.00000001) {
  		return 1;
  	}
  	float px = kPI * x;
  	double r = double(lobes) / (double(px) * double(px));
  	r *= sin(px);
  	r *= sin(px / float(lobes));
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
    vec2 support = lobes * nscale;
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
	/**** Update Palette ****/
	updates := &proto.PaletteUpdate{}
	updates.FragmentShader = FragmentShader
	updates.VertexShader = VertexShader

	updates.Buffers = []*proto.ArrayBuffer{
		&proto.ArrayBuffer{
			Name: "tex",
			Data: []float32{
				-1, -1, 0,
				1, -1, 0,
				-1, 1, 0,
				1, 1, 0,
			},
		}}
	updates.Indecies = []int32{0, 1, 3, 2, 3, 0}
	updates.VertexArrtibutes = []*proto.VertexAttribute{
		&proto.VertexAttribute{
			ArgumentName: "vert",
			BufferName:   "tex",
			ElementSize:  3,
			Offset:       0,
			Stride:       0,
		}}

	uniforms := make([]*proto.UniformVariable, 0)
	uniforms = append(uniforms, &proto.UniformVariable{
		Name: "lobes",
		IntValue: &proto.UniformIntValue{
			ElementSize: 1,
			Data:        []int32{int32(lobes)},
		},
	})
	updates.UniformVariables = uniforms

	_, err = serv.UpdatePalette(context.Background(), &proto.UpdatePaletteRequest{
		EaselId:   easelID,
		PaletteId: paletteID,
		Updates:   updates,
	})
	log.Debugf("Palette Updated (Lanzcos Filter lobes=%d): (%s > %s) err=%v", lobes, easelID, paletteID, err)
	return err
}

// RenderLanczos ...
func RenderLanczos(serv proto.EaselServiceClient, easelID, paletteID string, data []byte, src image.Image, width, height int) ([]byte, error) {
	resp, err := serv.Render(context.Background(), &proto.RenderRequest{
		EaselId:    easelID,
		PaletteId:  paletteID,
		OutFormat:  "image/png",
		OutQuality: 95,
		OutWidth:   int32(width),
		OutHeight:  int32(height),
		Updates: &proto.PaletteUpdate{
			UniformVariables: []*proto.UniformVariable{
				&proto.UniformVariable{
					Name:    "tex",
					Texture: data,
				},
				&proto.UniformVariable{
					Name: "scale",
					FloatValue: &proto.UniformFloatValue{
						ElementSize: 2,
						Data: []float32{
							float32(width) / float32(src.Bounds().Dx()),
							float32(height) / float32(src.Bounds().Dy())}}},
				&proto.UniformVariable{
					Name: "srcSize",
					FloatValue: &proto.UniformFloatValue{
						ElementSize: 2,
						Data: []float32{
							float32(src.Bounds().Dx()), float32(src.Bounds().Dy())}}}}}})
	if err != nil {
		return nil, err
	}
	return resp.Output, nil
}
