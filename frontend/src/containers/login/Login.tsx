import {useTitle} from "@hooks/use-title"
import {Alert, Button, Card, Form, Input, Space} from "antd"
import {useCallback} from "react"
import {useTranslation} from "react-i18next"
import {useNavigate} from "react-router-dom"
import {useLoginMutation} from "@api/authenticated-user.api"
import logo from "@assets/react.svg"
import {General} from "@constants/general"
import {
  FADE_IN_INITIAL_STYLE,
  FADE_IN_SLIDE_INITIAL_STYLE,
  fadeOut,
  fadeSlideOut,
  useFadeIn,
  useFadeInSlide,
} from "@hooks/use-animation"

type FieldType = {
  email: string
  password: string
}

export default function Login() {
  const {t} = useTranslation()
  const navigate = useNavigate()
  const [login, loginData] = useLoginMutation()

  useTitle("general.login")

  useFadeIn("#login-wrapper")
  useFadeInSlide("#login-box", 0.3, 0.3)
  useFadeInSlide("#login-logo", 0.3, 0.5)

  const handleSubmit = useCallback(async (values: FieldType) => {
    const result = await login(values)
    if (!(result as any).error) {
      fadeSlideOut("#login-logo").finished
      await fadeSlideOut("#login-box", 0.3, 0.1).finished
      await fadeOut("#login-wrapper").finished
      navigate(General.DEFAULT_PRIVATE_ROUTE)
    }
  }, [])

  return (
    <div
      id='login-wrapper'
      style={FADE_IN_INITIAL_STYLE}
      className='w-screen h-screen bg-gradient-to-b from-sky-500 to-sky-900 flex flex-col justify-center items-center'>
      <img
        id='login-logo'
        src={logo}
        style={FADE_IN_SLIDE_INITIAL_STYLE}
        sizes='small'
        className='mb-4 w-full max-w-[224px] mb-8'
      />
      <Card
        id='login-box'
        style={FADE_IN_SLIDE_INITIAL_STYLE}
        className='p-4 w-full max-w-md mx-6'>
        <Space direction='vertical' className='w-full'>
          {loginData.error && (
            <Alert type='error' message={t("login.unauthorized")} />
          )}
          <Form layout='vertical' onFinish={handleSubmit}>
            <Form.Item<FieldType>
              label={t("general.email")}
              name='email'
              rules={[{required: true}]}>
              <Input placeholder={t("login.emailPlaceholder")} size='large' />
            </Form.Item>
            <Form.Item<FieldType>
              label={t("general.password")}
              name='password'
              rules={[{required: true}]}>
              <Input
                type='password'
                placeholder={t("login.passwordPlaceholder")}
                size='large'
              />
            </Form.Item>
            <Form.Item className='w-full'>
              <Button
                htmlType='submit'
                className='w-full mt-6'
                size='large'
                loading={loginData.isLoading}>
                {t("general.login")}
              </Button>
            </Form.Item>
          </Form>
        </Space>
      </Card>
    </div>
  )
}
