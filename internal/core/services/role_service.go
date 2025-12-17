package services

import "github.com/amirbakhtiari/orga/internal/core/ports"

type RoleService struct {
	repo ports.RoleRepository
}

func NewRoleService(repo ports.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}
