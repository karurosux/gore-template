import {useEffect} from "react"
import {useTranslation} from "react-i18next"

export function useTitle(title: string, params?: Record<string, any>) {
  const {t} = useTranslation()
  useEffect(() => {
    document.title = t(title, params) + " | " + t("appName")
  }, [title, params])
}
