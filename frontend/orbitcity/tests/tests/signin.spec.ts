import { test, expect } from '@playwright/test';
const url = "https://instagram.com";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Signin', () => {
    test('should verify if input username is visible', async ({ page }) => {
        const usernameInput = await page.waitForSelector('input[name="username"]');
        expect(usernameInput).not.toBeNull();
    });
    test('should verify if input password is visible', async ({ page }) => {
        const passwordInput = await page.waitForSelector('input[name="password"]');
        expect(passwordInput).not.toBeNull();
    });
    test('should verify if login button is visible', async ({ page }) => {
        const loginButton = await page.waitForSelector('button[type="submit"]');
        expect(loginButton).not.toBeNull();
    });
    test('should verify if login button is disabled', async ({ page }) => {
        const loginButton = await page.waitForSelector('button[type="submit"]');
        expect(await loginButton.isEnabled()).toBeFalsy();
    });
    test('should verify if login button is enabled', async ({ page }) => {
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        const loginButton = await page.waitForSelector('button[type="submit"]');
        expect(await loginButton.isEnabled()).toBeTruthy();
    });
    test('should verify if login button is disabled after clearing username input', async ({ page }) => {
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="username"]', '');
        const loginButton = await page.waitForSelector('button[type="submit"]');
        expect(await loginButton.isEnabled()).toBeFalsy();
    });
    test('should verify if login button is disabled after clearing password input', async ({ page }) => {
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="password"]', '');
        const loginButton = await page.waitForSelector('button[type="submit"]');
        expect(await loginButton.isEnabled()).toBeFalsy();
    });
    test('should verify if url changes after clicking login button', async ({ page }) => {
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.click('button[type="submit"]');
        expect(page.url()).not.toBe(url);
    });
});