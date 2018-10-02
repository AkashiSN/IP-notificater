API     := `cat api`
LDFLAGS := -ldflags=" -w -X \"main.API=$(API)\""
GCFLAGS := -gcflags="-trimpath=${GOPATH}"

build:
	go build ${GCFLAGS} ${LDFLAGS}