package permissionsdto

type PatchPermissionDTO struct {
	PermissionByIdDTO `tstype:",extends,required"`
	Write             bool `json:"write"`
	Read              bool `json:"read"`
}
