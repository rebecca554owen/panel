package benchmark

type Test struct {
	Name  string `json:"name" validate:"required|in:image,machine,compile,encryption,compression,physics,json,memory,disk"`
	Multi bool   `json:"multi"`
}
