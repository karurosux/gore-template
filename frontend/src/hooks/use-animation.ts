import {useEffect} from "react"
import {animate} from "motion"

export function useFadeInSlide(
  selector: string,
  duration: number = 0.3,
  delay: number = 0,
) {
  useEffect(() => {
    animate(
      selector,
      {
        transform: "translateY(0)",
        opacity: 1,
      },
      {
        duration,
        delay,
      },
    )
  }, [])
}

export function useFadeIn(
  selector: string,
  duration: number = 0.3,
  delay: number = 0,
) {
  useEffect(() => {
    animate(
      selector,
      {
        opacity: 1,
      },
      {
        duration,
        delay,
      },
    )
  }, [])
}

export function fadeOut(
  selector: string,
  duration: number = 0.3,
  delay: number = 0,
) {
  return animate(selector, FADE_IN_INITIAL_STYLE, {
    duration,
    delay,
  })
}

export function fadeSlideOut(
  selector: string,
  duration: number = 0.3,
  delay: number = 0,
) {
  return animate(selector, FADE_IN_SLIDE_INITIAL_STYLE, {
    duration,
    delay,
  })
}

export const FADE_IN_SLIDE_INITIAL_STYLE = {
  opacity: 0,
  transform: "translateY(-100px)",
}

export const FADE_IN_INITIAL_STYLE = {
  opacity: 0,
}
