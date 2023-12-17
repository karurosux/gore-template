import {PoweroffOutlined} from "@ant-design/icons"
import {
  authenticatedUserApi,
  useLogoutMutation,
} from "@api/authenticated-user.api"
import logo from "@assets/react.svg"
import {Button, Layout, Space, Tooltip} from "antd"
import {useCallback} from "react"
import {useTranslation} from "react-i18next"
import {Link, Outlet, useNavigate} from "react-router-dom"
import {Breadcrumbs} from "./containers/breadcrumbs/Breadcrumbs"
import {MainNavMenu} from "./containers/main-nav-menu/MainNavMenu"
import {useAppDispatch} from "@src/store"
import {usersApi} from "@src/api/users.api"
import {rolesApi} from "@src/api/roles.api"
import {permissionsApi} from "@src/api/permissions.api"
import {
  FADE_IN_INITIAL_STYLE,
  FADE_IN_SLIDE_INITIAL_STYLE,
  fadeOut,
  fadeSlideOut,
  useFadeIn,
  useFadeInSlide,
} from "@src/hooks/use-animation"

export default function Root() {
  const {t} = useTranslation()
  const [logout, logoutMutation] = useLogoutMutation()
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  useFadeInSlide("#app-header")
  useFadeIn("#app-nav", 0.3, 0.1)
  useFadeIn("#app-breadcrumbs", 0.3, 0.5)
  useFadeInSlide("#app-content", 0.3, 0.3)

  const handleLogout = useCallback(async () => {
    const response: any = await logout()
    if (!response.error) {
      fadeSlideOut("#app-header")
      fadeOut("#app-nav")
      fadeOut("#app-breadcrumbs")
      await fadeSlideOut("#app-content", 0.3, 0.2).finished
      dispatch(authenticatedUserApi.util.resetApiState())
      dispatch(usersApi.util.resetApiState())
      dispatch(rolesApi.util.resetApiState())
      dispatch(permissionsApi.util.resetApiState())
      navigate("/")
    }
  }, [])

  return (
    <Layout className='w-screen h-screen'>
      <Layout.Header
        id='app-header'
        style={FADE_IN_SLIDE_INITIAL_STYLE}
        className='!bg-transparent border-b border-b-neutral-700 flex justify-center items-center'>
        <Link to='/root'>
          <img src={logo} alt={t("appName")} className='w-8' />
        </Link>
        <span className='flex-1' />
        <Space>
          <Tooltip title={t("general.logout")}>
            <Button
              type='text'
              loading={logoutMutation.isLoading}
              onClick={handleLogout}
              icon={<PoweroffOutlined />}
            />
          </Tooltip>
        </Space>
      </Layout.Header>
      <MainNavMenu id='app-nav' style={FADE_IN_INITIAL_STYLE} />
      <Breadcrumbs id='app-breadcrumbs' style={FADE_IN_INITIAL_STYLE} />
      <Layout id='app-content' style={FADE_IN_SLIDE_INITIAL_STYLE}>
        <Layout.Content className='!bg-transparent py-4 px-6'>
          <div className='container m-auto'>
            <Outlet />
          </div>
        </Layout.Content>
      </Layout>
    </Layout>
  )
}
