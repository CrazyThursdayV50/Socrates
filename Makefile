DOCKER_DIR=./docker
DOCKERFILE_APP=${DOCKER_DIR}/Dockerfile
COMPOSE_FILE_APP=${DOCKER_DIR}/compose.yml
ENV_FILE_APP=${DOCKER_DIR}/.env
APP=socrates

install:
	@go install .

docker-deploy:
	@docker buildx build --load --platform linux/amd64,linux/arm64 -f ${DOCKERFILE_APP} -t ${APP}:local .
	./scripts/shell/docker/deploy.sh ${APP} achillesss/${APP}
