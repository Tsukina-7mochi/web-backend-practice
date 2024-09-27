#!/bin/bash
set -euo pipefail
cd "$(dirname "$0")"

PROTOCOL="${PROTOCOL:-http}"
HOST="${HOST:-localhost:8080}"
ORIGIN="$PROTOCOL://$HOST"

echo "Ping"
curl -sSL "$ORIGIN/ping"

echo -e "\n\nCreate a user"
curl -sSL \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"name": "user1", "display_name": "User 1"}' \
    "$ORIGIN/users"

echo -e "\n\nGet the user"
curl -sSL \
    -X GET \
    "$ORIGIN/users/user1"

echo -e "\n\nCreate todos"
curl -sSL \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"title": "Todo 1"}' \
    "$ORIGIN/users/user1/todos"
echo ""
curl -sSL \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"title": "Todo 2"}' \
    "$ORIGIN/users/user1/todos"
echo ""
curl -sSL \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"title": "Todo 3"}' \
    "$ORIGIN/users/user1/todos"

echo -e "\n\nGet the todos"
todos="$(curl -sSL \
    -X GET \
    "$ORIGIN/users/user1/todos")"
echo "$todos"

todo_ref_1="$(echo "$todos" | jq -r '.[0].ref')"
todo_ref_2="$(echo "$todos" | jq -r '.[1].ref')"
todo_ref_3="$(echo "$todos" | jq -r '.[2].ref')"

echo -e "\n\nUpdate todos"
curl -sSL \
    -X PATCH \
    -H "Content-Type: application/json" \
    -d '{"done": true}' \
    "$ORIGIN/users/user1/todos/$todo_ref_1"
echo ""
curl -sSL \
    -X PATCH \
    -H "Content-Type: application/json" \
    -d '{"done": true}' \
    "$ORIGIN/users/user1/todos/$todo_ref_3"

echo -e "\n\nGet the todos"
curl -sSL \
    -X GET \
    "$ORIGIN/users/user1/todos"

echo -e "\n\nDelete a todo"
curl -sSL \
    -X DELETE \
    "$ORIGIN/users/user1/todos/$todo_ref_3"

echo -e "\n\nGet the todos"
curl -sSL \
    -X GET \
    "$ORIGIN/users/user1/todos"

echo -e "\n\nDelete the user"
curl -sSL \
    -X DELETE \
    "$ORIGIN/users/user1"
