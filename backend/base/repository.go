package base

import (
	userdto "backend/user/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScopeFunction func(*gorm.DB) *gorm.DB

type RepositoryInt[T any] interface {
	FindAll() (res []T, err error)
	FindPagedWithFilter(page int, limit int, filter string, scopes ...ScopeFunction) ([]T, error)
	FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...ScopeFunction) ([]T, error)
	FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (T, error)
	FindById(id uuid.UUID) (T, error)
	CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error)
	CountPagedWithFilter(filter string, scopes ...ScopeFunction) (int64, error)
	Create(nr T) (T, error)
	DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error
	DeleteById(id uuid.UUID, scopes ...ScopeFunction) error
	PatchByUserAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error
	PatchById(id uuid.UUID, body T) error
	PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error
}

type (
	FilterFunc       func(filter string) func(*gorm.DB) *gorm.DB
	PatchFunc[T any] func(body T) ([]string, T)
)

type Repository[T any] struct {
	Db         *gorm.DB
	FilterFunc FilterFunc
	PatchFunc  PatchFunc[T]
}

func (r Repository[T]) FindAll() ([]T, error) {
	var res []T
	err := r.Db.Find(&res).Error
	return res, err
}

func (r Repository[T]) FindPagedWithFilter(page int, limit int, filter string, scopes ...ScopeFunction) ([]T, error) {
	var res []T
	staticScopes := [](func(db *gorm.DB) *gorm.DB){AsPage(page, limit), r.FilterFunc(filter)}
	baseStatements := r.Db.Model(new(T)).Scopes(staticScopes...)

	for i := 0; i < len(scopes); i++ {
		s := scopes[i]
		baseStatements = s(baseStatements)
	}

	err := baseStatements.Find(&res).Error
	return res, err
}

func (r Repository[T]) FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...ScopeFunction) ([]T, error) {
	var res []T
	staticScopes := [](func(db *gorm.DB) *gorm.DB){AsPage(page, limit), ForUserBranch(user), r.FilterFunc(filter)}
	baseStatements := r.Db.Model(new(T)).Scopes(staticScopes...)

	for i := 0; i < len(scopes); i++ {
		s := scopes[i]
		baseStatements = s(baseStatements)
	}

	err := baseStatements.Find(&res).Error
	return res, err
}

func (r Repository[T]) FindById(id uuid.UUID) (T, error) {
	var res T
	err := r.Db.Model(new(T)).Where("id = ?", id).Find(&res).Error
	return res, err
}

func (r Repository[T]) FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (T, error) {
	var res T
	err := r.Db.Model(new(T)).Scopes(ForUserBranch(user)).Where("id = ?", id).Find(&res).Error
	return res, err
}

func (r Repository[T]) CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	var count int64
	err := r.Db.Model(new(T)).Scopes(ForUserBranch(user), r.FilterFunc(filter)).Count(&count).Error
	return count, err
}

func (r Repository[T]) CountPagedWithFilter(filter string, scopes ...ScopeFunction) (int64, error) {
	var count int64
	staticScopes := [](func(db *gorm.DB) *gorm.DB){r.FilterFunc(filter)}
	baseStatements := r.Db.Model(new(T)).Scopes(staticScopes...)

	for i := 0; i < len(scopes); i++ {
		s := scopes[i]
		baseStatements = s(baseStatements)
	}

	err := baseStatements.Count(&count).Error
	return count, err
}

func (r Repository[T]) Create(nr T) (T, error) {
	err := r.Db.Create(&nr).Error
	return nr, err
}

func (r Repository[T]) DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return r.Db.Model(new(T)).Scopes(ForUserBranch(user)).Delete(new(T), id).Error
}

func (r Repository[T]) DeleteById(id uuid.UUID, scopes ...ScopeFunction) error {
	baseStatement := r.Db.Model(new(T))

	for i := 0; i < len(scopes); i++ {
		s := scopes[i]
		baseStatement = s(baseStatement)
	}

	return baseStatement.Delete(new(T), id).Error
}

func (r Repository[T]) PatchByUserAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error {
	cols, nb := r.PatchFunc(body)
	return r.Db.Model(new(T)).Scopes(ForUserBranch(user)).Select(cols).Where("id = ?", id).UpdateColumns(nb).Error
}

func (r Repository[T]) PatchById(id uuid.UUID, body T) error {
	cols, nb := r.PatchFunc(body)
	return r.Db.Model(new(T)).Select(cols).Where("id = ?", id).UpdateColumns(nb).Error
}

func (r Repository[T]) PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error {
	return r.Db.Model(new(T)).Scopes(ForUserBranch(user)).Select(cols).Where("id = ?", id).UpdateColumns(patch).Error
}
