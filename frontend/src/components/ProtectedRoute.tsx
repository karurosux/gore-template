import {PropsWithChildren} from "react"
import {Navigate, useLocation} from "react-router-dom"
import {useGetAuthenticatedUserQuery} from "../api/authenticated-user.api"
import {General} from "../constants/general"
import {RouteLoadingSuspense} from "./RouteLoadingSuspense"

type Props = PropsWithChildren

export function ProtectedRoute({children}: Props) {
  const location = useLocation()
  const authenticatedUserQuery = useGetAuthenticatedUserQuery()

  if (authenticatedUserQuery?.isLoading) {
    return <span />
  }

  return (
    <RouteLoadingSuspense>
      {authenticatedUserQuery?.isSuccess ? (
        children
      ) : (
        <Navigate to={General.DEFAULT_PUBLIC_ROUTE} state={{from: location}} />
      )}
    </RouteLoadingSuspense>
  )
}
