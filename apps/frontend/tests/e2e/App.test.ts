import { test, expect } from "@playwright/test"

test.describe("Main page tests", () => {
  test("should load the main page with all elements", async ({ page }) => {
    await page.goto("/")
    await expect(page.locator("text=ClipLink")).toBeVisible()
    await expect(page.locator("input[placeholder='https://example.com']")).toBeVisible()
    await expect(page.locator("text=Send")).toBeVisible()
  })

  test("should show an error if URL is empty in form input", async ({ page }) => {
    await page.goto("/")
    await page.click("text=Send")
    await expect(page.locator("text=URL cannot be empty")).toBeVisible()
  })
})