pull_dump:
	@echo "Start pulling fresh dump..."
	docker exec beria_redis redis-cli save
	docker cp beria_redis:/app/dump.rdb $(shell pwd)/backup/
	@echo "Success!"

deploy_dump:
	@echo "Start deploying dump..."
	docker cp $(shell pwd)/backup/dump.rdb beria_redis:/app/
	docker restart beria_redis
