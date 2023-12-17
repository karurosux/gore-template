import {useTranslation} from "react-i18next"

export default function CustomerManagement() {
  const {t} = useTranslation()
  return <div>{t("customerManagmenet.title")}</div>
}
