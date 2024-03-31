run:
	# @templ generate
	@air -c .air.toml

clean:
	@sudo fuser -k 1234/tcp

createdb:
	touch data/newser.sqlite

dropdb:
	rm data/newser.sqlite

migrateup:
	migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose up

migratedown:
	migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose down