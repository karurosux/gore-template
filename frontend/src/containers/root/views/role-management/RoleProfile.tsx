import {DeleteOutlined, PlusOutlined} from "@ant-design/icons"
import {useGetAuthenticatedUserQuery} from "@src/api/authenticated-user.api"
import {
  useCreatePermissionMutation,
  useDeletePermissionMutation,
  useGetPermissionCategoriesQuery,
  useGetPermissionsQuery,
  useUpdatePermissionByIdMutation,
} from "@src/api/permissions.api"
import {useFetchRoleQuery} from "@src/api/roles.api"
import {useBreadcrumbs} from "@src/hooks/use-breadcrumbs"
import {useTitle} from "@src/hooks/use-title"
import {CreatePermissionDTO} from "@src/model/generated/permission.dto"
import {
  Alert,
  Button,
  Card,
  Checkbox,
  Form,
  Modal,
  Select,
  Space,
  Table,
  Tooltip,
} from "antd"
import {useForm} from "antd/es/form/Form"
import {ColumnsType} from "antd/es/table"
import {startCase} from "lodash"
import {useCallback} from "react"
import {useTranslation} from "react-i18next"
import {useParams} from "react-router-dom"

type PermissionForm = Omit<CreatePermissionDTO, "roleId">

export default function RoleProfile() {
  const [form] = useForm<PermissionForm>()
  const params = useParams()
  const {t} = useTranslation()
  const roleQuery = useFetchRoleQuery({id: params.roleId as string})
  const permissionCategoryQuery = useGetPermissionCategoriesQuery()
  const permissionsQuery = useGetPermissionsQuery({
    roleId: params.roleId as string,
  })
  const [deletePermissionMutation] = useDeletePermissionMutation()
  const getAuthenticatedUserQuery = useGetAuthenticatedUserQuery()
  const [createPermissionMutation, createPermissionMutationData] =
    useCreatePermissionMutation()
  const [updateById] = useUpdatePermissionByIdMutation()

  useTitle("roleManagement.roleProfileTitle", roleQuery.data)

  useBreadcrumbs([
    {
      title: "roleManagement.title",
      href: -1,
      useNavitagion: true,
    },
    {
      title: "roleManagement.roleList",
      href: -1,
      useNavitagion: true,
    },
    {
      title: roleQuery.data?.name || "...",
    },
  ])

  const handleDeletePermissionClicked = useCallback(
    (id: string) => () => {
      Modal.confirm({
        title: t("roleManagement.deletePermissionConfirmTitle"),
        content: t("roleManagement.deletePermissionConfirmContent"),
        okText: t("general.delete"),
        onOk: async () => {
          await deletePermissionMutation({id})
          await permissionsQuery.refetch()
          await getAuthenticatedUserQuery.refetch()
        },
      })
    },
    [deletePermissionMutation],
  )

  const handleUpdatePermissionClicked = useCallback(
    (id: string, prop: "write" | "read", currentVal: boolean, record: any) =>
      async () => {
        await updateById({
          id,
          permission: {
            ...{write: record.write, read: record.read},
            [prop]: !currentVal,
          },
        })
        await getAuthenticatedUserQuery.refetch()
        await permissionsQuery.refetch()
      },
    [updateById],
  )

  const handleCreatePermission = useCallback(
    async (values: PermissionForm) => {
      await createPermissionMutation({
        ...values,
        roleId: params.roleId as string,
        read: !!values.read,
        write: !!values.write,
      })
      await permissionsQuery.refetch()
      await getAuthenticatedUserQuery.refetch()
      form.resetFields()
    },
    [createPermissionMutation],
  )

  const permissionsColumns: ColumnsType<any> = [
    {
      title: t("general.category"),
      dataIndex: "category",
      key: "category",
      render: (val) => startCase(val),
    },
    {
      title: t("general.write"),
      dataIndex: "write",
      key: "write",
      render: (val, record) => (
        <Checkbox
          checked={val}
          onChange={handleUpdatePermissionClicked(
            record.id,
            "write",
            val,
            record,
          )}
        />
      ),
    },
    {
      title: t("general.read"),
      dataIndex: "read",
      key: "read",
      render: (val, record) => (
        <Checkbox
          checked={val}
          onChange={handleUpdatePermissionClicked(
            record.id,
            "read",
            val,
            record,
          )}
        />
      ),
    },
    {
      title: "",
      dataIndex: "",
      key: "actions",
      render: (_, record) => (
        <Tooltip title={t("general.delete")}>
          <Button
            type='link'
            onClick={handleDeletePermissionClicked(record.id)}>
            <DeleteOutlined />
          </Button>
        </Tooltip>
      ),
    },
  ]

  return (
    <div className='w-full h-full'>
      <Space direction='vertical' className='w-full'>
        {createPermissionMutationData.error && (
          <Alert
            type='error'
            message={t("roleManagement.errorCreatingPermission")}
          />
        )}
        <Card>
          <Form<PermissionForm>
            layout='inline'
            form={form}
            onFinish={handleCreatePermission}>
            <Form.Item<PermissionForm>
              name='category'
              label={t("general.category")}
              rules={[{required: true}]}
              className='min-w-[256px]'>
              <Select>
                {permissionCategoryQuery.data?.map((category) => (
                  <Select.Option key={category} value={category}>
                    {startCase(category)}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
            <Form.Item<PermissionForm>
              name='write'
              label={t("general.write")}
              valuePropName='checked'>
              <Checkbox />
            </Form.Item>
            <Form.Item<PermissionForm>
              name='read'
              label={t("general.read")}
              valuePropName='checked'>
              <Checkbox />
            </Form.Item>
            <Form.Item>
              <Button
                htmlType='submit'
                icon={<PlusOutlined />}
                ghost={false}
                disabled={createPermissionMutationData.isLoading}
                loading={createPermissionMutationData.isLoading}
                onClick={() => form.submit()}>
                {t("roleManagement.addPermission")}
              </Button>
            </Form.Item>
            <Form.Item>
              <Button
                ghost
                htmlType='reset'
                danger
                onClick={() => form.resetFields()}>
                {t("general.reset")}
              </Button>
            </Form.Item>
          </Form>
        </Card>
        <Card
          bodyStyle={{padding: 0}}
          title={t("general.permissions", roleQuery.data as any) as string}>
          <Table
            columns={permissionsColumns}
            dataSource={permissionsQuery.data ?? []}
            loading={permissionsQuery.isLoading}
            pagination={false}
          />
        </Card>
      </Space>
    </div>
  )
}
