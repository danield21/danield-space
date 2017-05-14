# danield.space

I am Daniel J Dominguez, and this is the source code for my personal site

## Getting Started

This project is a simple website that will be used for myself.

### Prerequisites

In order to run this, you will need to install a few pieces of software

* [Go](https://golang.org/dl/)
* [NodeJS](https://nodejs.org)
* [Google Cloud SDK](https://cloud.google.com/appengine/docs/standard/go/download)

### Installing

```
go get ./...
```

```
npm install
```

```
dev_appserver app
```

#### Optional
Although you can use the local installation in the node_modules, you can also install these globally
```
npm install -g webpack eslint
```

## Running the tests

The tests for Go will can be run using
```
go test server/...
```
Tests are written using [Testify](https://github.com/stretchr/testify)

The tests for JavaScript will can be run using
```
npm test
```
Tests haven't been written yet

## Deployment

Deployment is done through the gloud application
```
gcloud deploy
```

## TODO

* Turn repository into an interface rather than just package of functions
* Fix any broken tests
* Add more tests to backend
* Introduce testing framework for JavaScript
* Add edit and delete capabilities for articles
* Add edit and delete capabilities for categories
* Move SPA HTML information into a storage rather than in state
  * state:
    * id - Used to retrieve page some storage
    * scroll position - Used to emulate browser behavior
    * title - Used for updating the title
    * meta - Used for updating any meta tags
* Move all text into a resource file for multiple languages
* Implement a headless CMS

## Built With

* [Gorilla](https://github.com/gorilla) - Router and session 
* [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML sanitizing
* [Testify](https://github.com/stretchr/testify) - Useful functions for Go Testing 
* [Google App Engine](https://cloud.google.com/appengine) - Cloud service that site is built for

## Authors

* **Daniel J Dominguez** [danield21](https://github.com/danield21)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
