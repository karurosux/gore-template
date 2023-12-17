import {Endpoints} from "@constants/endpoints"
import {createApi, fetchBaseQuery} from "@reduxjs/toolkit/query/react"
import {PaginatedRequest} from "@src/model/generated/api-dto"
import {CreateUserDTO} from "@src/model/generated/api-user-dto"
import {Paginated} from "@src/model/generated/model"
import {
  UserDTO,
  UserWithRoleAndPermissions,
} from "@src/model/generated/service-user-dto"

export const usersApi = createApi({
  reducerPath: "users",
  baseQuery: fetchBaseQuery({baseUrl: Endpoints.USER}),
  endpoints: (builder) => ({
    createUser: builder.mutation<UserDTO, CreateUserDTO>({
      query: (body) => ({
        url: "",
        method: "POST",
        body,
      }),
    }),
    fetchUsers: builder.query<
      Paginated<UserWithRoleAndPermissions>,
      PaginatedRequest
    >({
      query: (params) => ({
        url: "",
        params,
      }),
    }),
    fetchUserById: builder.query<UserDTO, Pick<UserDTO, "id">>({
      query: ({id}) => ({
        method: "GET",
        url: `/${id}`,
      }),
    }),
    deleteById: builder.mutation<void, Pick<UserDTO, "id">>({
      query: ({id}) => ({
        url: `/${id}`,
        method: "DELETE",
      }),
    }),
    existByEmail: builder.mutation<boolean, Pick<UserDTO, "email">>({
      query: ({email}) => ({
        method: "GET",
        url: "/exist-by-email/" + email,
      }),
    }),
  }),
})

export const {
  useFetchUsersQuery,
  useDeleteByIdMutation,
  useCreateUserMutation,
  useExistByEmailMutation,
  useFetchUserByIdQuery,
} = usersApi
