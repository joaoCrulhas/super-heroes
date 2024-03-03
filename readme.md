# Dee See Super-Heroes

## Table of Contents

- [Dee See Super-Heroes](#dee-see-super-heroes)

- [Table of Contents](#table-of-contents)

- [Introduction](#introduction)

- [Features](#features)

- [Getting Started](#getting-started)

- [Prerequisites](#prerequisites)

- [Installation](#installation)

- [Usage](#usage)

## Introduction

<a href="#dee-see-super-heroes"></a>
Dee See Super-Heroes is a project developed by **_DeeSee GmbH_** to ensure the security and anonymity of superheroes.

The project allows the storage and retrieval of superhero data while encrypting their identities for anyone outside the company.

Only those within **_DeeSee GmbH_** have access to the real identities of the superheroes

The project was developed in [Golang](https://go.dev/), using as framework the [flamingo](https://docs.flamingo.me/). The database is currently implemented an in-memory db, and Docker to [dockerize](https://www.docker.com/) the application

The main goal of this project is created a webservice that allow us, from **_DeeSee GmbH_**, store new super-hero part of our company, and based on super-powers, retrieve them for a specific mission, or return a list of all super-heroes.

## Features

<a href="#features"></a>

- Store SuperHero: This feature allow us of **DeeSee GmbH** the opportunity to add new super-heroes into our database, to do this, we should use our admin credentials, and the example request can be found on the swagger

- Get all SuperHeroes: This is one of the most important feature developed, with this feature, we're able to get all super-heroes registered into our database, and to avoid that some villain have access to the secret identity of our super-heroes, the only way to get the real identity is using the DeeSee credentials.

- Get Heroes by SuperPower: Based on a mission, a specific super-power is necessary, so for this reason is possible fitler the super-hero based on his superPower.

## Getting Started

<a href="#getting-started"></a>
To run this project the easisest way to do it, is use docker, and import the swagger into your favorite rest-api-client(insomnia, postman), to have access in all features available.

### Prerequisites

<a href="#pre-requsites"></a>

#### With Docker

```
docker build -t super-heroes .
docker run -d -p 3222:3222 --rm -it --network=host super-heroes:latest
```

#### Without docker

```
go mod download
go run main.go serve
```

## Usage

<a href="#usage"></a>

The best way to use this project is loading the [file](https://github.com/joaoCrulhas/omnevo-super-heroes/blob/main/swagger/super-hero.yaml), into your insomnia or another Rest client application.

There 2 files into [http-request](https://github.com/joaoCrulhas/omnevo-super-heroes/tree/main/http-requests) that can help you to have a nice overview about the avialable endpoints, and if you're using vs-code, just need to install any plugins that allows you make execute the file, [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client), is one of them.
