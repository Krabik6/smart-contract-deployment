package api

type DeployRequest struct {
	SourceCode           string   `json:"source_code"`
	ConstructorArguments []string `json:"constructor_arguments"`
}

type SourceCodeRequest struct {
	SourceCode string `json:"source_code"`
}
