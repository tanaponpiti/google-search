# How does it work
![components.png](asset%2Fcomponents.png)

There are six components in the whole system in order to function properly.
- **nginx-proxy**: Acts as a reverse proxy, routing requests to the web and API services. It listens on ports 80 and 443.
- **web**: The web client user interface for utilize and visualize scraping API service.
- **api**: The API service that interacts with the web client, Redis, and PostgreSQL. It connects to an external HTML retriever service to fetch and store search results.
- **redis**: A Redis server used for authentication token store.
- **postgres**: A PostgreSQL database for storing search results and application data.
- **html-retriever**: A Puppeteer service to request and render html page from Google search.

Three of which, includes **web, api, html-retriever,** were implemented inside this repository.

# The problem

Surely, Google Search doesn't want its own data to be scraped. So if there are too many request sending to Google Search, 
it will block our IP from accessing it and demand CAPTCHA to be solved first.

I have tried to use TOR network to relocate IP to another. However, Google seems to list all most all of TOR node IP to be in their blacklist and still require CAPTCHA to be solved.

So, I have to find a way to allow my HTML Retriever service to gain new "clean" IP to avoid google blockade.

# The Solution

Normally, if I have a budget, I would use IP Proxy pool service and request new "clean" IP. However, I didn't want to spend more than Google Cloud Free trial credit.
Since I have to deploy this system on Google Cloud, my solutions is to use Google Cloud Run (Serverless) for the HTML Retriever service.

Google Cloud Run is a serverless service from Google that allow docker container to be run on demand.
It will start up when there is a request for the service and also scale up into multiple instance based on numbers of usage.
I happened to found out that each instance have set of different egress IP.
Therefore, I implement HTML Retriever service to terminate itself everytime its IP have been block by google search.     
When it restarted by Google Cloud Run it will automatically have different IP to continue scraping. 

Finally, our system diagram will look like this.

![deployment.png](asset%2Fdeployment.png)

# Demo Instance
Until my Google Cloud Trial free credit ran out.
You can simply access online demo [here](https://search.perthpiti.me/).

# Deployment Guide

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker: [Get Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Configuration

I have already built all of required service as a docker image.

- **web** image: ghcr.io/tanaponpiti/search-scraper-web-client:1.0.0-amd64
- **api** image: ghcr.io/tanaponpiti/search-scraper-api:1.0.0-amd64
- **api** image: ghcr.io/tanaponpiti/html-retriever:1.0.0-amd64

There are two deployment strategies available: using an external HTML retriever service hosted on Google Cloud Run for flexible IP rotation (`docker-compose.yml`) and a standalone version for local deployment (`standalone-docker-compose.yml`).

## Environment Variables Explained
### Web Service
- `VIRTUAL_HOST`: The hostname at which the service is accessible. Default is `localhost`.
- `VIRTUAL_PORT`: The port on which the service is running. For the web service, it's `80`.
- `VIRTUAL_PATH`: The path prefix for the service. Defaults to `/`.
- `API_URL`: The URL where the API service can be accessed. Typically points to the `api` service.

### API Service
- `VIRTUAL_HOST`, `VIRTUAL_PORT`, `VIRTUAL_PATH`: Similar to the web service, these variables define how the API service is accessed through the nginx proxy.
- `DB_URI`: Connection string for the PostgreSQL database.
- `JWT_SECRET`: Secret key for JWT authentication.
- `REDIS_URI`, `REDIS_PASSWORD`, `REDIS_DB`, `REDIS_CONNECTION_POOL`: Configuration for connecting to Redis, including URI, password, database index, and connection pool size.
- `TOKEN_EXPIRE_HOUR`: Defines the lifespan of authentication tokens.
- `HTML_RETRIEVER_STANDALONE`: Indicates whether the HTML retriever service is running in standalone mode.
- `HTML_RETRIEVER_URL`: URL of the HTML retriever service.
- `CLOUD_RUN_KEY_PATH`: Path to the Google Cloud service account key file (specific to the API service).
- `CONCURRENT_SCRAPE_LIMIT`: Defines how many concurrent request to scrape at a time.
- `CLOUD_RUN_KEY_PATH`: Path to the Google Cloud service account key file (specific to the API service).

### Postgres and Redis Services
- For `postgres`: `POSTGRES_DB`, `POSTGRES_USER`, `POSTGRES_PASSWORD` define the database, user, and password.
- For `redis`: The command includes `--requirepass` to set the password.

### Differences between Compose Files
- The `HTML_RETRIEVER_STANDALONE` variable in the `api` service is `false` in the default setup, indicating the HTML retriever runs as part of the overall application, in Google Cloud Run. In the standalone setup, it is `true`, and a separate `html-retriever` service is defined.
- When `HTML_RETRIEVER_STANDALONE` is `false`, an `CLOUD_RUN_KEY_PATH` env will be required for the Google Cloud Run authentication process.

## Deployment

To deploy the application, navigate to the directory containing your desired Docker Compose file and run:

```bash
docker-compose up -d
```

Replace `docker-compose.yml` with `standalone-docker-compose.yml` for local deployment.

