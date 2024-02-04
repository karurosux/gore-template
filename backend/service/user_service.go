package service

import (
	"app/entities"
	"app/repository"
	userdto "app/service/dto/user_dto"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type UserService struct {
	BaseService[entities.User, userdto.UserDTO]
	passwordService *PasswordService
	roleService     *RoleService
}

func NewUserService(i *do.Injector) (*UserService, error) {
	return &UserService{
		BaseService: BaseService[entities.User, userdto.UserDTO]{
			repository: do.MustInvoke[repository.UserRepository](i),
			toDTOFunc:  userdto.ToUserDTO,
			toDTOsFunc: userdto.ToUserDTOs,
		},
		passwordService: do.MustInvoke[*PasswordService](i),
		roleService:     do.MustInvoke[*RoleService](i),
	}, nil
}

func (us *UserService) Create(u entities.User) (userdto.UserDTO, error) {
	pass, err := us.passwordService.HashPassword(u.Password)
	if err != nil {
		return userdto.UserDTO{}, err
	}
	u.Password = pass
	return us.BaseService.Create(u)
}

func (us *UserService) FindByEmailWithPassword(email string) (userdto.UserWithPasswordDTO, error) {
	repo := us.repository.(repository.UserRepository)
	user, err := repo.FindByEmail(email)
	if err != nil {
		return userdto.UserWithPasswordDTO{}, err
	}

	return userdto.ToUserWithPassword(user), nil
}

func (us *UserService) FindByEmailWithRoleAndPermissions(email string) (userdto.UserWithRoleAndPermissions, error) {
	repo := us.repository.(repository.UserRepository)
	u, err := repo.FindByEmail(email)
	if err != nil {
		return userdto.UserWithRoleAndPermissions{}, err
	}

	r, err := us.roleService.GetRoleWithPermissionsById(u.RoleId)
	if err != nil {
		return userdto.UserWithRoleAndPermissions{}, err
	}

	return userdto.ToUserWithRoleAndPermissions(u, r), nil
}

func (us *UserService) FindPagedWithRoleAndPermissions(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string) ([]userdto.UserWithRoleAndPermissions, error) {
	res, err := us.FindPagedByUserBranchWithFilterRaw(user, page, limit, filter, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Role")
	})
	if err != nil {
		return []userdto.UserWithRoleAndPermissions{}, err
	}
	return userdto.ToUsersWithRoleAndPermissions(res), nil
}
