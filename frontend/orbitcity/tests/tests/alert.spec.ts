import { test, expect } from '@playwright/test';
const url = "https://www.latlong.net/";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Create Alert', () => {
    test('should verify if latitud input is visible', async ({ page }) => {
        const latitudInput = await page.waitForSelector('input[name="lat"]');
        expect(latitudInput).not.toBeNull();
    });
    test('should verify if longitud input is visible', async ({ page }) => {
        const longitudInput = await page.waitForSelector('input[name="lng"]');
        expect(longitudInput).not.toBeNull();
    });
    test('should verify if create alert button is visible', async ({ page }) => {
        const createAlertButton = await page.waitForSelector('button[type="submit"]');
        expect(createAlertButton).not.toBeNull();
    });
    test('should verify if create alert button is disabled', async ({ page }) => {
        const createAlertButton = await page.waitForSelector('button[type="submit"]');
        expect(await createAlertButton.isEnabled()).toBeFalsy();
    });
    test('should verify if create alert button is enabled', async ({ page }) => {
        await page.fill('input[name="lat"]', '1');
        await page.fill('input[name="lng"]', '1');
        const createAlertButton = await page.waitForSelector('button[type="submit"]');
        expect(await createAlertButton.isEnabled()).toBeTruthy();
    });
    test('should verify if create alert button is disabled after clearing latitud input', async ({ page }) => {
        await page.fill('input[name="lat"]', '1');
        await page.fill('input[name="lng"]', '1');
        await page.fill('input[name="lat"]', '');
        const createAlertButton = await page.waitForSelector('button[type="submit"]');
        expect(await createAlertButton.isEnabled()).toBeFalsy();
    });
    test('should verify if create alert button is disabled after clearing longitud input', async ({ page }) => {
        await page.fill('input[name="lat"]', '1');
        await page.fill('input[name="lng"]', '1');
        await page.fill('input[name="lng"]', '');
        const createAlertButton = await page.waitForSelector('button[type="submit"]');
        expect(await createAlertButton.isEnabled()).toBeFalsy();
    });
    test('should verify if alert is created', async ({ page }) => {
        await page.fill('input[name="lat"]', '1');
        await page.fill('input[name="lng"]', '1');
        await page.click('button[type="submit"]');
        const alert = await page.waitForSelector('div[class="alert alert-success"]');
        expect(alert).not.toBeNull();
    });
});