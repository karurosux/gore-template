import {createApi, fetchBaseQuery} from "@reduxjs/toolkit/query/react"
import {Endpoints} from "@src/constants/endpoints"
import {WithFilterRequestDTO} from "@src/model/generated/base.dto"
import {
  CreateRoleDTO,
  RoleDTO,
  RoleWithBranchDTO,
} from "@src/model/generated/role.dto"

export const rolesApi = createApi({
  reducerPath: "roles",
  baseQuery: fetchBaseQuery({baseUrl: Endpoints.ROLE}),
  endpoints: (builder) => ({
    createRole: builder.mutation<RoleDTO, CreateRoleDTO>({
      query: (body) => ({
        url: "",
        method: "POST",
        body,
      }),
    }),
    fetchRoles: builder.query<
      RoleWithBranchDTO[],
      Partial<WithFilterRequestDTO>
    >({
      query: ({filter}) => {
        const params: Record<string, any> = {}

        if (filter) {
          params.filter = filter
        }

        return {
          url: "",
          params,
        }
      },
    }),
    deleteById: builder.mutation<void, Pick<RoleDTO, "id">>({
      query: ({id}) => ({
        url: `/${id}`,
        method: "DELETE",
      }),
    }),
    fetchRole: builder.query<RoleDTO, Pick<RoleDTO, "id">>({
      query: ({id}) => ({
        url: `/${id}`,
      }),
    }),
  }),
})

export const {
  useFetchRolesQuery,
  useFetchRoleQuery,
  useDeleteByIdMutation,
  useCreateRoleMutation,
} = rolesApi
