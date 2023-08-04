package response

type ConversationRun struct {
	ID             string `json:"id"`
	Role           string `json:"role"`
	Content        string `json:"content"`
	ContentType    string `json:"contentType"`
	Model          string `json:"models"`
	Feedback       uint   `json:"feedback"`
	ParentMsgId    string `json:"parentMsgId"`
	ConversationId string `json:"conversationId"`
}

type ConversationWS struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Sender  string `json:"sender"`
}
