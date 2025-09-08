#!/usr/bin/env bash

$BASE_URL=http://localhost:8000

http POST $BASE_URL/api/login/ \
  email=peter@example.com \
  password=^uw6fJ3NY45peX