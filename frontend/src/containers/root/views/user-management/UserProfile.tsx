import {useFetchUserByIdQuery} from "@src/api/users.api"
import {useBreadcrumbs} from "@src/hooks/use-breadcrumbs"
import {useTitle} from "@src/hooks/use-title"
import {Card} from "antd"
import {useParams} from "react-router-dom"

export default function UserProfile() {
  const params = useParams()
  const fetchUserByIdQuery = useFetchUserByIdQuery({
    id: params.userId as string,
  })

  useTitle(
    fetchUserByIdQuery.isLoading
      ? "..."
      : fetchUserByIdQuery.data?.firstName +
          " " +
          fetchUserByIdQuery.data?.lastName,
  )

  useBreadcrumbs([
    {
      title: "userManagement.title",
      href: -1,
      useNavitagion: true,
    },
    {
      title: "userManagement.userList",
      href: -1,
      useNavitagion: true,
    },
    {
      title: fetchUserByIdQuery.isLoading
        ? "..."
        : (((fetchUserByIdQuery.data?.firstName as string) +
            " " +
            fetchUserByIdQuery.data?.lastName) as string),
    },
  ])

  return (
    <div>
      {fetchUserByIdQuery.isLoading ? (
        <div>Loading....</div>
      ) : (
        <Card>
          {fetchUserByIdQuery.data?.firstName +
            " " +
            fetchUserByIdQuery.data?.lastName}
        </Card>
      )}
    </div>
  )
}
