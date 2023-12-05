#!/bin/bash

# Define the list of services
services=(
"user-microservice"
"offer-microservice"
#"product-management" - TODO: implement and fix migrate/main.go and seed/main.go
"order-microservice"
)

# Save the current working directory
cwd=$(pwd)

# Process each service
for service in "${services[@]}"
do
    echo "Processing $service"

    # Change directory to the service folder
    cd "$cwd/$service"

    # Migrate
    echo "Migrating $service"
    go run cmd/migrate/main.go
    if [ $? -ne 0 ]; then
        echo "Migration of $service failed"
        exit 1
    fi

    # Seed
    echo "Seeding $service"
    go run cmd/seed/main.go
    if [ $? -ne 0 ]; then
        echo "Seeding of $service failed"
        exit 1
    fi

    echo "$service processed successfully"
done

# Change directory back to the original working directory
cd "$cwd"