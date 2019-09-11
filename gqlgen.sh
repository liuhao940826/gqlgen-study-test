#!/bin/sh

echo "Generating code from GraphQL schema..."

cd "$(dirname "$0")"

cd ./graphql
go run github.com/99designs/gqlgen