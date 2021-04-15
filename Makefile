SVC_AUT ?= svc_aut
SVC_DRIVER ?= svc_driver

# Start application
start-solution:
	cd svc_aut && docker build . -t ${SVC_AUT}
	cd svc_driver && docker build . -t ${SVC_DRIVER}
	cd docker  && docker-compose up -d;

# Push the docker image
docker-push:
	docker push ${SVC_AUTH}
	docker push ${SVC_DRIVER}

# Push the docker image
delete-solution:
	##cd docker && docker-compose down && docker volume prune -f && docker rmi -f $(docker images -a -q);

# Run tests
run-tests-svc-driver:
	cd svc_driver && go test ./... -v

