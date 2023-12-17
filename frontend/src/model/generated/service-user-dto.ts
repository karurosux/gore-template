// Code generated by tygo. DO NOT EDIT.

import {RoleWithPermissionsDTO} from "./service-role-dto.ts"

//////////
// source: user_dto.go

export interface UserDTO {
  id: string /* uuid */
  firstName: string
  lastName: string
  email: string
  branchId: string /* uuid */
  createdAt: string /* RFC3339 */
  updatedAt: string /* RFC3339 */
}

//////////
// source: user_with_password_dto.go

export interface UserWithPasswordDTO extends UserDTO {
  password: string
}

//////////
// source: user_with_role_and_permissions_dto.go

export interface UserWithRoleAndPermissions extends UserDTO {
  role: RoleWithPermissionsDTO
}
