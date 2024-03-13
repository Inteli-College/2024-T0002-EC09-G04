import { test, expect } from '@playwright/test';
const url = "http://localhost:3000/signup";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Signup', () => {
    test('should verify if signup button is visible', async ({ page }) => {
        const signupButton = await page.waitForSelector('a[href="/accounts/emailsignup/"]');
        expect(signupButton).not.toBeNull();
    });
    test('should verify if signup button redirects to signup page', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        expect(page.url()).toBe("https://www.instagram.com/accounts/emailsignup/");
    });
    test('should verify if input username is visible', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        const usernameInput = await page.waitForSelector('input[name="username"]');
        expect(usernameInput).not.toBeNull();
    });
    test('should verify if input password is visible', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        const passwordInput = await page.waitForSelector('input[name="password"]');
        expect(passwordInput).not.toBeNull();
    });
    test('should verify if password confirmation is visible', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        const passwordConfirmation = await page.waitForSelector('input[name="passwordConfirm"]');
        expect(passwordConfirmation).not.toBeNull();
    });
    test('should verify if signup button is disabled', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        const signupButton = await page.waitForSelector('button[type="submit"]');
        expect(await signupButton.isEnabled()).toBeFalsy();
    });
    test('should verify if signup button is enabled', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="passwordConfirm"]', 'testing123');
        const signupButton = await page.waitForSelector('button[type="submit"]');
        expect(await signupButton.isEnabled()).toBeTruthy();
    });
    test('should verify if signup button is disabled after clearing username input', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="passwordConfirm"]', 'testing123');
        await page.fill('input[name="username"]', '');
        const signupButton = await page.waitForSelector('button[type="submit"]');
        expect(await signupButton.isEnabled()).toBeFalsy();
    });
    test('should verify if signup button is disabled after clearing password input', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="passwordConfirm"]', 'testing123');
        await page.fill('input[name="password"]', '');
        const signupButton = await page.waitForSelector('button[type="submit"]');
        expect(await signupButton.isEnabled()).toBeFalsy();
    });
    test('should verify if url changes after clicking signup button', async ({ page }) => {
        await page.click('a[href="/accounts/emailsignup/"]');
        await page.fill('input[name="username"]', 'testing123');
        await page.fill('input[name="password"]', 'testing123');
        await page.fill('input[name="passwordConfirm"]', 'testing123');
        await page.click('button[type="submit"]');
        expect(page.url()).not.toBe("https://www.instagram.com/accounts/emailsignup/");
    });
});
