package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ledyba/easel/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var server *string = flag.String("server", "localhost:3000", "server to connect")
var help *bool = flag.Bool("help", false, "Print help and exit")

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
	%s [OPTIONS] FILES...
Options:
`, os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	args := flag.Args()
	printStartupBanner()
	if len(args) <= 0 || *help {
		usage()
		return
	}
	var err error
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	serv := proto.NewEaselServiceClient(conn)
	eresp, err := serv.NewEasel(context.Background(), &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	/**** Create Easel ****/
	log.Printf("Easel Created: %s", eresp.EaselId)
	defer func() {
		serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: eresp.EaselId,
		})
		log.Printf("Easel Deleted: %s", eresp.EaselId)
	}()

	/**** Create Palette ****/
	presp, err := serv.NewPalette(context.Background(), &proto.NewPaletteRequest{
		EaselId: eresp.EaselId,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Palette Created: (%s > %s)", presp.EaselId, presp.PaletteId)
	defer func() {
		serv.DeletePalette(context.Background(), &proto.DeletePaletteRequest{
			EaselId:   eresp.EaselId,
			PaletteId: presp.PaletteId,
		})
		log.Printf("Palette Deleted: (%s > %s)", presp.EaselId, presp.PaletteId)
	}()

	/**** Update Palette ****/
	updates := &proto.PaletteUpdate{}
	updates.FragmentShader = `
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
	updates.VertexShader = `
	#version 410 core
	layout(location = 0) in vec3 vert;
	out vec2 uv;

	void main() {
		uv = (vert.xy+vec2(1,1))/2.0;
		gl_Position = vec4(vert, 1);
	}
`
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
		log.Fatalf("sumk should be 1, but %f", sumk)
	}

	uniforms := make([]*proto.UniformVariable, 0)
	uniforms = append(uniforms, &proto.UniformVariable{
		Name: "offset",
		FloatValue: &proto.UniformFloatValue{
			ElementSize: 2,
			Data:        offset,
		},
	})
	uniforms = append(uniforms, &proto.UniformVariable{
		Name: "kernel",
		FloatValue: &proto.UniformFloatValue{
			ElementSize: 4,
			Data:        kernel,
		},
	})
	updates.UniformVariables = uniforms

	_, err = serv.UpdatePalette(context.Background(), &proto.UpdatePaletteRequest{
		EaselId:   presp.EaselId,
		PaletteId: presp.PaletteId,
		Updates:   updates,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Palette Updated: (%s > %s)", presp.EaselId, presp.PaletteId)

	/**** Render Image ****/
	fname := flag.Arg(0)
	freader, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer freader.Close()
	bytes, err := ioutil.ReadAll(freader)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := serv.Render(context.Background(), &proto.RenderRequest{
		EaselId:    presp.EaselId,
		PaletteId:  presp.PaletteId,
		OutFormat:  "image/png",
		OutQuality: 95,
		OutWidth:   128,
		OutHeight:  128,
		Updates: &proto.PaletteUpdate{
			UniformVariables: []*proto.UniformVariable{
				&proto.UniformVariable{
					Name:    "tex",
					Texture: bytes,
				}}}})
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s.out.png", strings.TrimSuffix(fname, path.Ext(fname))), resp.Output, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rendered: (%s > %s) %d bytes", presp.EaselId, presp.PaletteId, len(resp.Output))

}
