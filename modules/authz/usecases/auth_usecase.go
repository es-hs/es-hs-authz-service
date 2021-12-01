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

func (instance *AuthzUsecase) AddRolesForUserToDomain(userId uint, roles []string, shopId uint) (bool, error) {
	return instance.authz.AddRolesForUserToDomain(userId, roles, shopId)
}

func (instance *AuthzUsecase) RemoveRoleFromUser(userId uint, role string, shopId uint) (bool, error) {
	return instance.authz.RemoveRoleFromUser(userId, role, shopId)
}

func (instance *AuthzUsecase) RemoveRolesFromUser(userId uint, roles []string, shopId uint) (bool, error) {
	return instance.authz.RemoveRolesFromUser(userId, roles, shopId)
}

func (instance *AuthzUsecase) GetUsersForRoleInDomain(role string, shopId uint) []string {
	return instance.authz.GetUsersForRoleInDomain(role, shopId)
}

func (instance *AuthzUsecase) GetAllUsersByDomain(shopId uint) []string {
	return instance.authz.GetAllUsersByDomain(shopId)
}

func (instance *AuthzUsecase) DeleteDomains(shopIds []int64) (bool, error) {
	var shopIdsUint []uint
	for k := range shopIds {
		shopIdsUint = append(shopIdsUint, uint(shopIds[k]))
	}
	return instance.authz.DeleteDomains(shopIdsUint)
}

// func (instance *AuthzUsecase) GetRolesInDomain(userId uint, shopId uint) ([]string, error) {
// 	return instance.authz.GetRoleList(userId, shopId)
// }
