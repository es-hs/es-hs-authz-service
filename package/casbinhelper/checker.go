package casbinhelper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
)

type Auth struct {
	enforcer *casbin.Enforcer
}

func (a *Auth) CheckPermission(userID uint, shopID uint, act string) (result bool, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return a.Enforce(sub, domain, act)
}

func (a *Auth) GetRoleList(userID uint, shopID uint) (roleList []string, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return a.GetRolesForUser(sub, domain)
}

func (a *Auth) GetDetailRoleList(userID uint, shopID uint) (roleList []string, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return a.GetImplicitRolesForUser(sub, domain)
}

func (a *Auth) GenerateOwnerRole(userID uint, shopID uint) (err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	a.AddRoleToDomain(sub, OWNER_ROLE, domain)
	for role := range mapScope {
		if role == "owner" || role == "admin" {
			continue
		}
		a.AddRoleToDomain(ADMIN_ROLE, role, domain)
	}
	a.AddRoleToDomain(OWNER_ROLE, ADMIN_ROLE, domain)
	return nil
}

func (a *Auth) AddRoleForUserToDomain(userID uint, role string, shopID uint) (result bool, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return a.AddRoleToDomain(sub, role, domain)
}

func (a *Auth) AddRolesForUserToDomain(userID uint, roles []string, shopID uint) (result bool, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return a.AddRolesToDomain(sub, roles, domain)
}

func (a *Auth) RemoveRoleFromUser(userID uint, role string, shopID uint) (result bool, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	//
	e := a.enforcer
	return e.RemoveNamedGroupingPolicy("g", sub, role, domain)

	// return a.RemoveRoleFromUser(sub, role, domain)
}

func (a *Auth) RemoveRolesFromUser(userID uint, roles []string, shopID uint) (result bool, err error) {
	sub := fmt.Sprintf("user_%s", strconv.Itoa(int(userID)))
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	//
	e := a.enforcer
	rolesToRemove := [][]string{}
	for k := range roles {
		// if _, ok := mapScope[roles[k]]; !ok {
		// 	return false, errors.New("Role not found in scope")
		// }
		rolesToRemove = append(rolesToRemove, []string{sub, roles[k], domain})
	}
	return e.RemoveNamedGroupingPolicies("g", rolesToRemove)

	// return a.RemoveRoleFromUser(sub, role, domain)
}

func (a *Auth) Enforce(sub string, domain string, act string) (result bool, err error) {
	e := a.enforcer
	var roleList []string
	result = false
	if roleList, err = e.GetImplicitRolesForUser(sub, domain); err != nil {
		return false, err
	}
	for k := range roleList {
		if roleList[k] == OWNER_ROLE {
			return true, nil
		}
		if mapRolePermission[roleList[k]][act] {
			return true, nil
		}
	}
	return false, nil
}

func (a *Auth) GetRolesForUser(sub string, domain string) (roleList []string, err error) {
	e := a.enforcer
	if roleList, err = e.GetRolesForUser(sub, domain); err != nil {
		return roleList, err
	}
	return roleList, nil
}

func (a *Auth) GetImplicitRolesForUser(sub string, domain string) (roleList []string, err error) {
	e := a.enforcer
	if roleList, err = e.GetImplicitRolesForUser(sub, domain); err != nil {
		return roleList, err
	}
	return roleList, nil
}

func (a *Auth) AddRoleToDomain(sub string, role string, domain string) (result bool, err error) {
	e := a.enforcer
	if _, ok := mapScope[role]; !ok {
		return false, errors.New("Role not found in scope")
	}
	return e.AddNamedGroupingPolicy("g", sub, role, domain)
}

func (a *Auth) AddRolesToDomain(sub string, roles []string, domain string) (result bool, err error) {
	e := a.enforcer
	rolesToAdd := [][]string{}
	for k := range roles {
		if _, ok := mapScope[roles[k]]; !ok {
			return false, errors.New("Role not found in scope")
		}
		rolesToAdd = append(rolesToAdd, []string{sub, roles[k], domain})
	}
	return e.AddNamedGroupingPolicies("g", rolesToAdd)
}

func (a *Auth) GetUsersForRoleInDomain(role string, shopID uint) (result []string) {
	e := a.enforcer
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return e.GetUsersForRoleInDomain(role, domain)
}

func (a *Auth) GetAllUsersByDomain(shopID uint) (result []string) {
	e := a.enforcer
	domain := fmt.Sprintf("shop_%s", strconv.Itoa(int(shopID)))
	return e.GetAllUsersByDomain(domain)
}

func (a *Auth) DeleteDomains(shopIDs []uint) (result bool, err error) {
	e := a.enforcer
	var domains []string
	for k := range shopIDs {
		domains = append(domains, fmt.Sprintf("shop_%s", strconv.Itoa(int(shopIDs[k]))))
	}
	return e.DeleteDomains(domains...)
}

// func (a *Auth) RemoveRoleFromUser(sub string, role string, domain string) (result bool, err error) {
// 	e := a.enforcer
// }
