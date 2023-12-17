import {GroupOutlined, ShopOutlined, UserOutlined} from "@ant-design/icons"
import {usePermissionsMap} from "@hooks/use-permisisons-map"
import {Menu, Skeleton} from "antd"
import {ItemType, MenuItemType} from "antd/es/menu/hooks/useItems"
import {CSSProperties, useEffect, useState} from "react"
import {useTranslation} from "react-i18next"
import {useLocation, useNavigate} from "react-router-dom"

type Props = {
  id?: string
  style?: CSSProperties
}

export function MainNavMenu(props: Props) {
  const {t} = useTranslation()
  const location = useLocation()
  const navigate = useNavigate()
  const [menu, setMenu] = useState<ItemType<MenuItemType>[]>([])
  const {permissionsMap, loaded} = usePermissionsMap()

  useEffect(() => {
    if (!loaded) {
      return
    }

    const newMenu: ItemType<MenuItemType>[] = []

    if (
      permissionsMap?.USER_MANAGEMENT?.write ||
      permissionsMap?.USER_MANAGEMENT?.read
    ) {
      newMenu.push({
        key: "user-management",
        label: t("userManagement.title"),
        icon: <UserOutlined />,
        onClick: handleOptionClicked("user-management"),
      })
    }

    if (
      permissionsMap?.CUSTOMER_MANAGEMENT?.write ||
      permissionsMap?.CUSTOMER_MANAGEMENT?.read
    ) {
      newMenu.push({
        key: "customer-management",
        label: t("customerManagement.title"),
        icon: <ShopOutlined />,
        onClick: handleOptionClicked("customer-management"),
      })
    }

    if (
      permissionsMap?.ROLE_MANAGEMENT?.write ||
      permissionsMap?.ROLE_MANAGEMENT?.read
    ) {
      newMenu.push({
        key: "role-management",
        label: t("roleManagement.title"),
        icon: <GroupOutlined />,
        onClick: handleOptionClicked("role-management"),
      })
    }

    // TODO: uncomment when branch management is ready
    // wont be developed for now.
    // if (
    //   permissionsMap?.BRANCH_MANAGEMENT?.write ||
    //   permissionsMap?.BRANCH_MANAGEMENT?.read
    // ) {
    //   newMenu.push({
    //     key: "branch-management",
    //     label: t("branchManagement.title"),
    //     icon: <BranchesOutlined />,
    //     onClick: handleOptionClicked("branch-management"),
    //   });
    // }

    setMenu(newMenu)
  }, [permissionsMap, loaded])

  const handleOptionClicked = (link: string) => () => {
    navigate(link)
  }

  return !loaded && !menu.length ? (
    <div
      id={props.id}
      style={props.style}
      className='flex pt-2 px-9 border-b min-h-[48px] border-b-neutral-700'>
      <Skeleton
        paragraph={false}
        active
        title={{className: "!h-6"}}
        rootClassName='p-2 w-56'
      />
      <Skeleton
        paragraph={false}
        active
        title={{className: "!h-6"}}
        rootClassName='p-2 w-56'
      />
      <Skeleton
        paragraph={false}
        active
        title={{className: "!h-6"}}
        rootClassName='p-2 w-56'
      />
    </div>
  ) : (
    <div
      id={props.id}
      style={props.style}
      className='w-full border-b min-h-[48px] border-b-neutral-700 px-8'>
      <Menu
        className='w-full bg-transparent !border-none'
        items={menu}
        mode='horizontal'
        selectedKeys={[location.pathname.split("/")[2]]}
      />
    </div>
  )
}
