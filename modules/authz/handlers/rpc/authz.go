package rpc

import (
	"authz-service/modules/authz/usecases"
	"authz-service/package/casbinhelper"
	"context"

	pb "github.com/es-hs/erpc/authz"

	"google.golang.org/grpc"
)

type AuthzRpcServer struct {
	pb.UnimplementedAuthzServer
	usecase usecases.AuthzUsecase
}

func RegisterServer(server *grpc.Server, authz casbinhelper.Auth) {
	pb.RegisterAuthzServer(server, &AuthzRpcServer{
		usecase: usecases.NewAuthzUsecase(authz),
	})
}

func (instance *AuthzRpcServer) AddRoleToDomain(ctx context.Context, request *pb.AddRoleToDomainRequest) (*pb.AddRoleToDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	act := request.GetAct()
	result, err := instance.usecase.AddRoleForUserToDomain(userID, shopID, act)
	if err != nil {
		return nil, err
	}
	return &pb.AddRoleToDomainResult{
		Result: result,
		Code:   0,
	}, nil
}

func (instance *AuthzRpcServer) CheckPermission(ctx context.Context, request *pb.CheckPermissionRequest) (*pb.CheckPermissionResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	act := request.GetAct()
	result, err := instance.usecase.CheckPermission(userID, shopID, act)
	if err != nil {
		return nil, err
	}
	return &pb.CheckPermissionResult{
		Result: result,
		Code:   0,
	}, nil
}

func (instance *AuthzRpcServer) GetRolesInDomain(ctx context.Context, request *pb.GetRolesInDomainRequest) (*pb.GetRolesInDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	result, err := instance.usecase.GetRolesInDomain(userID, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.GetRolesInDomainResult{
		Roles: result,
		Code:  0,
	}, nil
}

func (instance *AuthzRpcServer) GetImplicitRolesInDomain(ctx context.Context, request *pb.GetImplicitRolesInDomainRequest) (*pb.GetImplicitRolesInDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	result, err := instance.usecase.GetImplicitRolesInDomain(userID, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.GetImplicitRolesInDomainResult{
		Roles: result,
		Code:  0,
	}, nil

}

func (instance *AuthzRpcServer) GenerateOwnerRole(ctx context.Context, request *pb.GenerateOwnerRoleRequest) (*pb.GenerateOwnerRoleResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	err := instance.usecase.GenerateOwnerRole(userID, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateOwnerRoleResult{
		Code: 0,
	}, nil

}

func (instance *AuthzRpcServer) AddRolesForUserToDomain(ctx context.Context, request *pb.AddRolesForUserToDomainRequest) (*pb.AddRolesForUserToDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	roles := request.GetAct()
	result, err := instance.usecase.AddRolesForUserToDomain(userID, roles, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.AddRolesForUserToDomainResult{
		Result: result,
		Code:   0,
	}, nil

}

func (instance *AuthzRpcServer) RemoveRolesFromDomain(ctx context.Context, request *pb.RemoveRolesFromDomainRequest) (*pb.RemoveRolesFromDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	roles := request.GetAct()
	result, err := instance.usecase.RemoveRolesFromUser(userID, roles, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveRolesFromDomainResult{
		Result: result,
		Code:   0,
	}, nil
}

func (instance *AuthzRpcServer) RemoveRoleFromDomain(ctx context.Context, request *pb.RemoveRoleFromDomainRequest) (*pb.RemoveRoleFromDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	role := request.GetAct()
	result, err := instance.usecase.RemoveRoleFromUser(userID, role, shopID)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveRoleFromDomainResult{
		Result: result,
		Code:   0,
	}, nil
}

func (instance *AuthzRpcServer) GetUsersForRoleInDomain(ctx context.Context, request *pb.GetUsersForRoleInDomainRequest) (*pb.GetUsersForRoleInDomainResult, error) {
	role := request.GetRole()
	shopID := request.GetShopId()
	result := instance.usecase.GetUsersForRoleInDomain(role, shopID)
	return &pb.GetUsersForRoleInDomainResult{
		UserIds: result,
		Code:    0,
	}, nil
}

func (instance *AuthzRpcServer) GetAllUsersByDomain(ctx context.Context, request *pb.GetAllUsersByDomainRequest) (*pb.GetAllUsersByDomainResult, error) {
	shopID := request.GetShopId()
	result := instance.usecase.GetAllUsersByDomain(shopID)
	return &pb.GetAllUsersByDomainResult{
		UserIds: result,
		Code:    0,
	}, nil
}

func (instance *AuthzRpcServer) DeleteDomains(ctx context.Context, request *pb.DeleteDomainsRequest) (*pb.DeleteDomainsResult, error) {
	shopIds := request.GetShopIds()
	result, err := instance.usecase.DeleteDomains(shopIds)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDomainsResult{
		Result: result,
		Code:   0,
	}, nil
}
