import { test, expect } from '@playwright/test';
const url = "http://localhost:3000/alert";

test.beforeEach(async ({ page }) => {
    await page.goto(url);
});  

test.describe('Create Alert', () => {
    test('should verify if latitud input is visible', async ({ page }) => {
        const latitudInput = await page.waitForSelector('input[name="Latitude"]');
        expect(latitudInput).not.toBeNull();
    });
    test('should verify if longitud input is visible', async ({ page }) => {
        const longitudInput = await page.waitForSelector('input[name="Longitude"]');
        expect(longitudInput).not.toBeNull();
    });
    test('should verify if create alert button is visible', async ({ page }) => {
        const createAlertButton = await page.waitForSelector('button[id="submit"]');
        expect(createAlertButton).not.toBeNull();
    });
    // test('should verify if create alert button is disabled', async ({ page }) => {
    //     const createAlertButton = await page.waitForSelector('button[id="submit"]');
    //     expect(await createAlertButton.isEnabled()).toBeFalsy();
    // });
    // test('should verify if create alert button is enabled', async ({ page }) => {
    //     await page.fill('input[name="Latitude"]', '1');
    //     await page.fill('input[name="Longitude"]', '1');
    //     const createAlertButton = await page.waitForSelector('button[id="submit"]');
    //     expect(await createAlertButton.isEnabled()).toBeTruthy();
    // });
    // test('should verify if create alert button is disabled after clearing latitud input', async ({ page }) => {
    //     await page.fill('input[name="Latitude"]', '1');
    //     await page.fill('input[name="Longitude"]', '1');
    //     await page.fill('input[name="Latitude"]', '');
    //     const createAlertButton = await page.waitForSelector('button[id="submit"]');
    //     expect(await createAlertButton.isEnabled()).toBeFalsy();
    // });
    // test('should verify if create alert button is disabled after clearing longitud input', async ({ page }) => {
    //     await page.fill('input[name="Latitude"]', '1');
    //     await page.fill('input[name="Longitude"]', '1');
    //     await page.fill('input[name="Longitude"]', '');
    //     const createAlertButton = await page.waitForSelector('button[id="submit"]');
    //     expect(await createAlertButton.isEnabled()).toBeFalsy();
    // });
    test('should verify if alert is created', async ({ page }) => {
        await page.fill('input[name="Latitude"]', '1');
        await page.fill('input[name="Longitude"]', '1');
        await page.click('button[id="submit"]');
        const alert = await page.waitForSelector('div[class="alert alert-success"]');
        expect(alert).not.toBeNull();
    });
});