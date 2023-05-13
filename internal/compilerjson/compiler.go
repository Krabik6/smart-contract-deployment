package compilerjson

type Compiler struct {
	WorkDir string
	Image   string
}

func NewCompiler(Workdir, image string) *Compiler {
	return &Compiler{
		WorkDir: Workdir,
		Image:   image,
	}
}
