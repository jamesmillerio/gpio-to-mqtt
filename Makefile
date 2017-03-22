build:
	go build
	go install

run:
	./gpio-to-mqtt

refresh:
	git pull

buildlatest: refresh build

runlatest: buildlatest run
