SVCUSER = wex_dev
SVCPASS = password

test-integration:
	go test -count=1 -p=1 -tags=integration -v ./...

test-unit:
	go test -count=1 ./...

test-clean-cache:
	go clean -testcache

serve:
	go run main.go --config=./config/config.yaml serve

postgres:
	docker-compose up -d postgres
	sleep 15
	make migrate

postgres-clean:
	docker rm -f postgres

migrate:
	docker run --network wex_basic --rm \
	-v ${PWD}/deployment:/deployment \
	-e PGPASSWORD=password \
	        postgres:15-alpine psql \
		-h postgres \
		-U postgres \
		-f /deployment/dependencies.sql \
		-v _user=postgres \
		-v _dbpass=${SVCPASS} \
		-v _dbuser=${SVCUSER}

	# run the flyway migration
	docker run --network wex_basic --rm \
	-v ${PWD}/deployment/migrations:/flyway/sql \
		flyway/flyway:9-alpine \
		-url=jdbc:postgresql://postgres:5432/postgres \
		-user=postgres \
		-password=password \
		-schemas=wex \
		-placeholders.service_user=${SVCUSER} \
		migrate

mocks:
	mockgen -package mocks -source internal/data/pgx/interface.go -destination test/mocks/mock_pgx_interface.go
	mockgen -package mocks -source internal/exchangerate/interface.go -destination test/mocks/mock_retriever_interface.go