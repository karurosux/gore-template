package service

import (
	"app/model"
	"app/repository"
	customerdto "app/service/dto/customer_dto"
	userdto "app/service/dto/user_dto"
	"app/utils"

	"github.com/samber/do"
)

type CustomerService struct {
	customerRepository *repository.CustomerRepository
}

func NewCustomerService(i *do.Injector) (*CustomerService, error) {
	return &CustomerService{
		customerRepository: do.MustInvoke[*repository.CustomerRepository](i),
	}, nil
}

func (cs *CustomerService) GetAllCustomers(user userdto.UserWithRoleAndPermissions) ([]customerdto.CustomerDTO, error) {
	u, err := cs.customerRepository.GetAllCustomers(user)
	if err != nil {
		return []customerdto.CustomerDTO{}, err
	}

	return customerdto.ToCustomerDTOs(u), nil
}

func (cs *CustomerService) GetPageByUser(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string) (model.Paginated[customerdto.CustomerDTO], error) {
	customers, err := cs.customerRepository.GetPageByUser(user, page, limit, filter)
	if err != nil {
		return model.Paginated[customerdto.CustomerDTO]{}, err
	}

	count, err := cs.customerRepository.GetCountByUser(user, filter)
	if err != nil {
		return model.Paginated[customerdto.CustomerDTO]{}, err
	}

	return model.Paginated[customerdto.CustomerDTO]{
		Data: customerdto.ToCustomerDTOs(customers),
		Meta: utils.CalculatePaginationMeta(count, page, limit),
	}, nil
}
