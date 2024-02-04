package repository

import (
	userdto "app/service/dto/user_dto"
	"app/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepoInt[T any] interface {
	FindAll() (res []T, err error)
	FindPagedWithFilter(page int, limit int, filter string) ([]T, error)
	FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...func(*gorm.DB) *gorm.DB) ([]T, error)
	FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (T, error)
	FindById(id uuid.UUID) (T, error)
	CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error)
	CountPagedWithFilter(filter string) (int64, error)
	Create(nr T) (T, error)
	DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error
	DeleteById(id uuid.UUID) error
	PatchByUserAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error
	PatchById(id uuid.UUID, body T) error
	PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error
}

type (
	FilterFunc       func(filter string) func(*gorm.DB) *gorm.DB
	PatchFunc[T any] func(body T) ([]string, T)
)

type BaseRepository[T any] struct {
	db         *gorm.DB
	filterFunc FilterFunc
	patchFunc  PatchFunc[T]
}

func (r BaseRepository[T]) FindAll() ([]T, error) {
	var res []T
	err := r.db.Find(&res).Error
	return res, err
}

func (r BaseRepository[T]) FindPagedWithFilter(page int, limit int, filter string) ([]T, error) {
	var res []T
	err := r.db.Model(new(T)).Scopes(utils.AsPage(page, limit), r.filterFunc(filter)).Find(&res).Error
	return res, err
}

func (r BaseRepository[T]) FindPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string, scopes ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var res []T
	staticScopes := [](func(*gorm.DB) *gorm.DB){utils.AsPage(page, limit), utils.ForUserBranch(user), r.filterFunc(filter)}
	baseStatements := r.db.Model(new(T)).Scopes(staticScopes...)

	for i := 0; i < len(scopes); i++ {
		s := scopes[i]
		baseStatements = s(baseStatements)
	}

	err := baseStatements.Find(&res).Error
	return res, err
}

func (r BaseRepository[T]) FindById(id uuid.UUID) (T, error) {
	var res T
	err := r.db.Model(new(T)).Where("id = ?", id).Find(&res).Error
	return res, err
}

func (r BaseRepository[T]) FindByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (T, error) {
	var res T
	err := r.db.Model(new(T)).Scopes(utils.ForUserBranch(user)).Where("id = ?", id).Find(&res).Error
	return res, err
}

func (r BaseRepository[T]) CountPagedByUserBranchWithFilter(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	var count int64
	err := r.db.Model(new(T)).Scopes(utils.ForUserBranch(user), r.filterFunc(filter)).Count(&count).Error
	return count, err
}

func (r BaseRepository[T]) CountPagedWithFilter(filter string) (int64, error) {
	var count int64
	err := r.db.Model(new(T)).Scopes(r.filterFunc(filter)).Count(&count).Error
	return count, err
}

func (r BaseRepository[T]) Create(nr T) (T, error) {
	err := r.db.Create(&nr).Error
	return nr, err
}

func (r BaseRepository[T]) DeleteByUserBranchAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return r.db.Model(new(T)).Scopes(utils.ForUserBranch(user)).Delete(new(T), id).Error
}

func (r BaseRepository[T]) DeleteById(id uuid.UUID) error {
	return r.db.Model(new(T)).Delete(new(T), id).Error
}

func (r BaseRepository[T]) PatchByUserAndId(user userdto.UserWithRoleAndPermissions, id uuid.UUID, body T) error {
	cols, nb := r.patchFunc(body)
	return r.db.Model(new(T)).Scopes(utils.ForUserBranch(user)).Select(cols).Where("id = ?", id).UpdateColumns(nb).Error
}

func (r BaseRepository[T]) PatchById(id uuid.UUID, body T) error {
	cols, nb := r.patchFunc(body)
	return r.db.Model(new(T)).Select(cols).Where("id = ?", id).UpdateColumns(nb).Error
}

func (r BaseRepository[T]) PatchByUserAndIdWithColumns(user userdto.UserWithRoleAndPermissions, id uuid.UUID, patch T, cols []string) error {
	return r.db.Model(new(T)).Scopes(utils.ForUserBranch(user)).Select(cols).Where("id = ?", id).UpdateColumns(patch).Error
}
