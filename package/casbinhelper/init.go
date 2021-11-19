package casbinhelper

import (
	"fmt"
	"os"
	"time"

	"github.com/casbin/casbin/v2"
	redisadapter "github.com/casbin/redis-adapter/v2"

	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/sirupsen/logrus"
)

func init() {
	mapRolePermission = make(map[string]map[string]bool)
	mapScope = make(map[string]bool)
	mapRolePermission[OWNER_ROLE] = map[string]bool{
		LOGIN_PERMISSION: true,
	}
	mapRolePermission[ADMIN_ROLE] = map[string]bool{
		LOGIN_PERMISSION: true,
	}
	mapRolePermission[SHOP_READ] = map[string]bool{
		SHOP_READ_PERMISSION: true,
	}
	mapRolePermission[SHOP_WRITE] = map[string]bool{
		SHOP_WRITE_PERMISSION: true,
	}
	mapRolePermission[PRODUCT_READ] = map[string]bool{
		PRODUCT_READ_PERMISSION: true,
	}
	mapRolePermission[PRODUCT_WRITE] = map[string]bool{
		PRODUCT_WRITE_PERMISSION: true,
	}
	mapRolePermission[PRODUCT_DELETE] = map[string]bool{
		PRODUCT_DELETE_PERMISSION: true,
	}
	mapRolePermission[THEME_READ] = map[string]bool{
		THEME_READ_PERMISSION: true,
	}
	mapRolePermission[THEME_WRITE] = map[string]bool{
		THEME_WRITE_PERMISSION: true,
	}
	mapRolePermission[THEME_DELETE] = map[string]bool{
		THEME_DELETE_PERMISSION: true,
	}
	mapRolePermission[SECTION_READ] = map[string]bool{
		SECTION_READ_PERMISSION: true,
	}
	mapRolePermission[SECTION_WRITE] = map[string]bool{
		SECTION_WRITE_PERMISSION: true,
	}
	mapRolePermission[SECTION_DELETE] = map[string]bool{
		SECTION_DELETE_PERMISSION: true,
	}
	for k := range mapRolePermission {
		mapScope[k] = true
	}
}
func InitAuthorization() (Auth, error) {
	t1 := time.Now()
	RedisURL := os.Getenv("REDIS_URL")

	fmt.Println("init authorization with DB")
	var err error
	adapter := redisadapter.NewAdapter("tcp", RedisURL)

	// adapter, err := gormadapter.NewAdapterByDB(driver.OpenDB())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// AuthEnforcer.InitWithAdapter("configs/policy/rbac_with_domain.conf", adapter)

	authEnforcer, err := casbin.NewEnforcer("./configs/policy/rbac_with_domain_no_obj.conf", adapter)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	logrus.Info("Init Authen Finished: ", time.Since(t1))
	return Auth{
		enforcer: authEnforcer,
	}, nil
}
