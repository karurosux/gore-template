import {BreadcrumbItem} from "@src/model/breadcrumb-item"
import {useAppSelector} from "@src/store"
import {Breadcrumb} from "antd"
import {CSSProperties, MouseEventHandler} from "react"
import {useTranslation} from "react-i18next"
import {useNavigate} from "react-router-dom"

type Props = {
  id?: string
  style?: CSSProperties
}

export function Breadcrumbs(props: Props) {
  const {t} = useTranslation()
  const breadcrumbs = useAppSelector((state) => state.breadcrumbs)
  const navigate = useNavigate()

  const handleClick =
    (
      item: BreadcrumbItem,
    ): MouseEventHandler<HTMLAnchorElement | HTMLSpanElement> =>
    (e) => {
      if (item.useNavitagion) {
        e.preventDefault()
        navigate(item.href as any)
      }
    }

  return (
    <div id={props.id} style={props.style} className='p-4 px-12'>
      <Breadcrumb>
        {breadcrumbs.map((breadcrumb) => (
          <Breadcrumb.Item
            key={breadcrumb.title}
            href={breadcrumb.href as string}
            onClick={handleClick(breadcrumb)}>
            {t(breadcrumb.title)}
          </Breadcrumb.Item>
        ))}
      </Breadcrumb>
    </div>
  )
}
