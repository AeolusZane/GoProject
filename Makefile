register:
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/register -H "Content-Type: application/json" -d '{"username":"aeolusj","password":"iloveu"}'
testurl:
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/login -H "Content-Type: application/json" -d '{"username":"aeolusj","password":"iloveu"}'
	@echo ""
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/login -H "Content-Type: application/json" -d '{"username":"aeolusj","password":"iloved"}'
	@echo ""