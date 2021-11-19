package usecases

import (
	"authz-service/package/casbinhelper"
)

type AuthzUsecase struct {
	authz casbinhelper.Auth
}

func NewAuthzUsecase(authz casbinhelper.Auth) AuthzUsecase {
	return AuthzUsecase{
		authz: authz,
	}
}

func (instance *AuthzUsecase) AddRoleForUserToDomain(userId uint, shopId uint, act string) (bool, error) {
	return instance.authz.AddRoleForUserToDomain(userId, act, shopId)
}

func (instance *AuthzUsecase) GetRolesInDomain(userId uint, shopId uint) ([]string, error) {
	return instance.authz.GetRoleList(userId, shopId)
}

func (instance *AuthzUsecase) CheckPermission(userId uint, shopId uint, act string) (bool, error) {
	return instance.authz.CheckPermission(userId, shopId, act)
}

func (instance *AuthzUsecase) GetImplicitRolesInDomain(userId uint, shopId uint) ([]string, error) {
	return instance.authz.GetDetailRoleList(userId, shopId)
}

func (instance *AuthzUsecase) GenerateOwnerRole(userId uint, shopId uint) error {
	return instance.authz.GenerateOwnerRole(userId, shopId)
}

// func (instance *AuthzUsecase) GetRolesInDomain(userId uint, shopId uint) ([]string, error) {
// 	return instance.authz.GetRoleList(userId, shopId)
// }
