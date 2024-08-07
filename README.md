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
EMAIL=<Twitter account email>
USERNAME=<Twitter username>
PASSWORD=<Twitter password>
AHBCC_DOMAIN=<Domain of the application with the endpoint /tweets/v1> --> In this case the app AHBCC
```
[^1]

[^1]: AHBCC: Adverse Human Behaviour Corpus Creator. More information [here](https://github.com/lhbelfanti/ahbcc)

### Run

In the root folder, run:

```
go run cmd/api/main.go --local
```

## Setting up & run (into a Docker container)

### Setup

1. Create a `.env` file at the root of the project, and add the following environment variables:

```
EMAIL=<Twitter account email>
USERNAME=<Twitter username>
PASSWORD=<Twitter password>
DRIVER_PATH=<The path to the Chrome driver> --> Example: /usr/bin/chromedriver
BROWSER_PATH=<The path to the Chrome browser> --> Example: /usr/bin/chromium
RABBITMQ_USER=<The RabbitMQ user>
RABBITMQ_PASS=<The RabbitMQ password>
AHBCC_DOMAIN=<Domain of the application with the endpoint /tweets/v1> --> In this case the app AHBCC
```
[^1]

### Build & Run

```
docker compose up --build
```

---

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Logo License

[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](https://creativecommons.org/licenses/by-nc-sa/4.0/)

The logo was obtained from https://github.com/ashleymcnamara/gophers, but it was slightly modified to be representative for this repository.

