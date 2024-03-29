// Code generated by tygo. DO NOT EDIT.

//////////
// source: branch_dto.go

export interface BranchDTO {
  id: string /* uuid */;
  name: string;
  city: string;
  state: string;
  createdAt: string /* RFC3339 */;
  updatedAt: string /* RFC3339 */;
}

//////////
// source: create_branch_dto.go

export interface CreateBranchDTO {
  name: string;
  city: string;
  state: string;
  country: string;
  zipcode: string;
}

//////////
// source: delete_branch_dto.go

export interface DeleteBranchDto {
  id: string /* uuid */;
}

//////////
// source: screate_branch_dto.go

export interface SCreateBranchDTO {
  name: string;
  city: string;
  state: string;
  zipcode: string;
  country: string;
}
