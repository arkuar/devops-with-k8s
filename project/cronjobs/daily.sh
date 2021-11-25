#!/usr/bin/env sh
set -e

if [[ -n $URL ]]; then
    TODO=$(curl -sI https://en.wikipedia.org/wiki/Special:Random | sed -En 's/^location: (.*)\r/\1/p')

    curl -sS -X POST $URL/todos \
         -H 'Content-Type: application/json' \
         -d '{"content":"Read '${TODO}'"}' \
         -o /dev/null

    echo "Sent new todo to backend"
else
    echo "Backend URL not found"
fi