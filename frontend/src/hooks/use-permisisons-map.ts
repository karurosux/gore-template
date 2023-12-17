import {useGetAuthenticatedUserQuery} from "@api/authenticated-user.api"
import {useEffect, useState} from "react"

export function usePermissionsMap(delay = 0) {
  const [loaded, setLoaded] = useState(false)
  const authenticatedUserQuery = useGetAuthenticatedUserQuery()
  const [permissionsMap, setPermissionsMap] = useState<Record<string, any>>(
    {} as any,
  )

  useEffect(() => {
    if (!authenticatedUserQuery.data) {
      return
    }

    setPermissionsMap(
      authenticatedUserQuery.data?.role.permissions.reduce(
        (acum: any, permission: any) => {
          return {...acum, [permission.category]: permission}
        },
        {},
      ),
    )

    setTimeout(() => setLoaded(true), delay)
  }, [authenticatedUserQuery.data])

  return {permissionsMap, loaded}
}
