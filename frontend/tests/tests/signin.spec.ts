import { test, expect } from '@playwright/test';
const url = "http://localhost:3000/login";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Signin', () => {
    test('input username should be visible', async ({ page }) => {
        const usernameInput = await page.waitForSelector('input[id="Email"]');
        expect(usernameInput).not.toBeNull();
    });
    test('input password should be visible', async ({ page }) => {
        const passwordInput = await page.waitForSelector('input[id="password"]');
        expect(passwordInput).not.toBeNull();
    });
    // test('checkbox should be visible', async ({ page }) => {
    //     const checkbox = await page.waitForSelector('input[type="checkbox"]');
    //     expect(checkbox).not.toBeNull();
    // });
    test('login button should be visible', async ({ page }) => {
        const loginButton = await page.waitForSelector('button[id="submit"]');
        expect(loginButton).not.toBeNull();
    });
    test('password should be hidden', async ({ page }) => {
        const passwordInput = await page.waitForSelector('input[id="password"]');
        expect(await passwordInput.getAttribute('type')).toBe('password');
    });
    // test('login button should be disabled', async ({ page }) => {
    //     const loginButton = await page.waitForSelector('button[type="submit"]');
    //     expect(await loginButton.isEnabled()).toBeFalsy();
    // });
    // test('login button should be enabled', async ({ page }) => {
    //     await page.fill('input[id="Email"]', 'testing123');
    //     await page.fill('input[id="password"]', 'testing123');
    //     const loginButton = await page.waitForSelector('button[type="submit"]');
    //     expect(await loginButton.isEnabled()).toBeTruthy();
    // });
    // test('login button should be disabled after clearing username input', async ({ page }) => {
    //     await page.fill('input[id="Email"]', 'testing123');
    //     await page.fill('input[id="password"]', 'testing123');
    //     await page.fill('input[id="Email"]', '');
    //     const loginButton = await page.waitForSelector('button[type="submit"]');
    //     expect(await loginButton.isEnabled()).toBeFalsy();
    // });
    test('message should appear after clearing username input', async ({ page }) => {
        await page.fill('input[id="Email"]', 'testing123');
        await page.fill('input[id="password"]', 'testing123');
        await page.fill('input[id="Email"]', '');
        expect(await expect(page.locator('div[class="ant-form-item-explain-error"]')).toHaveText('Por favor insira seu email!'));
    });
    // test('login button should be disabled after clearing password input', async ({ page }) => {
    //     await page.fill('input[id="Email"]', 'testing123');
    //     await page.fill('input[id="password"]', 'testing123');
    //     await page.fill('input[id="password"]', '');
    //     const loginButton = await page.waitForSelector('button[type="submit"]');
    //     expect(await loginButton.isEnabled()).toBeFalsy();
    // });
    test('message should appear after clearing password input', async ({ page }) => {
        await page.fill('input[id="Email"]', 'testing123');
        await page.fill('input[id="password"]', 'testing123');
        await page.fill('input[id="password"]', '');
        expect(await expect(page.locator('div[class="ant-form-item-explain-error"]')).toHaveText('Por favor digite sua senha!'));
    });
    // test('url should change after clicking login button', async ({ page }) => {
    //     await page.fill('input[id="Email"]', 'testing123');
    //     await page.fill('input[id="password"]', 'testing123');
    //     await page.click('button[type="submit"]');
    //     expect(page.url()).not.toBe(url);
    // });
});