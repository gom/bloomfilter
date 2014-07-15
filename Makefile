SAMPLE_BIN=filter_sample

default: all
all: sample

test:
	go test ./...

sample: test
	go build -o ${SAMPLE_BIN} sample/main.go

clean:
	rm ${SAMPLE_BIN}
