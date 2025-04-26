# Cracker
> Quick wasm functions


```bash
./cracker-runner-darwin-arm64 ./plugin.wasm say_hello 8081
```

```bash
curl -X POST \
http://localhost:8081 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
```

## Run the (local) Compose CI

### Requirements

You need Docker Desktop installed and running (eg: for vulnerability scan with Docker Scout).

Create a `.env` file in the root of the project with the following content:

```bash
DOCKER_HUB_USERNAME=<your_docker_hub_username>
DOCKER_HUB_PAT=<your_docker_hub_pat>
TAG=demo
ABOUT=Cracker project, the HTTP wasm runner
AUTHOR=@k33g
```
### Build and run

```bash
docker compose -f compose.ci.yml up --build
```
