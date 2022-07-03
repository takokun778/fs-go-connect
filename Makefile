.PHONY: gen

lint:
	@buf lint
gen:
	@buf generate
dev:
	@air
