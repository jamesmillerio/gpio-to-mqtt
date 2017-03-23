build:
	go build
	go install
	cp ./gpio-to-mqtt /usr/bin/gpio-to-mqtt

run:
	./gpio-to-mqtt

refresh:
	git pull

buildlatest: refresh build

runlatest: buildlatest run
