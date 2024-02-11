const express = require('express');
const puppeteer = require('puppeteer');
const bodyParser = require('body-parser');

puppeteer.launch({headless: true, args: ['--no-sandbox','--disable-features=site-per-process']}).then(browser => {
    const app = express();
    const port = process.env.PORT || 8081;
    app.use(bodyParser.json());

    async function requestHtml(url) {
        const page = await browser.newPage();

        // Listen for all responses. If any response has a 429 status code, terminate the program.
        page.on('response', async response => {
            if (response.status() === 429) {
                console.error('Error 429 encountered. Shutting down the server to reallocate new IP...');
                await browser.close();
                process.exit(1); // Terminate the program
            }
        });

        try {
            await page.goto(url, {waitUntil: 'networkidle2'});
            await page.waitForSelector('div#result-stats', {visible: true, timeout: 5000}).catch(error => {
                console.log('The element div#result-stats did not appear within the timeout period.', error);
            });
            return await page.content();
        } catch (error) {
            throw error;
        }
    }

    app.post('/request-html', async (req, res) => {
        const {url} = req.body;
        if (!url) {
            return res.status(400).json({error: 'URL is required'});
        }
        try {
            const htmlContent = await requestHtml(url);
            res.status(200).send(htmlContent);
        } catch (error) {
            res.status(500).json({error: error.message});
        }
    });

    app.listen(port, async () => {
        console.log(`Server running on port ${port}`);
    });
});
