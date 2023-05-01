ACCOUNT_NAME := autosetup
PROJECT_NAME := golang-rabbitmq-examples
TAG ?= 0.0.9

build::
	docker build -t $(PROJECT_NAME):$(TAG) .
	docker image tag $(PROJECT_NAME):$(TAG) $(ACCOUNT_NAME)/$(PROJECT_NAME):$(TAG)

install::
	docker push $(ACCOUNT_NAME)/$(PROJECT_NAME):$(TAG)
