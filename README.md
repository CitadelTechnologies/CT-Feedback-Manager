Citadel Technologies - Feedback Manager
=======================================

This repository contains a Golang API to manage projet feedbacks.

The API is wrapped inside a dedicated Docker container.

Production Usage
----------------

For production environment, you just have to pull the Docker image from Docker Hub and then run the container with the required environment variables.

```
docker pull citadeltechnologies/feedback-manager
```

The required environment variables are:

```
MONGO_DBNAME=feedbacks
MONGO_HOST=your_mongo_container_hostname
MONGO_PORT=27017
SERVER_PORT=80
```

Development Usage
-----------------

After cloning the project, you can apply changes to the code.

In order to compile it, just build the container using the following command:

```
docker build -t citaldetechnologies/feedback-manager .
```

Don't forget that you will need a MongoDB server to run the container.
