package shader

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
)

const path = "./internal/shader/ray_tracing.glsl"

type Shader struct {
	Width      int32
	Height     int32
	outTex     uint32
	sizeBuffer uint32
	program    uint32
}

func (s *Shader) Init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	var err error

	if err = glfw.Init(); err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			s.Delete()
		}
	}()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(640, 480, "", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	fmt.Printf("OpenGL VERSION %s\n", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Printf("OpenGL VENDOR %s\n", gl.GoStr(gl.GetString(gl.VENDOR)))
	fmt.Printf("OpenGL RENDERER %s\n", gl.GoStr(gl.GetString(gl.RENDERER)))

	compiledShader, err := compileShader()
	if err != nil {
		panic(err)
	}

	s.program = gl.CreateProgram()

	gl.AttachShader(s.program, compiledShader)
	gl.LinkProgram(s.program)

	var status int32
	gl.GetProgramiv(s.program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(s.program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.program, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(compiledShader)

	gl.GenTextures(1, &s.outTex)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.outTex)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA32F, s.Width, s.Height, 0, gl.RGB, gl.FLOAT, nil)
	gl.BindImageTexture(0, s.outTex, 0, false, 0, gl.WRITE_ONLY, gl.RGBA32F)

	sizeData := []float32{float32(s.Width), float32(s.Height)}

	gl.GenBuffers(1, &s.sizeBuffer)
	gl.BindBuffer(gl.SHADER_STORAGE_BUFFER, s.sizeBuffer)
	gl.BufferData(gl.SHADER_STORAGE_BUFFER, 2*4, gl.Ptr(sizeData), gl.STATIC_READ)
	gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER, 1, s.sizeBuffer)
}

func (s *Shader) Delete() {
	glfw.Terminate()
}

func compileShader() (uint32, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	shader := gl.CreateShader(gl.COMPUTE_SHADER)
	cSrc, free := gl.Strs(string(src) + "\x00")
	defer free()

	gl.ShaderSource(shader, 1, cSrc, nil)
	gl.CompileShader(shader)

	var status int32

	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %s: %s", src, log)
	}

	return shader, nil
}

func (s *Shader) Compute(scene scene.Scene) []float32 {
	gl.UseProgram(s.program)
	gl.DispatchCompute(uint32(s.Width), uint32(s.Height), 1)
	gl.MemoryBarrier(gl.ALL_BARRIER_BITS)

	collectionSize := s.Width * s.Height * 3
	computeData := make([]float32, collectionSize, collectionSize)

	gl.GetTexImage(gl.TEXTURE_2D, 0, gl.RGB, gl.FLOAT, gl.Ptr(computeData))

	return computeData
}
