import {PropsWithChildren, Suspense} from "react"

export function RouteLoadingSuspense({children}: PropsWithChildren<{}>) {
  return <Suspense fallback={<span />}>{children}</Suspense>
}
