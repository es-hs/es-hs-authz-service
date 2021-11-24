run:
	docker-compose -f docker-compose.yml --env-file .env up es-hs-authz
build:
	docker build -t $IMAGE_REPO_NAME:$IMAGE_TAG .
watch:
	./scripts/app.sh