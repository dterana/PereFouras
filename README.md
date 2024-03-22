# Fouras
A simple random 32 bytes key generator

## Requirements
**Fouras** rely exclusively on standard library; so it only require to have go installed.
This program was developed using go version 1.22.1

## Usage
- `go run main.go` to simply generate a new key
- `go install` to install Fouras; then `fouras` to generate a new key

## Possible improvement
- Make Fouras be able to generate different key size based on different characters sets
- Make possible to use an interractive interface

## Context
**Fouras** was to generate a [Gorilla CSRF](https://github.com/gorilla/csrf) [auth-key](https://github.com/gorilla/csrf?tab=readme-ov-file#examples). Thanks to golang *stdlib*, It was simplier than searching for an existing solution.

## Name origin
It is a reference to ***Père Fouras*** from the french game show ***[Fort Boyard](https://en.wikipedia.org/wiki/Fort_Boyard_(game_show))***.
He used to reward by a key a good answer to his riddle