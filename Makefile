default: build

dep-init:
	go get github.com/kardianos/govendor
	govendor init

dep-update:
	govendor remove +unused
	govendor add +external

fmt:
	go fmt $(go list ./... | grep -v /vendor/)

build:
	docker build -t zyfdedh/iploc .

run:
	docker run -p 80:80 zyfdedh/iploc

push: build
	docker push zyfdedh/iploc
