package service

import (
	"app/entities"
	"app/model"
	"app/repository"
	userdto "app/service/dto/user_dto"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type UserService struct {
	userRepository *repository.UserRepository
	roleService    *RoleService
}

func NewUserService(i *do.Injector) (*UserService, error) {
	return &UserService{
		userRepository: do.MustInvoke[*repository.UserRepository](i),
		roleService:    do.MustInvoke[*RoleService](i),
	}, nil
}

func (us *UserService) GetPaginated(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string) (model.Paginated[userdto.UserWithRoleAndPermissions], error) {
	users, err := us.userRepository.GetPageByUser(user, page, limit, filter)
	if err != nil {
		return model.Paginated[userdto.UserWithRoleAndPermissions]{}, err
	}

	count, err := us.userRepository.GetCountByUser(user, filter)
	if err != nil {
		return model.Paginated[userdto.UserWithRoleAndPermissions]{}, err
	}

	return model.Paginated[userdto.UserWithRoleAndPermissions]{
		Data: userdto.ToUsersWithRoleAndPermissions(users),
		Meta: model.PaginatedMeta{
			Total:       count,
			PerPage:     int32(limit),
			CurrentPage: int32(page),
		},
	}, nil
}

func (us *UserService) GetById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (userdto.UserDTO, error) {
	foundUser, err := us.userRepository.GetById(user, id)

	if err != nil {
		return userdto.UserDTO{}, err
	}

	return userdto.ToUserDto(foundUser), nil
}

func (us *UserService) GetByEmailWithPassword(email string) (userdto.UserWithPasswordDTO, error) {
	user, err := us.userRepository.GetByEmail(email)

	if err != nil {
		return userdto.UserWithPasswordDTO{}, err
	}

	return userdto.ToUserWithPassword(user), nil
}

func (us *UserService) GetByEmailWithRoleAndPermissions(email string) (userdto.UserWithRoleAndPermissions, error) {
	u, err := us.userRepository.GetByEmail(email)
	if err != nil {
		return userdto.UserWithRoleAndPermissions{}, err
	}

	r, err := us.roleService.GetRoleWithPermissionsById(u.RoleId)
	if err != nil {
		return userdto.UserWithRoleAndPermissions{}, err
	}

	return userdto.ToUserWithRoleAndPermissions(u, r), nil
}

func (us *UserService) DeleteById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return us.userRepository.DeleteById(user, id)
}

func (us *UserService) Create(user userdto.UserWithRoleAndPermissions, newUser userdto.CreateUserDTO) (userdto.UserDTO, error) {
	created, err := us.userRepository.CreateUser(entities.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		RoleId:    uuid.NullUUID{UUID: newUser.RoleId, Valid: true},
		BranchId:  user.BranchId,
	})
	return userdto.ToUserDto(created), err
}
