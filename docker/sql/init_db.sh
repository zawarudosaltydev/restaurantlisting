#!/usr/bin/env bash

mysql -uadmin -padmin restaurantlisting < "/docker-entrypoint-initdb.d/create_restaurants.sql"
mysql -uadmin -padmin restaurantlisting < "/docker-entrypoint-initdb.d/seed.sql"
