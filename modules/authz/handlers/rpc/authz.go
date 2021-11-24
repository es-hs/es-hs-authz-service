package rpc

import (
	"authz-service/modules/authz/usecases"
	"authz-service/package/casbinhelper"
	"context"

	pb "github.com/es-hs/erpc/authz"

	"google.golang.org/grpc"
)

type AuthzRpcServer struct {
	pb.UnimplementedAuthzRPCServer
	usecase usecases.AuthzUsecase
}

func RegisterServer(server *grpc.Server, authz casbinhelper.Auth) {
	pb.RegisterAuthzRPCServer(server, &AuthzRpcServer{
		usecase: usecases.NewAuthzUsecase(authz),
	})
}

func (instance *AuthzRpcServer) AddRoleToDomain(ctx context.Context, request *pb.AddRoleToDomainRequest) (*pb.AddRoleToDomainResult, error) {
	userID := request.GetUserId()
	shopID := request.GetShopId()
	act := request.GetAct()
	result, err := instance.usecase.AddRoleForUserToDomain(uint(userID), uint(shopID), act)
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
	result, err := instance.usecase.CheckPermission(uint(userID), uint(shopID), act)
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
	result, err := instance.usecase.GetRolesInDomain(uint(userID), uint(shopID))
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
	result, err := instance.usecase.GetImplicitRolesInDomain(uint(userID), uint(shopID))
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
	err := instance.usecase.GenerateOwnerRole(uint(userID), uint(shopID))
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
	result, err := instance.usecase.AddRolesForUserToDomain(uint(userID), roles, uint(shopID))
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
	result, err := instance.usecase.RemoveRolesFromUser(uint(userID), roles, uint(shopID))
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
	result, err := instance.usecase.RemoveRoleFromUser(uint(userID), role, uint(shopID))
	if err != nil {
		return nil, err
	}
	return &pb.RemoveRoleFromDomainResult{
		Result: result,
		Code:   0,
	}, nil
}
