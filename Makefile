.PHONY: postgres adminer

postgres: 
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=password postgres

adminer:
  docker run --rm -ti --network host adminer
