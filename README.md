# Larafiber

Welcome to Larafiber, a project using Golang (Gofiber), using some of the same structures as Laravel. This is not Laravel written in Go, but my take on creating a base for all my projects, inspired by Laravel. This project is using Gofiber 2, GORM 2, Vue 3, Tailwind 3, and Inertia.js.

The idea behind the project is to have a good starting point for web app development and prototyping. This setup has more than once being used as a base for rapidly developing web applications.

This project is using SQLight, but GORM can be configured to use a lot of different database engines. 


## Setup

Init and run the frontend by running: 

```npm install && npm run watch```

Init and run the backend:

```go get```

```go run main.go```

## Create an .env file

```
template_path=./views/
db_path=./database/db.sqlite
sparkpost=
domain=http://localhost:4000/
app_env=local
port=4000
jwt_token=
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_BUCKET=
```