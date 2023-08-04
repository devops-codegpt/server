// Package router registers the routing of this application
package router

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/handler"
	"github.com/labstack/echo/v4"
)

// Init initialize the routing of this application
func Init(e *echo.Echo, c container.Container) {
	setUserHandler(e, c)
	setRoleHandler(e, c)
	setHealthHandler(e, c)
	setConversationHandler(e, c)
	setLLMHandler(e, c)
	setCodeHandler(e, c)
}

func setUserHandler(e *echo.Echo, c container.Container) {
	user := handler.NewUserHandler(c, nil)
	e.POST(config.APIUserLogin, func(c echo.Context) error { return user.Login(c) })
	e.GET(config.APIUserList, func(c echo.Context) error { return user.GetUsers(c) })
	e.GET(config.APIUserInfo, func(c echo.Context) error { return user.GetCurrentUser(c) })
}

func setRoleHandler(e *echo.Echo, c container.Container) {
	role := handler.NewRoleHandler(c, nil)
	e.GET(config.APIRoleList, func(c echo.Context) error { return role.GetRoles(c) })
	e.PATCH(config.APIRoleUpdate, func(c echo.Context) error { return role.UpdateRoleById(c) })
	e.POST(config.APIRoleCreate, func(c echo.Context) error { return role.CreateRole(c) })
}

func setHealthHandler(e *echo.Echo, c container.Container) {
	health := handler.NewHealthHandler(c)
	e.GET(config.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })
}

func setConversationHandler(e *echo.Echo, c container.Container) {
	conversation := handler.NewConversationHandler(c, nil)
	e.GET(config.APIConversationList, func(c echo.Context) error { return conversation.GetConversations(c) })
	e.POST(config.APIConversation, func(c echo.Context) error { return conversation.RunConversation(c) })
	e.GET(config.APIConversationInfo, func(c echo.Context) error { return conversation.GetById(c) })
	e.DELETE(config.APIConversationBatchDelete, func(c echo.Context) error { return conversation.BatchDeleteByIds(c) })
	e.POST(config.APIMessageFeedback, func(c echo.Context) error { return conversation.FeedbackMessage(c) })
	e.GET(config.APIConversationWS, func(c echo.Context) error { return conversation.RunConversationWS(c) })
	e.POST(config.APIConversationBase, func(c echo.Context) error { return conversation.RunBase(c) })
}

func setCodeHandler(e *echo.Echo, c container.Container) {
	code := handler.NewCodeHandler(c)
	e.POST(config.APICodeGenerate, func(c echo.Context) error { return code.Generate(c) })
}

func setLLMHandler(e *echo.Echo, c container.Container) {
	llm := handler.NewLLMHandler(c, nil)
	e.GET(config.APILLMList, func(c echo.Context) error { return llm.GetLLMs(c) })
}
