import {createApi, fetchBaseQuery} from "@reduxjs/toolkit/query/react"
import {Endpoints} from "@src/constants/endpoints"
import {
  CreatePermissionDTO,
  GetByRoleIdDTO,
  PermissionsDTO,
} from "@src/model/generated/permission.dto"
import {PermissionCategoryVal} from "@src/model/generated/permission.entity"

export const permissionsApi = createApi({
  reducerPath: "permissions",
  baseQuery: fetchBaseQuery({
    baseUrl: Endpoints.PERMISSION,
  }),
  endpoints: (builder) => ({
    createPermission: builder.mutation<PermissionsDTO, CreatePermissionDTO>({
      query: (body) => ({
        url: "",
        method: "POST",
        body,
      }),
    }),
    getPermissionCategories: builder.query<PermissionCategoryVal[], void>({
      query: () => ({
        url: "/categories",
        method: "GET",
      }),
    }),
    getPermissions: builder.query<PermissionsDTO[], GetByRoleIdDTO>({
      query: ({roleId}) => ({
        url: "",
        method: "GET",
        params: {roleId},
      }),
    }),
    deletePermission: builder.mutation<void, Pick<PermissionsDTO, "id">>({
      query: ({id}) => ({
        url: `/${id}`,
        method: "DELETE",
      }),
    }),

    updatePermissionById: builder.mutation<
      PermissionsDTO,
      {
        id: PermissionsDTO["id"]
        permission: Partial<Pick<PermissionsDTO, "write" | "read">>
      }
    >({
      query: ({id, permission}) => ({
        url: "/" + id,
        method: "PATCH",
        body: permission,
      }),
    }),
  }),
})

export const {
  useGetPermissionCategoriesQuery,
  useGetPermissionsQuery,
  useDeletePermissionMutation,
  useCreatePermissionMutation,
  useUpdatePermissionByIdMutation,
} = permissionsApi
