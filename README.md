<p align="center">
  <img src="media/goxcrap-logo.png" width="100" alt="Repository logo" />
</p>
<h3 align="center">GoXCrap</h3>
<p align="center">X (formerly Twitter) web scrapper, written in Go<p>
<p align="center">
    <img src="https://img.shields.io/github/repo-size/lhbelfanti/goxcrap?label=Repo%20size" alt="Repo size" />
    <img src="https://img.shields.io/github/license/lhbelfanti/goxcrap?label=License" alt="License" />
    <img src="https://codecov.io/gh/lhbelfanti/goxcrap/graph/badge.svg?token=69LLNMKXRU" alt="Coverage" />
</p>

---

# GoXCrap

This application collects tweets based on a defined search criteria, and save them in a database.

## Set up & run (locally)

### Set up

1. First of all you need to download the Chrome Web Driver that matches with the installed version of Google Chrome (the
   browser used for testing this project). </br>
   You can download it from [here](https://googlechromelabs.github.io/chrome-for-testing/), or you can
   use `@puppeteer/browsers` with [this](https://pptr.dev/browsers-api) installation guide. </br>
   After that, copy it inside the [internal/webdriver](./internal/webdriver) folder.
2. Create a `.env` file at the root of the project, and add the following environment variables:

```
EMAIL=<your-twitter-account-email>
USERNAME=<your-twitter-username>
PASSWORD=<your-twitter-password>
```

### Run

In the root folder, run:

```
go run cmd/api/main.go --local
```

## Setting up & run (into a Docker container)

### Setup

1. Create a `.env` file at the root of the project, and add the following environment variables:

```
EMAIL=<your-twitter-account-email>
USERNAME=<your-twitter-username>
PASSWORD=<your-twitter-password>
DRIVER_PATH=/usr/bin/chromedriver
BROWSER_PATH=/usr/bin/chromium
```

### Build & Run

```
docker-compose up --build
```

---

## License

[MIT](https://choosealicense.com/licenses/mit/)
