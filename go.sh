#!/bin/sh

GO_ENV=""

deploy_dev()
{
  echo "running DEV" | tr "[a-z]" "[A-Z]"
  docker-compose -f ./config/docker-compose.dbDev.yml up -d

  sleep 4

  docker ps -a


  docker-compose  -f ./config/docker-compose.dbDev.yml -f ./config/docker-compose.dev.yml up -d

  sleep 6
  export ENDPOINT=http://localhost:7007

  prisma deploy
  echo "========== READY RUNNING DEV ===========" | tr "[a-z]" "[A-Z]"
}

deploy_prod()
{
  echo "running PRODUCTION" | tr "[a-z]" "[A-Z]"
  docker-compose -f ./config/docker-compose.dbProd.yml up -d

  sleep 4

  docker ps -a

  docker-compose -f ./config/docker-compose.dbProd.yml -f ./config/docker-compose.prod.yml up -d

  sleep 6
  export ENDPOINT=http://localhost:7000

  prisma deploy

  echo "========== READY RUNNING PRODUCTION ===========" | tr "[a-z]" "[A-Z]"
}

if [ "$1" = "dev" ]
then
  deploy_dev
elif [ "$1" = "prod" ]
then
  deploy_prod
else
  read -p "what running do you want ? prod or dev ?" GO_ENV
  if [ ${GO_ENV} = "dev" ]
  then
    deploy_dev
  elif  [ ${GO_ENV} = "prod" ]
  then
    deploy_prod
  else
    echo "can't found ${GO_ENV}, try again"| tr "[a-z]" "[A-Z]"
  fi
fi
