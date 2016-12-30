# Simple

> Dead simple CoC based Serverless CLI tool

**Note:** Simple was Hackathon project to deep dive into AWS SAM. This code will be updated and re-developed in a TDD way so that it's way more clean, better architected and DRY. I'll force-push once I'm done with it (and maybe force push in the meantime as well)!

## Table of contents

- [What is Simple?](#what-is-simple)
- [Getting started](#getting-started)
- [A Simple service structure](#a-simple-service-structure)
- [CLI reference](#cli-reference)
- [But I need more flexibility](#but-i-need-more-flexibility)

## What is Simple?

Simple is a CLI tool written in Go which enables you a dead simple way to deploy Serverless applications.

With Simple you can focus on your function code. Convention over configuration allows Simple to determine what resources to create and how to deploy your service. See more about this in the [getting started](#getting-started) section.

**Note:** Never use Simple in production as it has pretty candid defaults.

```bash
root@11fc58fa7244:/go/src/github.com/pmuens/simple/example-service# ls -lisa
total 4
2817035 0 drwxr-xr-x  4 root root 136 Jan  2 06:26 .
2847500 0 drwxr-xr-x 14 root root 476 Dec 31 12:12 ..
2915578 4 -rw-r--r--  1 root root   8 Dec 30 15:17 .gitignore
2907983 0 drwxr-xr-x  3 root root 102 Dec 30 13:38 timer-hello

root@11fc58fa7244:/go/src/github.com/pmuens/simple/example-service# ls -lisa .simple/
total 12
2970038 0 drwxr-xr-x 5 root root  170 Jan  2 06:29 .
2817035 0 drwxr-xr-x 5 root root  170 Jan  2 06:29 ..
2970039 4 -rwxr-xr-x 1 root root  192 Jan  2 06:29 create-stack.yml
2970041 4 -rwxr-xr-x 1 root root 1520 Jan  2 06:29 example-service.zip
2970040 4 -rwxr-xr-x 1 root root  579 Jan  2 06:29 update-stack.yml

root@11fc58fa7244:/go/src/github.com/pmuens/simple/example-service# go run ../main.go package
[SIMPLE]: Packaging...
[SIMPLE]: Creating create-stack.yml file...
[SIMPLE]: Creating update-stack.yml file...
[SIMPLE]: Creating .zip file...
[SIMPLE]: Compressing /go/src/github.com/pmuens/simple/example-service/.gitignore
[SIMPLE]: Compressing /go/src/github.com/pmuens/simple/example-service/.simple/create-stack.yml
[SIMPLE]: Compressing /go/src/github.com/pmuens/simple/example-service/.simple/example-service.zip
[SIMPLE]: Compressing /go/src/github.com/pmuens/simple/example-service/.simple/update-stack.yml
[SIMPLE]: Compressing /go/src/github.com/pmuens/simple/example-service/timer-hello/handler.js
[SIMPLE]: Successfully packaged service...

root@11fc58fa7244:/go/src/github.com/pmuens/simple/example-service# go run ../main.go deploy
[SIMPLE]: Deploying (this might take a few seconds)...
[SIMPLE]: Creating Changeset...
[SIMPLE]: CREATE_COMPLETE
[SIMPLE]: Creating Stack...
[SIMPLE]: CREATE_IN_PROGRESS
[SIMPLE]: CREATE_IN_PROGRESS
[SIMPLE]: CREATE_IN_PROGRESS
[SIMPLE]: CREATE_COMPLETE
[SIMPLE]: Stack successfully created...
[SIMPLE]: Uploading artifacts...
[SIMPLE]: Creating Changeset...
[SIMPLE]: CREATE_IN_PROGRESS
[SIMPLE]: CREATE_COMPLETE
[SIMPLE]: Updating Stack...
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_IN_PROGRESS
[SIMPLE]: UPDATE_COMPLETE
[SIMPLE]: Done...

root@11fc58fa7244:/go/src/github.com/pmuens/simple/example-service# go run ../main.go remove
[SIMPLE]: Removing (this might take a few seconds)...
[SIMPLE]: Removing artifacts...
[SIMPLE]: Successfully removed artifacts...
[SIMPLE]: Removing Stack...
[SIMPLE]: DELETE_IN_PROGRESS
[SIMPLE]: DELETE_IN_PROGRESS
[SIMPLE]: DELETE_IN_PROGRESS
[SIMPLE]: DELETE_IN_PROGRESS
[SIMPLE]: Done...
[SIMPLE]: Successfully removed service...
```

## Getting started

### Step by step guide

**Note:** Simple will create a deployment bucket with a name based on your services directory. You may want to rename the service so that the deployment bucket won't conflict with other buckets on AWS.

Furthermore Simple uses your `default` profile and deploys to `us-east-1`.

1. Make sure you have installed [Docker](http://docker.com) on your machine
2. Clone this repository
3. `cd` into the repository
4. Run `docker-compose run go-simple bash` to get into a container running Simple
5. Run `go get` to download all dependencies
6. Run `cd example-service` to get into the example service
7. Run `go run ../main.go package` to package the service
8. You can see the output in the `.simple` directory
9. Run `go run ../main.go deploy` to deploy the service (you can run this whenever you've changed something and want to re-deploy)
10. Run `go run ../main.go remove` to remove the service

### Building a binary

Just run the `go build` command to build your own executable binary.

After that you should be able to do something like `./simple package` or `./simple deploy`.

## A Simple service structure

Simple follows a strict convention so that you don't need any configuration.

### Service structure

Take a look at the [example-service](./example-service) to see a fully fledged example service.

**Note:** Directories which don't follow this naming convention will be added to the services zip file but not considered during resource compilation.

```
- <service-name>
  - <event>-<function-name>
    - handler.js (which exports a handler module)
```

### Available events

#### Alexa Skill

- **event:** alexa-skill
- **example:** alexa-skill-greeter/handler.js

#### API

- **event:** api
- **example:** api-greeter/handler.js

#### S3

- **event:** s3
- **example:** s3-greeter/handler.js

#### Timer

*Will run function every 2 minutes*

- **event:** timer
- **example:** timer-greeter/handler.js

#### None

*Just the function. No events will be added*

- **event:** none
- **example:** none-greeter/handler.js

## CLI reference

Usage: `simple [command] [--help]`

### Version

Prints out the version number.

`simple version`

### Package

**Note:** Must be run inside of the service

Creates deployment artifacts in a `.simple` directory in the services directory.

`simple package`

### Deploy

**Note:** Must be run inside of the service

Deploys the service artifacts. You must run `simple package` beforehand to create the deployment artifacts.

`simple deploy`

### Remove

**Note:** Must be run inside of the service

Removes the service. A service must be deployed (e.g. with `simple deploy`) in order to be removable.

`simple remove`

## But I need more flexibility

Simple aims to make things as simple as possible. It should help you explore the Serverless world.

Please look at [Serverless](http://github.com/serverless/serverless) if you want a feature rich, fully fledged Serverless tool.
