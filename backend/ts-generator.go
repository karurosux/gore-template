package main

import (
	"fmt"

	"github.com/gzuidhof/tygo/tygo"
)

func getOutputBase(fileName string) string {
	return "../frontend/src/model/generated/" + fileName
}

func GenerateTsTypes() {
	typeMappings := map[string]string{
		"time.Time":     "string /* RFC3339 */",
		"null.String":   "null | string",
		"null.Bool":     "null | boolean",
		"uuid.UUID":     "string /* uuid */",
		"uuid.NullUUID": "string /* uuid */",
	}
	fallbackType := "any"
	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			// Api DTOs
			{
				Path:         "app/api/dto",
				OutputPath:   getOutputBase("api-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "app/api/dto/user_dto",
				OutputPath:   getOutputBase("api-user-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "app/api/dto/role_dto",
				OutputPath:   getOutputBase("api-role-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "app/api/dto/permissions_dto",
				OutputPath:   getOutputBase("api-permissions-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "app/api/dto/auth_dto",
				OutputPath:   getOutputBase("api-auth-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			// Service DTOs
			{
				Path:         "app/service/dto/user_dto",
				OutputPath:   getOutputBase("service-user-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				ExcludeFiles: []string{
					"create_user_dto.go",
				},
				Frontmatter: `
import {RoleWithPermissionsDTO} from "./service-role-dto.ts";
				`,
			},
			{
				Path:         "app/service/dto/role_dto",
				OutputPath:   getOutputBase("service-role-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				ExcludeFiles: []string{
					"create_role_dto.go",
				},
				Frontmatter: `
import {PermissionsDTO} from "./service-permissions-dto.ts";
import {BranchDTO} from "./service-branch-dto.ts";
import {RoleTypeVal} from "./entities.ts";
				`,
			},
			{
				Path:         "app/service/dto/permissions_dto",
				OutputPath:   getOutputBase("service-permissions-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				ExcludeFiles: []string{
					"create_permission_dto.go",
				},
				Frontmatter: `
import {PermissionCategoryVal} from "./entities.ts";
				`,
			},
			{
				Path:         "app/service/dto/branch_dto",
				OutputPath:   getOutputBase("service-branch-dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "app/model",
				OutputPath:   getOutputBase("model.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				ExcludeFiles: []string{
					"controller.go",
				},
			},
			// Entities
			{
				Path:         "app/entities",
				OutputPath:   getOutputBase("entities.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
		},
	}
	gen := tygo.New(config)

	fmt.Println("Generating TypeScript definitions.")

	err := gen.Generate()

	if err != nil {
		fmt.Println("Error generating TypeScript definitions")
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("TypeScript Types generated, See %s\n", getOutputBase(""))
}
