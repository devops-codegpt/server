package migration

import (
	"errors"
	"fmt"
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/models"
	"gorm.io/gorm"
)

// InitData inits the data used in application
func InitData(c container.Container) error {
	if c.GetConfig().Database.InitData {
		initRoles(c)
		initUsers(c)
		initUrls(c)
		err := initCasbinRules(c)
		return err
	}
	return nil
}

var (
	roleAdminSort   = uint(models.AdminRoleSort)
	roleDefaultSort = uint(models.DefaultRoleSort)
	creator         = "auto create"
	adminKeyword    = "admin"
	testKeyword     = "test"
)

func initRoles(c container.Container) {
	rep := c.GetRepository()

	roles := []*models.Role{
		{
			Name:    "administrator",
			Keyword: adminKeyword,
			Sort:    &roleAdminSort,
			Creator: creator,
		},
		{
			Name:    "tester",
			Keyword: testKeyword,
			Sort:    &roleDefaultSort,
			Creator: creator,
		},
	}
	newRoles := make([]*models.Role, 0)
	for i, role := range roles {
		id := uint(i + 1)
		if err := rep.Where("id = ?", id).First(&models.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			role.Id = id
			newRoles = append(newRoles, role)
		}
	}
	if len(newRoles) > 0 {
		rep.Create(&newRoles)
	}
}

func initUsers(c container.Container) {
	rep := c.GetRepository()

	users := []*models.User{
		{
			Username: "12345678",
			ZhName:   "超管",
			Email:    "super@mail.com",
			Creator:  creator,
		},
		{
			Username: "87654321",
			ZhName:   "游客",
			Email:    "example@mail.com",
			Creator:  creator,
		},
	}
	newUsers := make([]*models.User, 0)
	for i, user := range users {
		id := uint(i + 1)
		if err := rep.Where("id = ?", id).First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			user.Id = id
			if user.RoleId == 0 {
				user.RoleId = id
			}
			newUsers = append(newUsers, user)
		}
	}
	if len(newUsers) > 0 {
		rep.Create(&newUsers)
	}
}

func getInitUrls() []*models.Url { // nolint:funlen
	urls := []*models.Url{
		{
			Method:   "GET",
			Path:     config.APIHealth,
			Category: "base",
			Desc:     "Service Activity Detection",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APIUserLogin,
			Category: "user",
			Desc:     "User Authentication",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIUserList,
			Category: "user",
			Desc:     "Get user list",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIUserInfo,
			Category: "user",
			Desc:     "Get user info",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIRoleList,
			Category: "role",
			Desc:     "Get role list",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APIRoleCreate,
			Category: "role",
			Desc:     "Creating a Role",
			Creator:  creator,
		},
		{
			Method:   "PATCH",
			Path:     config.APIRoleUpdate,
			Category: "role",
			Desc:     "Update role",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIConversationList,
			Category: "conversation",
			Desc:     "Get conversation list",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APIConversation,
			Category: "conversation",
			Desc:     "Create conversation and Continue it",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIConversationInfo,
			Category: "conversation",
			Desc:     "Get conversation info",
			Creator:  creator,
		},
		{
			Method:   "DELETE",
			Path:     config.APIConversationBatchDelete,
			Category: "conversation",
			Desc:     "Batch delete conversations",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APIMessageFeedback,
			Category: "conversation",
			Desc:     "Score a  message record",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APIConversationBase,
			Category: "conversation",
			Desc:     "Provide basic chatGPT service",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APIConversationWS,
			Category: "conversation",
			Desc:     "Conversation websocket",
			Creator:  creator,
		},
		{
			Method:   "GET",
			Path:     config.APILLMList,
			Category: "llm",
			Desc:     "Get llm list",
			Creator:  creator,
		},
		{
			Method:   "POST",
			Path:     config.APICodeGenerate,
			Category: "code",
			Desc:     "Code generation",
			Creator:  creator,
		},
	}
	return urls
}

func initUrls(c container.Container) {
	rep := c.GetRepository()
	urls := getInitUrls()
	newUrls := make([]*models.Url, 0)
	for i, url := range urls {
		id := uint(i + 1)
		if err := rep.Where("id = ?", id).First(&models.Url{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			url.Id = id
			newUrls = append(newUrls, url)
		}
	}
	if len(newUrls) > 0 {
		rep.Create(&newUrls)
	}
}

func initCasbinRules(c container.Container) error {
	rep := c.GetRepository()
	auth := c.GetAuthentication()

	rules := make([][]string, 0)
	urls := getInitUrls()
	// Admin role has access to all paths
	for _, url := range urls {
		rules = append(rules, []string{
			adminKeyword,
			url.Path,
			url.Method,
		})
	}

	// Other role has access to base paths
	basePaths := []string{
		config.APIHealth,
		config.APIUserLogin,
	}
	baseUrls := make([]*models.Url, 0)
	err := rep.Model(&models.Url{}).Where("path IN (?)", basePaths).Find(&baseUrls).Error
	if err != nil {
		fmt.Printf("Query base urls failed: %v\n", err)
		return err
	}
	for _, url := range baseUrls {
		rules = append(rules, []string{
			testKeyword,
			url.Path,
			url.Method,
		})
	}
	// Add rule to database
	if len(rules) > 0 {
		_, err = auth.AddPolicies(rules)
		if err != nil {
			fmt.Printf("Add rules failed: %v", err)
			return err
		}
	}
	return nil
}
