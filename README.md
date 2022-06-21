# Go Project REST API Example - Hello World

This is an example for Go project.

The motivation behind this project is to learn and broaden my knowledge about programming, APIs, writing test cases for APIs and also containerizing REST API and publishing to Image Registry. 

Here I am taking a different path when I publish the Docker Image. I am making use of the github container registry. Please refer [Github Packages registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

## A bit about me : ) 
[Albert Sebastian](https://www.albertstech.com/) 

## The Project

This is a simple REST API containerized to Docker Image whose test,build and publish is taken care by [Github Actions](https://github.com/features/actions)


This project is using:

1. Golang as the API Programming Language.
2. GitHub Actions to perform the pipeline which contains very basic stages test,build and publish
3. Container Image registry used is gcr.

## Getting Started

This is guide to get started with this project and installing dependencies required for running this project locally

### Requirements

1. Golang
2. Modules and versions are maintained in go.mod file. Mainly records the direct dependency.
3. Direct and indirect dependencies are generated in go.sum. You don't need to modify this file by any chance. 
3. Recommended to have a development einvironment created by using and IDE like VScode with required extensions installed specific to go build and test. 

### Database Dependency 
This project does not have a dependency to a Database. 

### Build and test via commandline

Follow below steps to build and test the application if IDE is not configufred. 
- Checkout the code and cd into `hello-world` directory
- `go test -v` To run the tests. 
- `go build` To build the REST API application

### Accessing the Hello World! application 

- After successful `build`, access the application by URL http://localhost