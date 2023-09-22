# edit this to match your env dev|prod|test etc
ENV=dev

spinup:
	bash -c "sudo docker compose --project-directory=infrastructure/$(ENV) up -d"
dev:
	bash -c "air -c .air.toml"
build:
	bash -c "echo Build command comes heres"
test:
	bash -c "echo Test command comes here"
test-e2e:
	bash -c "echo e2e test command comes here"
test-cov:
	bash -c "echo code coverage test command comes here"