package request

type CodeGenerate struct {
	Prompt string `json:"prompt" form:"prompt" validate:"required"`
	Num    int    `json:"n" form:"num"`
	Lang   string `json:"lang" form:"lang"`
}
