import i18n from "i18next"
import {initReactI18next} from "react-i18next"
import enJSON from "./i18n/en.json"

i18n.use(initReactI18next).init({
  resources: {
    en: {
      translation: enJSON,
    },
  },
  lng: "en",
})
