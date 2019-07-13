#! /usr/bin/env bash

if [[ "$1" = "buildup" ]]; then
    docker-compose build && docker-compose up -d
    docker-compose ps

    exit;
elif [[ "$1" = "build" ]]; then
    docker-compose build

    exit;
elif [[ "$1" = "up" ]]; then
    docker-compose up -d

    exit;
elif [[ "$1" = "down" ]]; then
    docker-compose down

    exit;
elif [[ "$1" = "logs" ]]; then
    docker-compose logs

    exit;
elif [[ "$1" = "test" ]]; then
    ginkgo -r

    exit;
fi

