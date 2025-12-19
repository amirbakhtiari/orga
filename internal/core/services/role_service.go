package services

import (
	"context"
	"errors"
	"github.com/amirbakhtiari/orga/internal/core/domain"
	"github.com/amirbakhtiari/orga/internal/core/ports"
	dto "github.com/amirbakhtiari/orga/internal/core/services/dto/role"
	"strings"
)

type RoleService struct {
	repo ports.RoleRepository
}

func NewRoleService(repo ports.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(
	ctx context.Context,
	input dto.CreateRoleDTO,
) (*domain.Role, error) {
	input.Code = strings.TrimSpace(input.Code)
	input.Name = strings.TrimSpace(input.Name)

	if input.Code == "" {
		return nil, errors.New("role code is required")
	}

	if input.Name == "" {
		return nil, errors.New("role name is required")
	}

	exists, err := s.repo.GetByCode(ctx, input.Code)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("role already exists")
	}
	newRole := &domain.Role{
		Name:        input.Name,
		Code:        input.Code,
		Description: input.Description,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	if err := s.repo.Create(ctx, newRole); err != nil {
		return nil, err
	}
	return newRole, nil
}

func (s *RoleService) UpdateRole(
	ctx context.Context,
	input dto.UpdateRoleDTO,
) (*domain.Role, error) {
	if input.ID == 0 {
		return nil, errors.New("role id is required")
	}

	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.New("role name is required")
	}

	role, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	role.Name = input.Name
	role.Description = input.Description

	if err := s.repo.Update(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) GetRoleByID(ctx context.Context, id uint) (*domain.Role, error) {
	if id == 0 {
		return nil, errors.New("invalid role id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *RoleService) GetRoleByCode(ctx context.Context, code string) (*domain.Role, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return nil, errors.New("role code is required")
	}
	return s.repo.GetByCode(ctx, code)
}

func (s *RoleService) ListRole(ctx context.Context) ([]*domain.Role, error) {
	return s.repo.List(ctx)
}

func (s *RoleService) DeleteRole(
	ctx context.Context,
	id uint,
) error {
	if id == 0 {
		return errors.New("invalid role id")
	}
	return s.repo.Delete(ctx, id)
}
