build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/nithinchandra/shipping/consignment-service \
    	  proto/consignment/consignment.proto
	docker build -t consignment-service .



run:
	docker run -p 50051:50051 \
	 -e MICRO_SERVICE_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns	consignment-service
