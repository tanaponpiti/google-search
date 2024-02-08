const express = require('express');
const puppeteer = require('puppeteer');
const bodyParser = require('body-parser');
//TODO fix this --no-sandbox security hole
puppeteer.launch({headless: true, args: ['--no-sandbox','--disable-features=site-per-process']}).then(browser => {
    const app = express();
    const port = process.env.PORT || 8081;
    app.use(bodyParser.json());

    async function requestHtml(url) {
        const page = await browser.newPage();
        try {
            await page.goto(url, {waitUntil: 'networkidle2'});
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
            if (error.message.includes('429')) {
                await browser.close()
                console.error('Error 429 encountered. Shutting down the server...to reallocate new IP');
                process.exit(1);
            }
            res.status(500).json({error: error.message});
        }
    });

    app.listen(port, async () => {
        console.log(`Server running on port ${port}`);
    });
});
