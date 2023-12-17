import {useCreateRoleMutation} from "@src/api/roles.api"
import {Form, Input, Modal} from "antd"
import {forwardRef, useCallback, useImperativeHandle, useState} from "react"
import {useTranslation} from "react-i18next"

export type CreateRoleRef = {
  show: () => void
}

type Props = {
  onRoleCreated?: () => void
}

type CreateRoleForm = {
  name: string
}

export const CreateRole = forwardRef<CreateRoleRef, Props>(
  function CreateRole(props, ref) {
    const [form] = Form.useForm<CreateRoleForm>()
    const {t} = useTranslation()
    const [open, setOpen] = useState(false)
    const [createRole, creteRoleData] = useCreateRoleMutation()

    useImperativeHandle(ref, () => ({
      show: () => setOpen(true),
    }))

    const handleClose = useCallback(() => {
      setOpen(false)
    }, [setOpen])

    const handleSubmit = useCallback(
      async (values: CreateRoleForm) => {
        await createRole(values)
        props.onRoleCreated?.()
        handleClose()
      },
      [props.onRoleCreated],
    )

    return (
      <Modal
        destroyOnClose
        title={t("roleManagement.createRole")}
        open={open}
        onOk={() => form.submit()}
        onCancel={handleClose}
        confirmLoading={creteRoleData.isLoading}
        okText={t("general.create")}>
        <Form
          layout='vertical'
          preserve={false}
          form={form}
          onFinish={handleSubmit}>
          <Form.Item<CreateRoleForm>
            name='name'
            label={t("general.name")}
            rules={[{required: true}]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    )
  },
)
