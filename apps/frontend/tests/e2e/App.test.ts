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

  test("should show an error if URL is invalid in form input", async ({ page }) => {
    await page.goto("/")
    await page.fill("input[placeholder='https://example.com']", "invalidurl")
    await page.click("text=Send")
    await expect(page.locator("text=Please enter a valid URL")).toBeVisible()
  })

  test("should handle URL submission in form input and copy to clipboard", async ({ page }) => {
    await page.goto("/")
    await page.fill("input[placeholder='https://example.com']", "https://www.google.fr/")
    await page.click("text=Send")

    await expect(page.locator("text=Your link is ready !!!")).toBeVisible()
    const linkText = await expect(page.locator("#link")).not.toBeEmpty()
    await expect(page.locator("#copy")).toBeVisible()

    await page.evaluate(() => {
      navigator.clipboard.writeText = (data) => {
        window["copiedText"] = data
        return Promise.resolve()
      }
    })

    await page.click("#copy")
    const copiedText = await page.evaluate(() => window["copiedText"])

    expect(copiedText).toBe(linkText)
  })
})