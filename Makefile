PACKAGE_NAME=k8s-homebase

BUILD_DIR=build/
BUILD_FILE=$(BUILD_DIR)$(PACKAGE_NAME)

make: compile run

compile:
	go build -o $(BUILD_FILE)

clean:
	rm -f $(BUILD_FILE)

run:
	$(BUILD_FILE)
