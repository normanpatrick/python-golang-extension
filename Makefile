rebuild: clean build

build:
	@echo "  + Building ..."
	@go build -buildmode=c-shared -o experiment.so

clean:
	@echo "  + Cleaning..."
	@rm -f experiment.so experiment.h

test: build
	@echo "Dumb test.."
	@python3 -c \
	'import experiment; print(experiment.experiment2("hoo",["sure", "thing"]))'
