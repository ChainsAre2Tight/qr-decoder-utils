run-image:
	go run cmd/main.go --input ./data/animation.jpg --output ./data/result --mode image
run-excel:
	go run cmd/main.go --input ./data/animation.jpg --output ./data/result --mode excel
run-decoding:
	go run cmd/main.go --input ./data/animation.jpg --mode value

build:
	go build -o ./bin/qr_decoder ./cmd/main.go
