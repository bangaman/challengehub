# challengehub
Challenge hub makes it easier for developers to find simple and great code challenges 

## Database managed by hasura
database used in this project is managed by hasura 
**Hasura** is a Graphql api for managing your database database check out more here [Hasura](https://hasura.io)
**Hasura documentation** [doc](https://hasura.io/docs/latest/graphql/core/index/)

## check out Hashnode
[Hashnode](https://www.hashnode.com) platform for developers

## Tech stack and Deployment
**Heroku** for deployment 
**front end** -> html, css & javascript
**backend** -> hasura & golang

## How to use it
**database is managed by hasura graphql api** [Register](https://hasura.io) to get started
update the **HASURA_SECRET** in .env file to your own hasura secret key
in http.go file modify **http.NewRequest("POST", "https://addprojectnamehere.hasura.app/v1/graphql")** 
to your hasura api gateway
start the server with go run main.go
## you are up and running
