#!/bin/bash

BASE_URL="http://localhost:8080"

function do_request() {
  method=$1
  url=$2
  data=$3

  echo "=== $method $url ==="

  if [ -z "$data" ]; then
    response=$(curl -s -w "\nHTTP_STATUS:%{http_code}" -X "$method" "$url")
  else
    response=$(curl -s -w "\nHTTP_STATUS:%{http_code}" -X "$method" "$url" \
      -H "Content-Type: application/json" -d "$data")
  fi

  body=$(echo "$response" | sed -e '/HTTP_STATUS:/d')
  status=$(echo "$response" | grep HTTP_STATUS | cut -d: -f2)

  echo "Status: $status"
  echo "Body:"
  echo "$body"
  echo -e "\n"
}

do_request POST "$BASE_URL/quotes" '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
do_request GET "$BASE_URL/quotes"
do_request GET "$BASE_URL/quotes/random"
do_request GET "$BASE_URL/quotes?author=Confucius"
do_request DELETE "$BASE_URL/quotes/1"
do_request GET "$BASE_URL/quotes"
