# Service composition

> The docker-compose.yml file can be found [here](../docker-compose.yml).

## R1

### Cluster

It is composed of three services. Those are:

- HarvestCCode. This one encapsulates the software. Created from the Dockerfile.hcc file.

- Nginx. This service works as a proxy, forwarding all the requests issued to ports 80 and 443 to HarvestCCode. It also adds a layer of security by implementing and using a SSL certificate, so all the requests are performed. More info about the configuration below.

- MongoDB. This database service is used as-is, with default configuration. Currently it has no data persistance, but that property could be achieved by using a volume in the docker-compose file.

#### Discussion

##### Why Nginx?

There are a lot of alternatives when it comes to proxys, like HAProxy, lighttpd or traefik and the [performance is quite similar](https://www.loggly.com/blog/benchmarking-5-popular-load-balancers-nginx-haproxy-envoy-traefik-and-alb/). It has a lot of features, like SSL implementation, logging or load balancing, so in the case it is required to scalate the HarvestCCode service I only need to add a couple of lines in the configuration file.

The SSL certificates are located in the `/certs` directory, and those should be regenerated if this compose is going to be used somewhere else.

After testing some of them I decided to use Nginx, since I've already used it in some [other projects](https://github.com/harvestcore/tfg) and I'm used to work with its main configuration file.

This is its configuration file content.

```nginx
# Http server. This server listens to port 80 and forwards
# all the requests to port 443.
#
# Server name is set to `_` since we want to forward all
# the requests.
server {
    listen      80;
    server_name _;

    return 301 https://$host$request_uri;
}

# Https server. This one listens to port 443.
#
# The server name is set to `localhost`, but it should be
# changed to the actual hostname to be used.
server {
    listen      443 ssl;
    server_name localhost;

    # SSL certificate files.
    ssl_dhparam             /etc/nginx/certs/dhparam.pem;
    ssl_certificate         /etc/nginx/certs/default.crt;
    ssl_certificate_key     /etc/nginx/certs/default.key;

    # Allowed protocols and cipher methods.
    ssl_protocols           TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers             HIGH:!aNULL:!MD5;

    # Proxy all requests to http://harvestccode:8080.
    #
    # Note that the hostname of the service is the one
    # used in the docker-compose file.
    #
    # http://harvestccode:8080 is where HarvestCCode is
    # listening.
    location / {
        proxy_pass http://harvestccode:8080;

        # Add some headers so the HarvestCCode service
        # can know the real IP and hostname of the client
        # that performed the request.
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
    }
}
```

##### Why there is no data persistance?

This service composition (a.k.a [docker-compose.yml](../docker-compose.yml)) aims to be the backbone of a bigger and more complex composition. Since this compose file is primarily going to be used for [testing purposes](../.github/workflows/benchmark.yml), there is no need to add data persistance. If required, the only step required is to add a volume in the service as shown below:

##### Why the composition includes a network declaration?

In fact, there is no need to declare a network in a docker-compose file since all the services share the same network. I created it as a way to keep track of all services IP addresses, and also to make 100% sure that all of the services share the same network. In the past, working on [other projects](https://github.com/harvestcore/tfg) I had a lot of networking issues and I opted to manually assign the IP addresses to all the services. This is not required but for me it makes more sense to have it this way. (It also makes debugging sessions more easy).

```yml
  mongo:
    image: mongo:4.4.2
    restart: always
    volumes:
      - /path/to/db:/data/db
```

## R2 & R3

This section explains the configuration of each service in the [docker-compose.yml](../docker-compose.yml) file and also the network I'm using.

### Networking

```yml
networks:
  hcc:
    driver: bridge
    ipam:
      config:
        - subnet: 172.25.0.0/16
```

- `hcc` - This indicates the name of the network, `hcc` in this case.
- `driver: bridge` - The driver this network will use. `bridge` allows all containers that are connected to the same network to communicate between them and also provices isolation, so all the services that are not connected to this network can't "see" the services.
- `ipam` - This stands for "IP Address Management", which means that some configuration will take place.
- `subnet` - The subnet this network will be using, in this case `172.25.0.0/16`.

### HarvestCCode

```yml
  harvestccode:
    build:
      context: .
      dockerfile: Dockerfile.hcc
    restart: always
    expose:
      - 8080
    environment:
      - HCC_MONGO_URI=mongodb://mongo:27017
    healthcheck:
      test: curl --fail -s http://localhost:8080/api/healthcheck || exit 1
      interval: 30s
      timeout: 10s
      retries: 3
    networks:
      hcc:
        ipv4_address: 172.25.0.3
    depends_on:
      - mongo
```

- `build` - Build the image to be used from the Dockerfile named `Dockerfile.hcc` whose context is the current one (`.`). This key could be changed to `image`, since the image for HarvestCCode is hosted in both DockerHub and GH Registry. (If so, the value of this `image` key would be either `harvestcore/harvestccode-backend:latest` or `ghcr.io/harvestcore/harvestccode-backend:latest`).
- `restart: always` - Always restart the service if it goes down for some reason.
- `expose: 8080` - Expose the port 8080 to the network. In this case there is no need to do port mapping (like it happens in the `nginx` service).
- `environment` - Environment variables. In this case only `HCC_MONGO_URI` is set. Note that the hostname in the URI matches the name of the service in the compose file, this is due to all services having as hostname the same value it is set in the compose.
- `healtcheck` - This allows the health check of the service. In this case it issues a request to the `/api/healthcheck` endpoint every 30 seconds (with a max timeout of 10s and 3 retries). If the request is unsuccessful the service will be restarted automatically.
- `networks` - Set the network to be used. In this case the `hcc` network with IP `172.25.0.3`.
- `depends_on` - This service depends on the mongo service, so it will be started as soon as the data service is up and running.

### Nginx

```yml
  nginx:
    image: nginx:alpine
    restart: always
    ports:
      - 80:80
      - 443:443
    volumes:
      - ./nginx.conf:/etc/nginx/conf.d/default.conf
      - ./certs:/etc/nginx/certs
    depends_on:
      - harvestccode
    networks:
      hcc:
        ipv4_address: 172.25.0.4
```

- `image: nginx:alpine` - Base image. A lightweight Alpine image with Nginx installed.
- `restart: always` - Always restart the service if it goes down for some reason.
- `ports` - Map the host machine ports with the service ports. In this case 80 and 443.
- `volumes` - Mount the volumes needed. The first one contains the nginx configuration file. The second one the SSL certificate keys.
- `depends_on` - This service depends on the harvesttcode service, so it will be started as soon as it is up and running.
- `networks` - Set the network to be used. In this case the `hcc` network with IP `172.25.0.4`.

### Mongo

```yml
  mongo:
    image: mongo:4.4.2
    restart: always
    networks:
      hcc:
        ipv4_address: 172.25.0.2
```

- `image: mongo:4.4.2` - Base image. Mongo version 4.4.2.
- `restart: always` - Always restart the service if it goes down for some reason.
- `networks` - Set the network to be used. In this case the `hcc` network with IP `172.25.0.4`.

## R4

There are two scenarios when it comes to testing the service composition:

- Create my own tests, something like a script in some language.
- Use a framework or tool that (probably) can measure better the performance and also will also keep in mind all the metrics.

In this case I've followed the second route. After a quick search in Google I found that there are quite a lot of tools that can do this task:

- [Gatling](https://gatling.io/): This one is really interesting since it has a lot of features, but it is more focused on internal performance of the code. It has capabilities to test endpoints, but not as deep as other alternatives in this list. Its sytax is also a bit complex compared to others and it is only available for Java.
- [Locust](https://locust.io/): Designed to test web applications. The best 
- K6

## R5
