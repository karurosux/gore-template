package user

import (
	"backend/base"
	"backend/role"
	"backend/user/dto"
	"backend/user/entity"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type Service struct {
	base.Service[entity.User, dto.UserDTO]
	passwordService *base.PasswordService
	roleService     *role.Service
}

func NewService(i *do.Injector) (*Service, error) {
	return &Service{
		Service: base.Service[entity.User, dto.UserDTO]{
			Repository: do.MustInvoke[Repository](i),
			ToDTOFunc:  dto.ToUserDTO,
			ToDTOsFunc: dto.ToUserDTOs,
		},
		passwordService: do.MustInvoke[*base.PasswordService](i),
		roleService:     do.MustInvoke[*role.Service](i),
	}, nil
}

func (us *Service) Create(u entity.User) (dto.UserDTO, error) {
	pass, err := us.passwordService.HashPassword(u.Password)
	if err != nil {
		return dto.UserDTO{}, err
	}
	u.Password = pass
	return us.Service.Create(u)
}

func (us *Service) FindByEmailWithPassword(email string) (dto.UserWithPasswordDTO, error) {
	repo := us.Repository.(Repository)
	user, err := repo.FindByEmail(email)
	if err != nil {
		return dto.UserWithPasswordDTO{}, err
	}

	return dto.ToUserWithPassword(user), nil
}

func (us *Service) FindByEmailWithRoleAndPermissions(email string) (dto.UserWithRoleAndPermissions, error) {
	repo := us.Repository.(Repository)
	u, err := repo.FindByEmail(email)
	if err != nil {
		return dto.UserWithRoleAndPermissions{}, err
	}

	res, err := us.roleService.GetRoleWithPermissionsById(u.RoleId)
	if err != nil {
		return dto.UserWithRoleAndPermissions{}, err
	}

	return dto.ToUserWithRoleAndPermissions(u, res), nil
}

func (us *Service) FindPagedWithRoleAndPermissions(user dto.UserWithRoleAndPermissions, page int, limit int, filter string) ([]dto.UserWithRoleAndPermissions, error) {
	res, err := us.FindPagedByUserBranchWithFilterRaw(user, page, limit, filter, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Role")
	})
	if err != nil {
		return []dto.UserWithRoleAndPermissions{}, err
	}
	return dto.ToUsersWithRoleAndPermissions(res), nil
}
