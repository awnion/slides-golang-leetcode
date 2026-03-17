.PHONY: test

test:
	@for d in problems/*/; do \
		echo "=== $$d ==="; \
		(cd "$$d" && go test ./...); \
	done
