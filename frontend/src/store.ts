import {usersApi} from "@api/users.api"
import {authenticatedUserApi} from "./api/authenticated-user.api"
import {configureStore} from "@reduxjs/toolkit"
import {TypedUseSelectorHook, useDispatch, useSelector} from "react-redux"
import {breadcrumbSlice} from "@slices/breadcrumbs.slice.ts"
import {rolesApi} from "./api/roles.api"
import {permissionsApi} from "./api/permissions.api"

export const store = configureStore({
  reducer: {
    [breadcrumbSlice.name]: breadcrumbSlice.reducer,
    [usersApi.reducerPath]: usersApi.reducer,
    [authenticatedUserApi.reducerPath]: authenticatedUserApi.reducer,
    [rolesApi.reducerPath]: rolesApi.reducer,
    [permissionsApi.reducerPath]: permissionsApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware()
      .concat(authenticatedUserApi.middleware)
      .concat(usersApi.middleware)
      .concat(rolesApi.middleware)
      .concat(permissionsApi.middleware),
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
