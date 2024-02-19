run:
	@templ generate
	@air -c .air.toml

clean:
	@sudo fuser -k 1234/tcp