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
