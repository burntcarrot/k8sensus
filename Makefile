MAKEFLAGS += --always-make

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

apply: ## Apply definitions.
	kubectl apply -f k8s/rbac.yaml
	kubectl apply -f k8s/deployment.yaml

clean: ## Clean k8sensus deployment.
	kubectl delete deployment k8sensus
