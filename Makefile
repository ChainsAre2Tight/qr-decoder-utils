run-image:
	go run cmd/main.go --input ./data/animation.jpg --output ./data/result.png --mode image
run-excel:
	go run cmd/main.go --input ./data/animation.jpg --output ./data/result --mode excel