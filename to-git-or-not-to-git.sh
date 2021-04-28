#! /bin/bash
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"chalvoker\"}}){id}}"}' | jq '.data.user[0].id'
