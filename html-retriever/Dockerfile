FROM --platform=linux/amd64 node:18 AS dependencies
WORKDIR /usr/src/app
COPY .puppeteerrc.cjs ./
COPY package*.json ./
RUN apt-get update && apt-get install gnupg wget -y && \
    wget --quiet --output-document=- https://dl-ssl.google.com/linux/linux_signing_key.pub | gpg --dearmor > /etc/apt/trusted.gpg.d/google-archive.gpg && \
    sh -c 'echo "deb [arch=amd64] https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list' && \
    apt-get update && \
    apt-get install google-chrome-stable -y --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

RUN npm install
COPY main.js ./
EXPOSE 8081
CMD [ "npm", "run", "start" ]
