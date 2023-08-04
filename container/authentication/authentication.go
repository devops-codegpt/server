// Package authentication uses Casbin(https://github.com/casbin/casbin) to load auth rules from database.
package authentication

import (
	_ "embed"
	"fmt"
	"github.com/casbin/casbin/v2"
	casbin_model "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/devops-codegpt/server/container/repository"
)

//go:embed rbac_model.conf
var confContent string

// CasbinRule represents  casbin rules
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:16"`
	V0    string `gorm:"size:128;comment:'role keyword'"`
	V1    string `gorm:"size:128;comment:'request path'"`
	V2    string `gorm:"size:256;comment:'request method'"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

type Authentication interface {
	AddPolicies(rules [][]string) (bool, error)
	AddPoliciesEx(rules [][]string) (bool, error)
	AddPolicy(params ...any) (bool, error)
	RemovePolicy(params ...any) (bool, error)
	RemovePolicies(rules [][]string) (bool, error)
	GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string
	GetCasbinEnforcer() *casbin.Enforcer
}

type authentication struct {
	casbin *casbin.Enforcer
}

func (a *authentication) AddPolicies(rules [][]string) (bool, error) {
	return a.casbin.AddPolicies(rules)
}

func (a *authentication) AddPolicy(params ...any) (bool, error) {
	return a.casbin.AddPolicy(params...)
}

func (a *authentication) RemovePolicy(params ...any) (bool, error) {
	return a.casbin.RemovePolicy(params...)
}

func (a *authentication) RemovePolicies(rules [][]string) (bool, error) {
	return a.casbin.RemovePolicies(rules)
}

func (a *authentication) GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string {
	return a.casbin.GetFilteredPolicy(fieldIndex, fieldValues...)
}

func (a *authentication) GetCasbinEnforcer() *casbin.Enforcer {
	return a.casbin
}

func (a *authentication) AddPoliciesEx(rules [][]string) (bool, error) {
	return a.casbin.AddPoliciesEx(rules)
}

// NewAuthentication is constructor for authentication
func NewAuthentication(rep repository.Repository) (Authentication, error) {
	rule := CasbinRule{}
	// Initialize a Gorm adapter and use it in a Casbin enforcer
	a, err := gormadapter.NewAdapterByDBWithCustomTable(rep.Model(&rule), &rule)
	if err != nil {
		fmt.Printf("Failed to create casbin rule table: %v", err)
		return nil, err
	}
	// Load config
	m, err := casbin_model.NewModelFromString(confContent)
	if err != nil {
		fmt.Printf("Failed to load casbin config: %v", err)
		return nil, err
	}
	// Create enforcer
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		fmt.Printf("Failed to init authentication: %v", err)
		return nil, err
	}
	// Load rules
	err = e.LoadPolicy()
	if err != nil {
		fmt.Printf("Failed to load casbin rules: %v", err)
		return nil, err
	}

	return &authentication{casbin: e}, nil
}
