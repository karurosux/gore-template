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
		"time.Time":                   "string /* RFC3339 */",
		"null.String":                 "null | string",
		"null.Bool":                   "null | boolean",
		"uuid.UUID":                   "string /* uuid */",
		"uuid.NullUUID":               "string /* uuid */",
		"permissionEntity.Permission": "Permission",
		"branchEntity.Branch":         "Branch",
		"branchentity.Branch":         "Branch",
		"userentity.User":             "User",
		"roleEntity.Role":             "Role",
		"gorm.DeletedAt":              "Date",
	}
	fallbackType := "any"
	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			// Api DTOs
			{
				Path:         "backend/auth/dto",
				OutputPath:   getOutputBase("auth.dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "backend/branch/dto",
				OutputPath:   getOutputBase("branch.dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "backend/branch/entity",
				OutputPath:   getOutputBase("branch.entity.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "backend/permission/dto",
				OutputPath:   getOutputBase("permission.dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				Frontmatter:  "import {PermissionCategoryVal} from \"./permission.entity\"",
			},
			{
				Path:         "backend/permission/entity",
				OutputPath:   getOutputBase("permission.entity.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
			},
			{
				Path:         "backend/role/dto",
				OutputPath:   getOutputBase("role.dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				Frontmatter: `import {RoleTypeVal} from "./role.entity"
import {BranchDTO} from "./branch.dto"
import {PermissionsDTO} from "./permission.dto"`,
			},
			{
				Path:         "backend/role/entity",
				OutputPath:   getOutputBase("role.entity.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				Frontmatter: `import {Permission} from "./permission.entity"
import {Branch} from "./branch.entity"`,
			},
			{
				Path:         "backend/user/dto",
				OutputPath:   getOutputBase("user.dto.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				Frontmatter:  "import {RoleWithPermissionsDTO} from \"./role.dto\"",
			},
			{
				Path:         "backend/user/entity",
				OutputPath:   getOutputBase("user.entity.ts"),
				FallbackType: fallbackType,
				TypeMappings: typeMappings,
				Frontmatter:  "import {Branch} from \"./branch.entity\"\nimport {Role} from \"./role.entity\"",
			},
			{
				Path:         "backend/base/dto",
				OutputPath:   getOutputBase("base.dto.ts"),
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
