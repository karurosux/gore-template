import {useFetchRolesQuery} from "@api/roles.api"
import {
  useCreateUserMutation,
  useExistByEmailMutation,
} from "@src/api/users.api"
import {CreateUserDTO} from "@src/model/generated/user.dto"
import {Form, Input, Modal, Select} from "antd"
import {debounce} from "lodash"
import React, {useCallback, useImperativeHandle, useState} from "react"
import {useTranslation} from "react-i18next"

export type CreateUserRef = {
  show: () => void
}

type Props = {
  onUserCreated?: () => void
}

export const CreateUser = React.forwardRef<CreateUserRef, Props>(
  function CreateUser(props, ref) {
    const [form] = Form.useForm<CreateUserDTO>()
    const [createUserMutation, createUserMutationData] = useCreateUserMutation()
    const rolesQuery = useFetchRolesQuery({})
    const [emailExistsMutation] = useExistByEmailMutation()
    const [isOpen, setIsOpen] = useState(false)
    const {t} = useTranslation()

    const doesEmailExist = useCallback(
      debounce((email: string, callback: (exist: boolean) => void) => {
        emailExistsMutation({email}).then((result) => {
          callback((result as {data: boolean}).data)
        })
      }, 500),
      [emailExistsMutation],
    )

    useImperativeHandle(ref, () => ({
      show: () => {
        setIsOpen(true)
      },
    }))

    const handleOkClick = useCallback(() => form.submit(), [setIsOpen, form])

    const handleSubmit = useCallback(
      (value: CreateUserDTO) => {
        createUserMutation(value).then(() => {
          form.resetFields()
          setIsOpen(false)
          props?.onUserCreated?.()
        })
      },
      [setIsOpen, props, form],
    )

    const handleCancelClick = useCallback(() => {
      setIsOpen(false)
    }, [setIsOpen])

    return (
      <Modal
        destroyOnClose
        title={t("userManagement.createUser")}
        okText={t("general.create")}
        okType='primary'
        open={isOpen}
        confirmLoading={createUserMutationData.isLoading}
        okButtonProps={{
          htmlType: "submit",
        }}
        onOk={handleOkClick}
        onCancel={handleCancelClick}>
        <div className='pt-2'>
          <Form
            layout='vertical'
            preserve={false}
            form={form}
            onFinish={handleSubmit}>
            <Form.Item<CreateUserDTO>
              name='email'
              label={t("general.email")}
              rules={[
                {required: true, type: "email"},
                {
                  validator(_rule, value, callback) {
                    if (value) {
                      doesEmailExist(value, (exist) => {
                        if (exist) {
                          callback(t("userManagement.emailExist"))
                        } else {
                          callback()
                        }
                      })
                    } else {
                      callback()
                    }
                  },
                },
              ]}>
              <Input />
            </Form.Item>
            <Form.Item<CreateUserDTO>
              name='firstName'
              label={t("general.firstName")}
              rules={[{required: true}]}>
              <Input />
            </Form.Item>
            <Form.Item<CreateUserDTO>
              name='lastName'
              label={t("general.lastName")}
              rules={[{required: true}]}>
              <Input />
            </Form.Item>
            <Form.Item<CreateUserDTO>
              name='roleId'
              label={t("general.role")}
              rules={[{required: true}]}>
              <Select loading={rolesQuery.isLoading}>
                {rolesQuery.data?.map((role) => (
                  <Select.Option key={role.id} value={role.id}>
                    {role.name}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Form>
        </div>
      </Modal>
    )
  },
)
