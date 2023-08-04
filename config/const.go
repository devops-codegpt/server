package config

const (
	// configType is the config file type
	configType = "yml"
	// developmentConfig is the path of development config yml
	developmentConfig = "config.dev.yml"
)

const (
	// API represents the group of API
	API = "/api"
)

const (
	APIHealth = API + "/health"
)

const (
	// APIUser represents the group of user management API
	APIUser = API + "/user"
	// APIUserLogin represents the API to login by authentication
	APIUserLogin = APIUser + "/login"
	APIUserList  = APIUser + "/list"
	APIUserInfo  = APIUser + "/info"
)

const (
	APIRole       = API + "/role"
	APIRoleList   = APIRole + "/list"
	APIRoleCreate = APIRole + "/create"
	APIRoleUpdate = APIRole + "/update/:roleId"
)

const (
	APIConversation            = "/api/conversation"
	APIConversationList        = APIConversation + "/list"
	APIConversationInfo        = APIConversation + "/:conversationId"
	APIMessageFeedback         = APIConversation + "/message_feedback"
	APIConversationBatchDelete = APIConversation + "/delete/batch"
	APIConversationWS          = APIConversation + "/ws"
	APIConversationBase        = APIConversation + "/base"
)

const (
	APICode         = "/api/code"
	APICodeGenerate = APICode + "/generate"
)

const (
	APILLM     = "/api/llm"
	APILLMList = APILLM + "/list"
)
