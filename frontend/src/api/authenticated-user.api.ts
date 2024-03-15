import {Endpoints} from "@constants/endpoints"
import {createApi, fetchBaseQuery} from "@reduxjs/toolkit/query/react"
import {LoginBodyDTO} from "@src/model/generated/auth.dto"
import {UserWithRoleAndPermissions} from "@src/model/generated/user.dto"

export const authenticatedUserApi = createApi({
  reducerPath: "authenticatedUser",
  baseQuery: fetchBaseQuery({baseUrl: Endpoints.BASE_V1}),
  endpoints: (builder) => ({
    login: builder.mutation<void, LoginBodyDTO>({
      query: (body) => ({
        url: "/auth/login",
        method: "POST",
        body,
      }),
    }),
    logout: builder.mutation<void, void>({
      query: () => ({
        url: "/auth/logout",
        method: "POST",
      }),
    }),
    getAuthenticatedUser: builder.query<UserWithRoleAndPermissions, void>({
      query: () => "/users/me",
    }),
  }),
})

export const {
  useGetAuthenticatedUserQuery,
  useLoginMutation,
  useLogoutMutation,
} = authenticatedUserApi
