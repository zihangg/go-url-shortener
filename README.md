# Introduction

This is a simple implementation for a url shortener that was done for fun.

## How it works

The application exposes two APIs: `/api/v1/shorten` and `/`. The shorten endpoints creates an id (stored with a 5 minute expiry in Redis), and the base endpoint redirects the user to the original embedded endpoint.

```bash
curl --location 'localhost:8080/api/v1/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.google.com"
}'
```

## Running the application

`docker compose up --build`
