# mobydev-internship-admin

This is a small application to demonstrate admin access to change data. It's built using [Go](https://golang.org/) and uses a Makefile for easy running.

## Features

- Admin registration
- Admin authentication
- Form data updating

## Prerequisites

Before running this application, make sure you have [Go](https://golang.org/dl/) installed on your machine.

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/qPyth/mobydev-internship-auth
```


## Configuration

This application uses environment variables for configuration. Copy the `.env.example` file to `.env` and adjust the settings according to your environment:

```bash
cp .env.example .env
```


Edit the `.env` file and set your secret key for JWTAuth, and edit the `config.yaml` and set your credentials and other necessary configurations.


## Running the Application

To run the application, use the `make run` command from the root directory of the project:

```bash
make run
```
or run with docker
```bash
make docker-build && make docker-run
```


This command compiles the Go application and starts the server on the default port.

## API Endpoints

The application exposes the following endpoints:

- `POST /admin/register`: Register a new admin user. Requires a JSON body with `email`, `password` and `pass_conf` fields.
- `POST /admin/login`: Authenticate a admin. Requires a JSON body with `email` and `password`. Returns a JWT token upon successful authentication.
- `POST /project/update`: Update the project data. Requires a JSON body with from a task form.
