build-image:
	cd ./deployments && docker-compose build

run-app:
	cd ./deployments && docker-compose up

run-app-bg:
	cd ./deployments && docker-compose up -d

bash:
	cd ./deployments && docker-compose exec app bash
