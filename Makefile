register:
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/register -H "Content-Type: application/json" -d '{"username":"primeaeolusj","password":"iloveu"}'
	@echo ""
loginy:
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/login -H "Content-Type: application/json" -d '{"username":"primeaeolusj","password":"iloveu"}'
	@echo ""
loginn:	
	@curl -X POST https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/login -H "Content-Type: application/json" -d '{"username":"primeaeolusj","password":"iloved"}'
	@echo ""
protected:
	@curl -X GET https://gsy4i6ldmi.execute-api.us-east-1.amazonaws.com/prod/protected -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNzE4MDE5OTE0LCJ1c2VyIjoicHJpbWVhZW9sdXNqIn0._CMgSmUUUCY_635rOO8Got7E4yvA8r5gLqRcPPTuQpM"