import {ConfigProvider, theme} from "antd"
import React from "react"
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom"
import "./App.css"
import {ProtectedRoute} from "@components/ProtectedRoute"
import {General} from "@constants/general"
import {PublicRoute} from "@components/PublicRoute"
import ErrorBoundary from "antd/es/alert/ErrorBoundary"
import colors from "tailwindcss/colors"
import {RouteLoadingSuspense} from "@components/RouteLoadingSuspense"

// Lazy loaded components
const Login = React.lazy(() => import("./containers/login/Login"))
const Root = React.lazy(() => import("./containers/root/Root"))
const UserManagement = React.lazy(
  () => import("./containers/root/views/user-management/UserManagement"),
)
const UserProfile = React.lazy(
  () => import("./containers/root/views/user-management/UserProfile"),
)
const RoleManagement = React.lazy(
  () => import("./containers/root/views/role-management/RoleManagement"),
)
const RoleProfile = React.lazy(
  () => import("./containers/root/views/role-management/RoleProfile"),
)
const BranchManagement = React.lazy(
  () => import("./containers/root/views/branch-management/BranchManagement"),
)
const CustomerManagement = React.lazy(
  () =>
    import("./containers/root/views/customer-management/CustomerManagement"),
)

function App() {
  return (
    <ErrorBoundary>
      <ConfigProvider
        theme={{
          token: {
            colorBgLayout: colors.neutral[900],
            colorBgMask: colors.neutral[900],
            borderRadius: 8,
            colorPrimary: colors.sky[500],
          },
          algorithm: theme.darkAlgorithm,
        }}>
        <BrowserRouter>
          <Routes>
            <Route
              key='login'
              path='/login'
              element={
                <PublicRoute>
                  <Login />
                </PublicRoute>
              }
            />
            <Route
              path='/root/*'
              element={
                <ProtectedRoute>
                  <Root />
                </ProtectedRoute>
              }>
              <Route
                key='user-management'
                path='user-management'
                element={
                  <RouteLoadingSuspense>
                    <UserManagement />
                  </RouteLoadingSuspense>
                }
              />
              <Route
                key='user-profile'
                path='user-management/:userId'
                element={
                  <RouteLoadingSuspense>
                    <UserProfile />
                  </RouteLoadingSuspense>
                }
              />
              <Route
                key='role-management'
                path='role-management'
                element={
                  <RouteLoadingSuspense>
                    <RoleManagement />
                  </RouteLoadingSuspense>
                }
              />
              <Route
                key='role-profile'
                path='role-management/:roleId'
                element={
                  <RouteLoadingSuspense>
                    <RoleProfile />
                  </RouteLoadingSuspense>
                }
              />
              <Route
                key='branch-management'
                path='branch-management'
                element={
                  <RouteLoadingSuspense>
                    <BranchManagement />
                  </RouteLoadingSuspense>
                }
              />
              <Route
                key='customer-management'
                path='customer-management'
                element={
                  <RouteLoadingSuspense>
                    <CustomerManagement />
                  </RouteLoadingSuspense>
                }
              />
              <Route path='*' element={<Navigate to='user-management' />} />
            </Route>
            <Route
              path='*'
              element={<Navigate to={General.DEFAULT_PUBLIC_ROUTE} />}
            />
          </Routes>
        </BrowserRouter>
      </ConfigProvider>
    </ErrorBoundary>
  )
}

export default App
