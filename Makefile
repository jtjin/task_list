.PHONY : doc 
doc:
	swag init

.PHONY : run 
run:
	go run main.go

.PHONY : test 
test:
	go test ./...

.PHONY : seed-flush 
seed-flush:
	# flush mysql
	docker exec mysql mysql -uroot -psecret -e \
	"SELECT CONCAT('TRUNCATE TABLE ', table_schema, '.', TABLE_NAME, ';') FROM INFORMATION_SCHEMA.TABLES \
	WHERE table_schema IN ('task_list') AND TABLE_NAME != 'migrations'" | grep "task_list*" | xargs -I {} docker exec mysql mysql -uroot -psecret -e {}
	# exec seeder
	go run ./cmd/seeder/main.go	

.PHONY : gen-mock 
gen-mock:
	mockgen -source=./internal/task/repository.go -destination=./mock/task/task_repository_mock.go -package=task_mock