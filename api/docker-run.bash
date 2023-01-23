GITHUB_REPOSITORY=kwanok/spatial-query-study/api-server
DOCKER_IMAGE_URL=ghcr.io/$(echo $GITHUB_REPOSITORY | tr '[:upper:]' '[:lower:]')
GITHUB_SHA=$(git rev-parse HEAD)
DOCKER_IMAGE_TAG=$(echo $GITHUB_SHA | cut -c1-7)
docker build -t $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG .
docker tag $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG $DOCKER_IMAGE_URL:latest

docker rm -f api-server
docker run -p 2586:8080 --name api-server $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG