import { test, expect } from '@playwright/test';

test('should not acess alert page without login', async ({ page }) => {
    await page.goto('http://localhost:3000/alert');
    expect(page.url()).toBe('http://localhost:3000/');
});

test('should not acess sensor page without login', async ({ page }) => {
    await page.goto('http://localhost:3000/sensor');
    expect(page.url()).toBe('http://localhost:3000/');
});

test('should access alert page after login', async ({ page }) => {
    await page.goto('http://localhost:3000/login');
    await page.fill('input[id="Email"]', 'testing123');
    await page.fill('input[id="password"]', 'testing123');
    await page.click('button[type="submit"]');
    await page.goto('http://localhost:3000/alert');
    expect(page.url()).toBe('http://localhost:3000/alert');
});

test('should access sensor page after login', async ({ page }) => {
    await page.goto('http://localhost:3000/login');
    await page.fill('input[id="Email"]', 'testing123');
    await page.fill('input[id="password"]', 'testing123');
    await page.click('button[type="submit"]');
    await page.goto('http://localhost:3000/sensor');
    expect(page.url()).toBe('http://localhost:3000/sensor');
});