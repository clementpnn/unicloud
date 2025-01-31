import { test, expect } from "@playwright/test"

test.describe("404 Not Found page tests", () => {
  test("should display 404 Not Found", async ({ page }) => {
    await page.goto("/non-existent-route")

    await expect(page.locator("text=404")).toBeVisible()
    await expect(page.locator("text=Page Not Found")).toBeVisible()

    await expect(page.locator("text=Go to home")).toBeVisible()
  })

  test("redirect to home page", async ({ page }) => {
    await page.goto("/non-existent-route")

    await page.click("text=Go to home")

    await expect(page).toHaveURL("/")
  })
})
