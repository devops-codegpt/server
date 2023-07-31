package response

import "time"

type Input struct {
	Lang string `json:"lang"`
	N    int    `json:"n"`
	Stop []any  `json:"stop"`
	Text string `json:"text"`
}

type Output struct {
	Code               []string `json:"code"`
	ErrCode            int      `json:"errCode"`
	CompletionTokenNum int      `json:"completion_token_num"`
	PromptTokenNum     int      `json:"prompt_token_num"`
}

type User struct {
	AppId string `json:"appId"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

type Progress struct {
}

type Result struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	App         string    `json:"app"`
	Status      string    `json:"status"`
	TaskId      string    `json:"task_id"`
	ProcessTime int       `json:"process_time"`
	Input       Input     `json:"input"`
	Output      Output    `json:"output"`
	User        User      `json:"user"`
	Progress    Progress  `json:"progress"`
}

type CodeGenerate struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Result  Result `json:"result"`
}
