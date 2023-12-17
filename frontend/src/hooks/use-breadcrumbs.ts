import {BreadcrumbItem} from "@src/model/breadcrumb-item"
import {setBreadcrumbs} from "@src/slices/breadcrumbs.slice"
import {useAppDispatch} from "@src/store"
import {useEffect} from "react"

export function useBreadcrumbs(items: BreadcrumbItem[]) {
  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(setBreadcrumbs(items))
  }, [items])
}
