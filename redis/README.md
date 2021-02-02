# Redis DB

## Params

   * **local**🐳

      * **host**: `localhost`
      * **port**: `6379`
      * **exposed container port**: `6380`

   * **remote test**

      * **root user**: `beria`
      * **root pass**: `dd721260bfe1b3d9601e7fbab36de6d04e2e67b0ef1c53de59d45950db0dd3zc`
      * **host**: `http://apple-pie.fun`
      * **port**: `6379`
      * **exposed container port**: `6380`

## Tools🐳

 * `docker exec -it beria_redis sh` enter container
 * `docker logs -f beria_redis` watch logs
 * `docker exec -it beria_redis redis-benchmark -q -n 100000 -c 50 -P 12` performance banchmark (Выполнит 100 000 запросов от 50 клиентов по 12 команд одновременно)
 * `make pull_dump` download fresh dump
 * `make deploy_dump` deploy dump

## Docs

 * [DB models](./SCHEMA_DOC.md)    
