package response

import "github.com/devops-codegpt/server/models"

// PageData Encapsulates paginated data
type PageData struct {
	models.PageInfo
	List any `json:"list"`
}
