include .env
export $(shell sed 's/=.*//' .env)

export MODEL_DIR=./internal/model/mysql/mysqlpetland
export DDL_DIR=./internal/migration/petland
export MYSQL_URL="mysql://root:123456@tcp(127.0.0.1:33061)/petland?"

run:
	go run main.go

# make create-migration name=file_name
create-migration:
	migrate create -ext sql -dir ${DDL_DIR} -seq ${name}

# make migration-up num=1
migrate-up:
	migrate -database ${MYSQL_URL} -path ${DDL_DIR} up ${num}

# make migration-down num=1
migrate-down:
	migrate -database ${MYSQL_URL} -path ${DDL_DIR} down ${num}

gen-model:
	gentool -modelPkgName mysqlpetland -onlyModel -outPath ${MODEL_DIR} -c "./gen.yml"