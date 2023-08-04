package handler

import (
	"encoding/json"
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/internal/llm/memory"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
	"github.com/devops-codegpt/server/service"
	"github.com/devops-codegpt/server/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// allow cross-domain.
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ConversationHandler interface {
	GetConversations(c echo.Context) error
	GetById(c echo.Context) error
	RunConversation(c echo.Context) error
	RunConversationWS(c echo.Context) error
	RunBase(c echo.Context) error
	BatchDeleteByIds(c echo.Context) error
	FeedbackMessage(c echo.Context) error
}

type conversationHandler struct {
	context container.Container
	service service.ConversationService
}

// GetConversations get all conversations of the current user
func (ch *conversationHandler) GetConversations(c echo.Context) error {
	var req request.ConversationList
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Get user information from the current jwt claims
	claims := getJwtLoginClaims(c)
	username := claims.UserName
	conversations, err := ch.service.FindConversations(username, &req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}

	var resp response.PageData
	resp.PageInfo = req.PageInfo
	resp.List = conversations
	return successWithData(c, resp)
}

// GetById gets the chat messages of a conversation
func (ch *conversationHandler) GetById(c echo.Context) error {
	id := c.Param("conversationId")
	conversation, err := ch.service.FindById(id)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return successWithData(c, conversation)
}

func (ch *conversationHandler) RunConversation(c echo.Context) error {
	var req request.ConversationRun
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Params validate
	if err := c.Validate(req); err != nil {
		return err
	}

	// Get current user
	jwtClaim := getJwtLoginClaims(c)
	username := jwtClaim.UserName
	callback := &streamCallback{
		context: c,
	}

	message, err := ch.service.RunConversation(username, &req, callback)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	var resp response.ConversationRun
	utils.Struct2Struct(message, &resp)
	return successWithData(c, resp)
}

// RunBase provides basic chatGPT service, that is, enter prompt and return the answer result
func (ch *conversationHandler) RunBase(c echo.Context) error {
	var req request.BaseRun
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Params validate
	if err := c.Validate(req); err != nil {
		return err
	}
	resp, err := ch.service.RunBase(req.Prompt)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return successWithData(c, resp)
}

func (ch *conversationHandler) BatchDeleteByIds(c echo.Context) error {
	var req request.ReqDelete
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	err := ch.service.DeleteByIds(req.Ids)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return success(c)
}

func (ch *conversationHandler) FeedbackMessage(c echo.Context) error {
	var req request.MessageFeedback
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Params validate
	if err := c.Validate(req); err != nil {
		return err
	}

	err := ch.service.FeedbackMessage(&req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return success(c)
}

func (ch *conversationHandler) RunConversationWS(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	buffer := memory.NewBuffer()
	callback := &wsStreamCallback{ws: ws}
	defer func(ws *websocket.Conn) {
		_ = buffer.Clear()
		_ = ws.Close()
	}(ws)

	for {
		// Read
		req := request.ConversationWS{}
		err = ws.ReadJSON(&req)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				ch.context.GetLogger().GetZapLogger().Errorf("error: %v", err)
			}
			break
		}
		startResp := response.ConversationWS{
			Message: "",
			Type:    "start",
			Sender:  "bot",
		}
		_ = ws.WriteJSON(&startResp)
		err = ch.service.RunConversationWS(req.Command, req.Prompt, buffer, callback)
		if err != nil {
			ch.context.GetLogger().GetZapLogger().Errorf("failed to get response information: %v", err)
			errResp := &response.ConversationWS{
				Message: fmt.Sprintf("sorry, something failed: %v", err),
				Type:    "error",
				Sender:  "bot",
			}
			_ = ws.WriteJSON(&errResp)
		}
		endResp := response.ConversationWS{
			Message: "",
			Type:    "end",
			Sender:  "bot",
		}
		_ = ws.WriteJSON(&endResp)
	}
	return nil
}

func NewConversationHandler(c container.Container, s service.ConversationService) ConversationHandler {
	if s == nil {
		s = service.NewConversationService(c)
	}
	return &conversationHandler{
		context: c,
		service: s,
	}
}

type streamCallback struct {
	context echo.Context
}

// OnLLMNewToken Handling Streaming Responses
func (sc *streamCallback) OnLLMNewToken(token string) {
	c := sc.context
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())
	if err := enc.Encode(fmt.Sprintf("data: %s\n\n", token)); err != nil {
		fmt.Printf("encode response tokens failed")
		return
	}
	c.Response().Flush()
}

type wsStreamCallback struct {
	ws *websocket.Conn
}

func (w *wsStreamCallback) OnLLMNewToken(token string) {
	resp := response.ConversationWS{
		Message: token,
		Type:    "stream",
		Sender:  "bot",
	}
	if err := w.ws.WriteJSON(&resp); err != nil {
		fmt.Printf("Sending data using websocket fails： %v", err)
	}
}
