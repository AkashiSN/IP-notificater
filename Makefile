API     := `cat api`
LDFLAGS := -ldflags=" -w -X \"main.API=$(API)\""
GCFLAGS := -gcflags="-trimpath=${GOPATH}"
USER    := ${USER}

build:
	go build ${GCFLAGS} ${LDFLAGS}

install:
	sudo mkdir -p /opt/IP-notificater
	sudo chown -R ${USER}:${USER} /opt/IP-notificater
	cp IP-notificater /opt/IP-notificater/IP-notificater
	crontab cron.conf
