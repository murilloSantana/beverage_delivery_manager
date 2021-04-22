# Beverage Delivery Manager

### Introduction
The Beverage Delivery Manager (**BDM**) was developed in Golang and exposed to external services with Graphql. The project was structured with the goal of keeping the business rules decoupled from the rest of the code, all communication with "details" is done through abstractions. Redis was used to avoid overloading the database read/write and also the network traffic with APQ. The database used was MongoDB, because besides being performant, mongo brings a great simplicity in the manipulation of GEOJSON data that are the base of the business rules of the challenge. 

### How to run the project locally
Run the `make run-docker` command to start the application (must be run in the project root)
##### Prerequisites
- [Docker 19.x](https://docs.docker.com/engine/install/)
- [Docker-compose 1.26.0](https://docs.docker.com/compose/install/)
    
### Testing the Features
The playground is enabled for local execution, access it through the url http://localhost:5000/ (this feature isn't enabled in production for security reasons)

#### Create a partner

```
mutation($input: PdvInput!) {
  savePdv(input: $input) {
    id
    tradingName
    ownerName
    document
    coverageArea
    address
  }
}
```

***Query variables***
```
{
	"input": {
		"tradingName": "Adega da Cerveja - Pinheiros",
		"ownerName": "Zé da Silva",
		"document": "1432132123891/0001",
		"coverageArea": {
			"type": "MultiPolygon",
			"coordinates": [
				[
					[
						[30, 20],
						[45, 40],
						[10, 40],
						[30, 20]
					]
				],
				[
					[
						[15, 5],
						[40, 10],
						[10, 20],
						[5, 10],
						[15, 5]
					]
				]
			]
		},
		"address": {
			"type": "Point",
			"coordinates": [29.355468750000004, 11.005904459659451]
		}
	}
}
```

#### Load partner by id

```
query($input: PdvIdInput!) {
  findPdvById(input: $input) {
    id
    tradingName
    ownerName
    document
    coverageArea
    address
  }
}
```

***Query variables***
```
{
  "input": {
    "id": "<some id>"
  }
}
```
***OBS***: Fill in "some id" in the field above, as it is a dynamically generated field it isn't possible to predict any valid option  

#### Search partner

```
query($input: PdvAddressInput!) {
  findPdvByAddress(input: $input) {
    id
    tradingName
    ownerName
    document
    coverageArea
    address
  }
}
```

***Query variables***
```
{
  "input": {
    "longitude": 30.355468750000004,
    "latitude":  11.005904459659451
  }
}
```

### Useful Commands
- `make run-docker`: Starts the application and all its dependencies (Redis and MongoDB) inside docker containers
- `make stop-docker`: Stop execution of the application and all its dependencies (Redis and MongoDB) that are in docker containers
- `make remove-docker`: Removes all containers created by the `make run-docker` command
- `make run-test`: Runs unit and integration testing of the application

### Step by step to put the project into production
- Set up a MongoDB server with a database called beverageDeliveryManagerDB, you also need to create a [2dsphere](https://docs.mongodb.com/manual/core/2dsphere/) index for the **address** and **coverageArea** fields
- Setting up a Redis server
- Configure the **BDM**. The project is [dockerized](docker/Dockerfile "dockerized"), so it is advisable to deploy the application in some service or tool that supports docker containers. The project uses github actions to perform a series of validations and automatically generate releases, an integration with the [docker hub](https://hub.docker.com/) was configured where a new image with the **BDM** is generated each new version generated in the main branch, so to deploy the **BDM** in your production environment, just point your container service to this [image](https://hub.docker.com/r/murillosantana/beverage_delivery_manager), it is also necessary to configure the following environment variables:
    
    - PORT=5000
    - ENV=production
    - MONGO_DB_NAME=beverageDeliveryManagerDB
    - MONGO_COLLECTION_NAME=pdvs
    - MONGO_URL=<MONGO SERVER URL WITH THE PORT: Ex mongodb://localhost:27017>
    - MONGO_MIN_POOL_SIZE=10
    - MONGO_MAX_POOL_SIZE=30
    - MONGO_MAX_CONN_IDLE_TIME=10  
    - REDIS_URL: <REDIS SERVER URL WITH THE PORT: Ex localhost:6379>
    - REDIS_PASSWORD: <YOUR REDIS SERVER PASSWORD>
    - REDIS_DB: 0
    - REDIS_MIN_IDLE_CONN=10
    - REDIS_POOL_SIZE=30   