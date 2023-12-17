import {PayloadAction, createSlice} from "@reduxjs/toolkit"
import {BreadcrumbItem} from "../model/breadcrumb-item"

export const breadcrumbSlice = createSlice({
  name: "breadcrumbs",
  initialState: [] as BreadcrumbItem[],
  reducers: {
    setBreadcrumbs: (_, action: PayloadAction<BreadcrumbItem[]>) =>
      action.payload,
  },
})

export const {setBreadcrumbs} = breadcrumbSlice.actions
