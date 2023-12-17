import {useTranslation} from "react-i18next"

export default function BranchManagement() {
  const {t} = useTranslation()
  return <div>{t("branchManagement.title")}</div>
}
