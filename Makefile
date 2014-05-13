.PHONY: test live-test live-testv test testv fmt

live-test:
	TWILIO_TEST_CONFIG=c:/cygwin/home/kyle/.twilio.json go test

live-testv:
	TWILIO_TEST_CONFIG=c:/cygwin/home/kyle/.twilio.json go test -test.v

test:
	go test

testv:
	go test -test.v

fmt:
	go fmt .
