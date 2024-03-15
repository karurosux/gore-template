package base

import (
	userdto "backend/user/dto"

	"github.com/google/uuid"
)

type Service[T any, DTO any] struct {
	Repository RepositoryInt[T]
	ToDTOFunc  func(T) DTO
	ToDTOsFunc func([]T) []DTO
}

func (s *Service[T, DTO]) FindAll() ([]DTO, error) {
	res, err := s.Repository.FindAll()
	if err != nil {
		return []DTO{}, err
	}
	return s.ToDTOsFunc(res), nil
}

func (s *Service[T, DTO]) FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...ScopeFunction) ([]DTO, error) {
	res, err := s.Repository.FindPagedByUserBranchWithFilter(user, page, limit, filter, scopes...)
	if err != nil {
		return []DTO{}, err
	}
	return s.ToDTOsFunc(res), nil
}

func (s *Service[T, DTO]) FindPagedByUserBranchWithFilterRaw(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...ScopeFunction) ([]T, error) {
	res, err := s.Repository.FindPagedByUserBranchWithFilter(user, page, limit, filter, scopes...)
	if err != nil {
		return []T{}, err
	}
	return res, nil
}

func (s *Service[T, DTO]) FindPagedWithFilter(page int, limit int, filter string) ([]DTO, error) {
	res, err := s.Repository.FindPagedWithFilter(page, limit, filter)
	if err != nil {
		return []DTO{}, err
	}
	return s.ToDTOsFunc(res), nil
}

func (s *Service[T, DTO]) FindById(id uuid.UUID) (DTO, error) {
	res, err := s.Repository.FindById(id)
	if err != nil {
		return *new(DTO), err
	}
	return s.ToDTOFunc(res), nil
}

func (s *Service[T, DTO]) FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (DTO, error) {
	res, err := s.Repository.FindByUserBranchAndId(user, id)
	if err != nil {
		return *new(DTO), err
	}
	return s.ToDTOFunc(res), nil
}

func (s *Service[T, DTO]) CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	count, err := s.Repository.CountPagedByUserBranchWithFilter(user, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Service[T, DTO]) CountPagedWithFilter(filter string) (int64, error) {
	count, err := s.Repository.CountPagedWithFilter(filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Service[T, DTO]) Create(n T) (DTO, error) {
	res, err := s.Repository.Create(n)
	if err != nil {
		return *new(DTO), err
	}
	return s.ToDTOFunc(res), nil
}

func (s *Service[T, DTO]) DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return s.Repository.DeleteByUserBranchAndId(user, id)
}

func (s *Service[T, DTO]) DeleteById(id uuid.UUID) error {
	return s.Repository.DeleteById(id)
}

func (s *Service[T, DTO]) PatchById(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error {
	return s.Repository.PatchByUserAndId(user, id, body)
}

func (s *Service[T, DTO]) PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error {
	return s.Repository.PatchByUserAndIdWithColumns(user, id, patch, cols)
}
