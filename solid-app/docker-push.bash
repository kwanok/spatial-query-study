npm run build

GITHUB_REPOSITORY=kwanok/spatial-query-study/frontend
DOCKER_IMAGE_URL=ghcr.io/$(echo $GITHUB_REPOSITORY | tr '[:upper:]' '[:lower:]')
GITHUB_SHA=$(git rev-parse HEAD)
DOCKER_IMAGE_TAG=$(echo $GITHUB_SHA | cut -c1-7)
LATEST=latest
docker build -t $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG -t $DOCKER_IMAGE_URL:$LATEST .
docker tag $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG $DOCKER_IMAGE_URL:$LATEST

echo $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG
docker push $DOCKER_IMAGE_URL:$DOCKER_IMAGE_TAG
docker push $DOCKER_IMAGE_URL:$LATEST
