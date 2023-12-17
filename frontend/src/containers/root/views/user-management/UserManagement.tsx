import {
  DeleteOutlined,
  GroupOutlined,
  MailOutlined,
  MoreOutlined,
  UserAddOutlined,
  UserOutlined,
} from "@ant-design/icons"
import {useDeleteByIdMutation, useFetchUsersQuery} from "@api/users.api"
import {useTitle} from "@hooks/use-title"
import {useBreadcrumbs} from "@src/hooks/use-breadcrumbs"
import {UserWithRoleAndPermissions} from "@src/model/generated/service-user-dto"
import {Button, Card, Dropdown, Input, Modal, Space, Table, Tag} from "antd"
import {ColumnsType} from "antd/es/table"
import {debounce} from "lodash"
import {ChangeEvent, useCallback, useRef, useState} from "react"
import {useTranslation} from "react-i18next"
import {Link} from "react-router-dom"
import {CreateUser, CreateUserRef} from "./CreateUser"

const {confirm} = Modal
const DEFAULT_PAGE_SIZE = 5

export default function UserManagement() {
  const createUserRef = useRef<CreateUserRef>(null)
  const [usersQueryParams, setUsersQueryParams] = useState({
    limit: DEFAULT_PAGE_SIZE,
    page: 1,
    filter: "",
  })
  const {t} = useTranslation()
  const usersQuery = useFetchUsersQuery(usersQueryParams)
  const [deleteByIdMutation] = useDeleteByIdMutation()

  useTitle("userManagement.title")
  useBreadcrumbs([
    {title: "userManagement.title"},
    {title: "userManagement.userList"},
  ])

  const setFilter = (filter: string) =>
    setUsersQueryParams((prev) => ({...prev, page: 1, filter: filter}))

  const handleFilterChange = useCallback(
    debounce((e: ChangeEvent<HTMLInputElement>) => {
      setFilter(e.target.value || "")
    }, 500),
    [setFilter],
  )

  const handleUserCreated = useCallback(() => {
    usersQuery.refetch()
  }, [])

  const handleDeleteClick = (row: UserWithRoleAndPermissions) => {
    confirm({
      title: t("userManagement.deleteUser"),
      content: t("userManagement.deleteUserConfirmation"),
      onOk: async () => {
        await deleteByIdMutation({id: row.id})
        await usersQuery.refetch()
      },
    })
  }

  const handleCreateUserClick = useCallback(() => {
    createUserRef.current?.show()
  }, [createUserRef])

  const columns: ColumnsType<UserWithRoleAndPermissions> = [
    {
      title: t("general.fullName"),
      dataIndex: "fullName",
      key: "fullName",
      ellipsis: true,
      render: (_value, record) => {
        return (
          <Link
            key={record.firstName + record.lastName}
            to={record.id as string}>
            <span>
              <UserOutlined className='mr-2' />
              {record.firstName} {record.lastName}
            </span>
          </Link>
        )
      },
    },
    {
      title: t("general.email"),
      key: "email",
      dataIndex: "email",
      render: (_value, record) => {
        return (
          <span>
            <MailOutlined className='mr-2' />
            {record.email}
          </span>
        )
      },
    },
    {
      title: t("general.role"),
      dataIndex: "roleId",
      key: "roleId",
      ellipsis: true,
      render: (_, record) => {
        return (
          <Tag>
            <GroupOutlined className='mr-2' />
            {record.role?.name}
          </Tag>
        )
      },
    },
    {
      title: "",
      render: (_, record) => {
        return (
          <Dropdown
            menu={{
              items: [
                {
                  key: "delete",
                  label: (
                    <Space onClick={() => handleDeleteClick(record)}>
                      <DeleteOutlined />
                      <span>{t("general.delete")}</span>
                    </Space>
                  ),
                },
              ],
            }}>
            <a>
              <Space>
                <MoreOutlined />
              </Space>
            </a>
          </Dropdown>
        )
      },
    },
  ]

  return (
    <div className='w-full h-full'>
      <Space direction='vertical'>
        <Card bodyStyle={{display: "flex"}}>
          <Input
            type='text'
            className='flex-1'
            onChange={handleFilterChange}
            placeholder={t("general.filter")}
          />
          <span className='flex-1' />
          <Button onClick={handleCreateUserClick}>
            <Space>
              <UserAddOutlined />
              <span>{t("userManagement.createUser")}</span>
            </Space>
          </Button>
        </Card>
        <Card className='w-full h-full' bodyStyle={{padding: 0}}>
          <Table
            columns={columns}
            dataSource={usersQuery.data?.data || []}
            loading={usersQuery.isLoading || usersQuery.isFetching}
            pagination={{
              defaultPageSize:
                usersQuery.data?.meta?.perPage ?? DEFAULT_PAGE_SIZE,
              total: usersQuery.data?.meta?.total ?? 0,
              current: usersQueryParams.page,
              onChange: (page, pageSize) => {
                setUsersQueryParams((prev) => ({
                  ...prev,
                  page: page,
                  limit: pageSize ?? DEFAULT_PAGE_SIZE,
                }))
              },
              pageSizeOptions: ["10", "20", "50", "100"],
            }}
          />
        </Card>
      </Space>
      <CreateUser ref={createUserRef} onUserCreated={handleUserCreated} />
    </div>
  )
}
