npm run build

GITHUB_REPOSITORY=kwanok/spatial-query-study/frontend
DOCKER_IMAGE_URL=ghcr.io/$(echo $GITHUB_REPOSITORY | tr '[:upper:]' '[:lower:]')
GITHUB_SHA=$(git rev-parse HEAD)
DOCKER_IMAGE_TAG=$(echo $GITHUB_SHA | cut -c1-7)
docker build -t $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG .
docker tag $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG $DOCKER_IMAGE_URL:latest

docker rm -f frontend
docker run -p 3987:80 --name frontend $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG