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

## Creating new models

All the models are located under the app-folder. There is a base file that contains a BASE struct. This is the base for all models, since it holds both ID and UUID. UUID is used in frontend, so we don´t have to expose the ID.

To add the model to GORM, simply add it to the auto migrate definition in the init function, and GORM will greate the database schema:

```
db.Debug().AutoMigrate(
    &ForgottenPassword{},
    &User{},
    &Account{},
    &Version{},
    &Task{},
)
```

## How it all comes together

When a request is sent to the server, the following happens:
* The router file is invoced (routes/web.go), where the matching path and method are located.
* If there are any middelware hoocked up, this will be invoced from routes/web.go
* When a route is matched, and controller function is called.
* The inertia.Render function is called, with the data and the name of the vue file.
* Inertia.js on the frontend is called with the data and the vue file (plus version and path, but that is described on Inertia.js´ own web page), and rendered based on the vue files. 