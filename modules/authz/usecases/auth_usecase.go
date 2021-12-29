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

func (instance *AuthzUsecase) AddRoleForUserToDomain(userId uint64, shopId uint64, act string) (bool, error) {
	return instance.authz.AddRoleForUserToDomain(userId, act, shopId)
}

func (instance *AuthzUsecase) GetRolesInDomain(userId uint64, shopId uint64) ([]string, error) {
	return instance.authz.GetRoleList(userId, shopId)
}

func (instance *AuthzUsecase) CheckPermission(userId uint64, shopId uint64, act string) (bool, error) {
	return instance.authz.CheckPermission(userId, shopId, act)
}

func (instance *AuthzUsecase) GetImplicitRolesInDomain(userId uint64, shopId uint64) ([]string, error) {
	return instance.authz.GetDetailRoleList(userId, shopId)
}

func (instance *AuthzUsecase) GenerateOwnerRole(userId uint64, shopId uint64) error {
	return instance.authz.GenerateOwnerRole(userId, shopId)
}

func (instance *AuthzUsecase) AddRolesForUserToDomain(userId uint64, roles []string, shopId uint64) (bool, error) {
	return instance.authz.AddRolesForUserToDomain(userId, roles, shopId)
}

func (instance *AuthzUsecase) RemoveRoleFromUser(userId uint64, role string, shopId uint64) (bool, error) {
	return instance.authz.RemoveRoleFromUser(userId, role, shopId)
}

func (instance *AuthzUsecase) RemoveRolesFromUser(userId uint64, roles []string, shopId uint64) (bool, error) {
	return instance.authz.RemoveRolesFromUser(userId, roles, shopId)
}

func (instance *AuthzUsecase) GetUsersForRoleInDomain(role string, shopId uint64) []string {
	return instance.authz.GetUsersForRoleInDomain(role, shopId)
}

func (instance *AuthzUsecase) GetAllUsersByDomain(shopId uint64) []string {
	return instance.authz.GetAllUsersByDomain(shopId)
}

func (instance *AuthzUsecase) DeleteDomains(shopIds []uint64) (bool, error) {
	var shopIdsUint []uint64
	for k := range shopIds {
		shopIdsUint = append(shopIdsUint, shopIds[k])
	}
	return instance.authz.DeleteDomains(shopIdsUint)
}

// func (instance *AuthzUsecase) GetRolesInDomain(userId uint64, shopId uint64) ([]string, error) {
// 	return instance.authz.GetRoleList(userId, shopId)
// }
