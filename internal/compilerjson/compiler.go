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

type SolcOutput struct {
	Errors []struct {
		SourceLocation struct {
			File  string `json:"file"`
			Start int    `json:"start"`
			End   int    `json:"end"`
		} `json:"sourceLocation"`
		SecondarySourceLocations []struct {
			File    string `json:"file"`
			Start   int    `json:"start"`
			End     int    `json:"end"`
			Message string `json:"message"`
		} `json:"secondarySourceLocations"`
		Type             string `json:"type"`
		Component        string `json:"component"`
		Severity         string `json:"severity"`
		ErrorCode        string `json:"errorCode"`
		Message          string `json:"message"`
		FormattedMessage string `json:"formattedMessage"`
	} `json:"errors"`
	Sources map[string]struct {
		Id  int `json:"id"`
		Ast struct {
		} `json:"ast"`
	} `json:"sources"`
	Contracts map[string]map[string]struct {
		Abi      []interface{} `json:"abi"`
		Metadata string        `json:"metadata"`
		Userdoc  struct {
		} `json:"userdoc"`
		Devdoc struct {
		} `json:"devdoc"`
		Ir            string `json:"ir"`
		StorageLayout struct {
			Storage []interface{} `json:"storage"`
			Types   struct {
			} `json:"types"`
		} `json:"storageLayout"`
		Evm struct {
			Assembly       string `json:"assembly"`
			LegacyAssembly struct {
			} `json:"legacyAssembly"`
			Bytecode struct {
				FunctionDebugData map[string]struct {
					EntryPoint     int `json:"entryPoint"`
					Id             int `json:"id"`
					ParameterSlots int `json:"parameterSlots"`
					ReturnSlots    int `json:"returnSlots"`
				} `json:"functionDebugData"`
				Object           string `json:"object"`
				Opcodes          string `json:"opcodes"`
				SourceMap        string `json:"sourceMap"`
				GeneratedSources []struct {
					Ast struct {
					} `json:"ast"`
					Contents string `json:"contents"`
					Id       int    `json:"id"`
					Language string `json:"language"`
					Name     string `json:"name"`
				} `json:"generatedSources"`
				LinkReferences map[string]map[string][]struct {
					Start  int `json:"start"`
					Length int `json:"length"`
				} `json:"linkReferences"`
			} `json:"bytecode"`
			DeployedBytecode struct {
				ImmutableReferences map[string][]struct {
					Start  int `json:"start"`
					Length int `json:"length"`
				} `json:"immutableReferences"`
			} `json:"deployedBytecode"`
			MethodIdentifiers map[string]string `json:"methodIdentifiers"`
			GasEstimates      struct {
				Creation struct {
					CodeDepositCost string `json:"codeDepositCost"`
					ExecutionCost   string `json:"executionCost"`
					TotalCost       string `json:"totalCost"`
				} `json:"creation"`
				External map[string]string `json:"external"`
				Internal map[string]string `json:"internal"`
			} `json:"gasEstimates"`
		} `json:"evm"`
		Ewasm struct {
			Wast string `json:"wast"`
			Wasm string `json:"wasm"`
		} `json:"ewasm"`
	} `json:"contracts"`
}
