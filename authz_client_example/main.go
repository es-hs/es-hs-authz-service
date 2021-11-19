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

	pb "authz-service/authz"
	"authz-service/package/casbinhelper"

	"google.golang.org/grpc"
)

const (
	// address     = "localhost:50051"
	// address     = "3.0.95.112:50051"
	address     = "authz.gempages.xyz:50051"
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
	log.Println(r4)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
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
}
