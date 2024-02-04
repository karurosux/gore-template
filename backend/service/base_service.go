package service

import (
	"app/repository"
	userdto "app/service/dto/user_dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseService[T any, DTO any] struct {
	repository repository.BaseRepoInt[T]
	toDTOFunc  func(T) DTO
	toDTOsFunc func([]T) []DTO
}

func (s *BaseService[T, DTO]) FindAll() ([]DTO, error) {
	res, err := s.repository.FindAll()
	if err != nil {
		return []DTO{}, err
	}
	return s.toDTOsFunc(res), nil
}

func (s *BaseService[T, DTO]) FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...func(*gorm.DB) *gorm.DB) ([]DTO, error) {
	res, err := s.repository.FindPagedByUserBranchWithFilter(user, page, limit, filter, scopes...)
	if err != nil {
		return []DTO{}, err
	}
	return s.toDTOsFunc(res), nil
}

func (s *BaseService[T, DTO]) FindPagedByUserBranchWithFilterRaw(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	res, err := s.repository.FindPagedByUserBranchWithFilter(user, page, limit, filter, scopes...)
	if err != nil {
		return []T{}, err
	}
	return res, nil
}

func (s *BaseService[T, DTO]) FindPagedWithFilter(page int, limit int, filter string) ([]DTO, error) {
	res, err := s.repository.FindPagedWithFilter(page, limit, filter)
	if err != nil {
		return []DTO{}, err
	}
	return s.toDTOsFunc(res), nil
}

func (s *BaseService[T, DTO]) FindById(id uuid.UUID) (DTO, error) {
	res, err := s.repository.FindById(id)
	if err != nil {
		return *new(DTO), err
	}
	return s.toDTOFunc(res), nil
}

func (s *BaseService[T, DTO]) FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (DTO, error) {
	res, err := s.repository.FindByUserBranchAndId(user, id)
	if err != nil {
		return *new(DTO), err
	}
	return s.toDTOFunc(res), nil
}

func (s *BaseService[T, DTO]) CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	count, err := s.repository.CountPagedByUserBranchWithFilter(user, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *BaseService[T, DTO]) CountPagedWithFilter(filter string) (int64, error) {
	count, err := s.repository.CountPagedWithFilter(filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *BaseService[T, DTO]) Create(n T) (DTO, error) {
	res, err := s.repository.Create(n)
	if err != nil {
		return *new(DTO), err
	}
	return s.toDTOFunc(res), nil
}

func (s *BaseService[T, DTO]) DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return s.repository.DeleteByUserBranchAndId(user, id)
}

func (s *BaseService[T, DTO]) DeleteById(id uuid.UUID) error {
	return s.repository.DeleteById(id)
}

func (s *BaseService[T, DTO]) PatchById(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error {
	return s.repository.PatchByUserAndId(user, id, body)
}

func (s *BaseService[T, DTO]) PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error {
	return s.repository.PatchByUserAndIdWithColumns(user, id, patch, cols)
}
