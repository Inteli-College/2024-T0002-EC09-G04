import { test, expect } from '@playwright/test';
const url = "http://localhost:3000";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Home Page', () => {
    test('signup button should be visible', async ({ page }) => {
        const usernameInput = await page.waitForSelector('a[href="/login"]');
        expect(usernameInput).not.toBeNull();
    });
});