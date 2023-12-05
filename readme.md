# microservice-challenge-order-api

### Environment

Each microservice has an .env.dist file which shows the necessary environment variables you need to provide.
1. Create an .env.dev file for use in development environment
2. Create an .env.prod file for use in production environment

### Offer Environment
.env.dev
 ```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=localhost
MYSQL_DATABASE=offerdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3302

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATOR_MICROSERVICE_URL=http://localhost:3001

# Router Configuration
GIN_PORT=3002
GIN_HOST=localhost

# Application Configuration
TZ=europe/paris
ENV=dev
API_KEY=secret
```
.env.prod
```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=host.docker.internal
MYSQL_DATABASE=offerdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3302

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATOR_MICROSERVICE_URL=http://host.docker.internal:3001

# Router Configuration
GIN_PORT=3002
GIN_HOST=0.0.0.0

# Application Configuration
TZ=europe/paris
ENV=prod
API_KEY=secret
```

### Order Environment
.env.dev
 ```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=localhost
MYSQL_DATABASE=orderdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3303

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATOR_MICROSERVICE_URL=http://localhost:3001

# Microservice Configuration
OFFER_MICROSERVICE_URL=http://localhost:3002

# Router Configuration
GIN_PORT=3003
GIN_HOST=localhost

# Application Configuration
TZ=europe/paris
ENV=dev
API_KEY=secret
```
.env.prod
```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=host.docker.internal
MYSQL_DATABASE=orderdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3303

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATOR_MICROSERVICE_URL=http://host.docker.internal:3001

# Microservice Configuration
OFFER_MICROSERVICE_URL=http://host.docker.internal:3002

# Router Configuration
GIN_PORT=3003
GIN_HOST=0.0.0.0

# Application Configuration
TZ=europe/paris
ENV=prod
API_KEY=secret
```

### Product Environment
.env.dev
```
# Database Configuration
MYSQL_HOST=localhost
MYSQL_TCP_PORT=3306
MYSQL_PORT=3304
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_DATABASE=productdb
MYSQL_ROOT_PASSWORD=myrootpassword

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATION_API_KEY=secret
AUTHENTICATOR_MICROSERVICE_URL=http://localhost:3001

# Router Configuration
GIN_PORT=3004
GIN_HOST=localhost

# Application Configuration
TZ=europe/paris
ENV=dev
```
.env.prod
```
# Database Configuration
MYSQL_HOST=host.docker.internal
MYSQL_TCP_PORT=3306
MYSQL_PORT=3304
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_DATABASE=productdb
MYSQL_ROOT_PASSWORD=myrootpassword

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATION_API_KEY=secret
AUTHENTICATOR_MICROSERVICE_URL=http://host.docker.internal:3001

# Router Configuration
GIN_PORT=3004
GIN_HOST=0.0.0.0

# Application Configuration
TZ=europe/paris
ENV=prod
```

### User Environment
.env.dev
```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=0.0.0.0
MYSQL_DATABASE=userdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3301

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATION_API_KEY=secret

# Router Configuration
GIN_PORT=3001
GIN_HOST=0.0.0.0

# Application Configuration
TZ=europe/paris
ENV=dev
```
.env.prod
```
# Database Configuration
DATABASE_DRIVER=mariadb
MYSQL_HOST=host.docker.internal
MYSQL_DATABASE=userdb
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_ROOT_PASSWORD=myrootpassword
MYSQL_TCP_PORT=3306
MYSQL_PORT=3301

# Authentication Configuration
AUTHENTICATOR_DRIVER=jwt
AUTHENTICATION_API_KEY=secret

# Router Configuration
GIN_PORT=3001
GIN_HOST=0.0.0.0

# Application Configuration
TZ=europe/paris
ENV=prod
```

### Setup Docker Environment

2. Deploy all docker containers.
```
docker-compose up -d
```

3. Run migrations and seeds
```
/bin/bash prepare.sh
```

### Routes and additional documentation
Use postman and import the ```postman_collection.json``` for further information regarding routes, requests and responses.