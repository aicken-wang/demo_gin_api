APPNAME=gin_api
TARGET=${APPNAME}
.PHONY: clean debug install

${TARGET}:
	go mod tidy && go build -o ${APPNAME} main.go
	# && strip ${APPNAME}
debug: 
	go mod tidy && go build -o ${APPNAME} main.go
clean: 
	rm -rf ${APPNAME}