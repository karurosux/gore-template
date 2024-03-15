import {
  BuildOutlined,
  CheckCircleFilled,
  DeleteOutlined,
  GroupOutlined,
  MoreOutlined,
} from "@ant-design/icons"
import {useDeleteByIdMutation, useFetchRolesQuery} from "@src/api/roles.api"
import {useBreadcrumbs} from "@src/hooks/use-breadcrumbs"
import {useTitle} from "@src/hooks/use-title"
import {Button, Card, Dropdown, Input, Modal, Space, Table} from "antd"
import {ColumnsType} from "antd/es/table"
import {debounce} from "lodash"
import {useCallback, useRef, useState} from "react"
import {useTranslation} from "react-i18next"
import {Link} from "react-router-dom"
import {CreateRole, CreateRoleRef} from "./CreateRole"
import {RoleWithBranchDTO} from "@src/model/generated/role.dto"

export default function RoleManagement() {
  const {t} = useTranslation()
  const createRoleRef = useRef<CreateRoleRef>(null)
  const [rolesQueryParams, setRoleQueryParams] = useState({filter: ""})
  const fetchRolesQuery = useFetchRolesQuery(rolesQueryParams)
  const [deleteRoleById] = useDeleteByIdMutation()

  useTitle("roleManagement.title")

  useBreadcrumbs([
    {
      title: "roleManagement.title",
    },
    {
      title: "roleManagement.roleList",
    },
  ])

  const handleFilterChange = useCallback(
    debounce((e: React.ChangeEvent<HTMLInputElement>) => {
      setRoleQueryParams({...rolesQueryParams, filter: e.target.value})
    }, 500),
    [],
  )

  const handleDeleteRoleClicked = useCallback(
    (id: string) => {
      Modal.confirm({
        title: t("roleManagement.deleteRole"),
        content: t("roleManagement.deleteRoleConfirmation"),
        okText: t("general.delete"),
        onOk: async () => {
          await deleteRoleById({id})
          await fetchRolesQuery.refetch()
        },
      })
    },
    [deleteRoleById, Modal.confirm, fetchRolesQuery.refetch],
  )

  const columns: ColumnsType<RoleWithBranchDTO> = [
    {
      title: t("general.name"),
      dataIndex: "name",
      key: "name",
      render: (val, record) => (
        <Link to={record.id as string}>
          <Button type='link'>
            <Space>
              <GroupOutlined />
              <span>{val}</span>
            </Space>
          </Button>
        </Link>
      ),
    },
    {
      title: t("roleManagement.isSuperAdmin"),
      dataIndex: "roleType",
      key: "roleType",
      render: (val) => val === "SUPER_ADMIN" && <CheckCircleFilled />,
    },
    {
      title: t("general.branch"),
      dataIndex: "branch.name",
      key: "branch.name",
      render: (_, record) => (
        <Space>
          <BuildOutlined />
          <span>{record.branch?.name}</span>
        </Space>
      ),
    },
    {
      title: "",
      dataIndex: "action",
      key: "action",
      render: (_, record) => (
        <Dropdown
          menu={{
            items: [
              {
                key: "delete",
                label: (
                  <Space onClick={() => handleDeleteRoleClicked(record.id)}>
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
      ),
    },
  ]

  return (
    <div className='w-full h-full'>
      <Space direction='vertical' className='w-full'>
        <Card bodyStyle={{display: "flex"}}>
          <Input
            type='text'
            className='flex-1'
            placeholder={t("general.filter")}
            onChange={handleFilterChange}
          />
          <span className='flex-1' />
          <Button onClick={() => createRoleRef.current?.show()}>
            <Space>
              <GroupOutlined />
              <span>{t("roleManagement.createRole")}</span>
            </Space>
          </Button>
        </Card>
        <Card className='w-full h-full' bodyStyle={{padding: 0}}>
          <Table
            loading={fetchRolesQuery.isLoading}
            columns={columns}
            dataSource={fetchRolesQuery.data ?? []}
            pagination={false}
          />
        </Card>
      </Space>
      <CreateRole
        ref={createRoleRef}
        onRoleCreated={() => fetchRolesQuery.refetch()}
      />
    </div>
  )
}
