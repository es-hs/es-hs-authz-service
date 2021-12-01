/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	"authz-service/package/casbinhelper"

	pb "github.com/es-hs/erpc/authz"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// address     = "3.0.95.112:50051"
	// address     = "authz.gempages.xyz:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthzRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	t1 := time.Now()
	r, err := c.AddRoleToDomain(ctx, &pb.AddRoleToDomainRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    casbinhelper.OWNER_ROLE,
	})
	log.Println("Result ", r.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	result, err := c.AddRolesForUserToDomain(ctx, &pb.AddRolesForUserToDomainRequest{
		UserId: 2007,
		ShopId: 2011,
		Act:    []string{casbinhelper.ADMIN_ROLE, casbinhelper.PRODUCT_DELETE, casbinhelper.SECTION_DELETE},
	})
	result, _ = c.AddRolesForUserToDomain(ctx, &pb.AddRolesForUserToDomainRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    []string{casbinhelper.ADMIN_ROLE, casbinhelper.PRODUCT_DELETE, casbinhelper.SECTION_DELETE},
	})
	result, _ = c.AddRolesForUserToDomain(ctx, &pb.AddRolesForUserToDomainRequest{
		UserId: 2009,
		ShopId: 2011,
		Act:    []string{casbinhelper.ADMIN_ROLE, casbinhelper.PRODUCT_DELETE, casbinhelper.SECTION_DELETE},
	})
	log.Println("Result ", result.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//check roles list
	r2, err := c.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r2.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//check roles detail list
	r3, err := c.GetImplicitRolesInDomain(ctx, &pb.GetImplicitRolesInDomainRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r3.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//check roles detail list
	r4, err := c.CheckPermission(ctx, &pb.CheckPermissionRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    casbinhelper.LOGIN_PERMISSION,
	})
	log.Println(r4.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("get all user in domain ")
	r10, err := c.GetAllUsersByDomain(ctx, &pb.GetAllUsersByDomainRequest{
		ShopId: 2011,
	})
	log.Println("Result ", r10.UserIds)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("get all user by roles in domain ")

	r11, err := c.GetUsersForRoleInDomain(ctx, &pb.GetUsersForRoleInDomainRequest{
		ShopId: 2011,
		Role:   casbinhelper.ADMIN_ROLE,
	})
	log.Println("Result ", r11.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r7, err := c.RemoveRoleFromDomain(ctx, &pb.RemoveRoleFromDomainRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    casbinhelper.OWNER_ROLE,
	})
	log.Println("Result ", r7.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r8, err := c.RemoveRolesFromDomain(ctx, &pb.RemoveRolesFromDomainRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    []string{casbinhelper.ADMIN_ROLE, casbinhelper.PRODUCT_DELETE, casbinhelper.SECTION_DELETE},
	})
	log.Println("Result ", r8.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("after remove admin and product delete seciton delete role")
	//SECOND TIME AFTER REMOVE ROLES
	//check roles list
	r2, err = c.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r2.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//check roles detail list
	r3, err = c.GetImplicitRolesInDomain(ctx, &pb.GetImplicitRolesInDomainRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r3.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//check roles detail list
	r4, err = c.CheckPermission(ctx, &pb.CheckPermissionRequest{
		UserId: 2008,
		ShopId: 2011,
		Act:    casbinhelper.LOGIN_PERMISSION,
	})
	log.Println(r4.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//
	//gemerate role for
	r5, err := c.GenerateOwnerRole(ctx, &pb.GenerateOwnerRoleRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r5.Code)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//remove domain
	log.Println("let  remove domains 2011")
	//SECOND TIME AFTER REMOVE ROLES
	//check roles list
	r12, err := c.DeleteDomains(ctx, &pb.DeleteDomainsRequest{
		ShopIds: []int64{2011},
	})
	log.Println(r12.Result)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//
	log.Println("after remove domains")
	//SECOND TIME AFTER REMOVE ROLES
	//check roles list
	r2, err = c.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
		UserId: 2008,
		ShopId: 2011,
	})
	log.Println(r2.Roles)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}
