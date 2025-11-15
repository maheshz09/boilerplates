#aws api gateway

 - basics of api gateway
 - API methods (GET, post, patch, delete & more) there are lot more api options we have on api gateway aws
 - api gateway create 
 - difference between HTTP Apis & rest api
 - proxy & non proxy lambda integration
 - resource policy IAM
 - authorizers in Api gateway
 
 user --> api gateway --> access data/business logic/backend services(ec2,lambda others)
api gateway helps in create publish maintain monitor and secure various endpoints

How does it work?

API Gateway acts as a reverse proxy, accepting API calls from clients and routing them to the appropriate backend service Lambda, EC2, HTTP endpoint, or other AWS services.

what we can create with api gateway
	-	 restful Apis
	-   WebSocket Apis
	-   containerized workflows
	-   serverless workloads

Why use it?

- Scalability: Automatically scales to handle incoming traffic.
- Security: Offers throttling, authorization, and access control.
- Efficiency: Supports multiple API versions for iterative development.
- Monitoring: Integrated with Amazon CloudWatch.

API stands for Application Programming Interface. Think of it as a special messenger that lets different software programs talk to each other and share information.

supported options:
![[Pasted image 20250914115644.png]]

GET request: (purpose fetch info)
![[Pasted image 20250914115617.png]]
red one is request and green one response

POST request (purpose save)
![[Pasted image 20250914115825.png]]
to save information to the backend system at that case we use post api( the post request contains metadata). if we receive a request response of 200 that means the information is stored on backend system.


DELETE: if we want to remove information from backend system
![[Pasted image 20250914120323.png]]

PATCH to update the information which is already existed in the backend 
![[Pasted image 20250914120421.png]]

Implementation: Creating api gateway
![[Pasted image 20250914120727.png]]

so the lambda function will run the business logic depends on our need and thru api gateway the lambda function at as backend for us

as you can see response variable in lambda function which it is going to throw back. when we will see get request to and api endpoint 
the response code it will throw is 200 and content-type: application/json

```python
import json

def lambda_handler(event, context):
    name = "World"
    if event.get('queryStringParameters'):
        name = event.get('queryStringParameters').get('name', name)

    response_data = {
        "message": f"Hello, {name}! Your request was successful.",
        "status": "success"
    }

    return {
        "statusCode": 200,
        "headers": {"Content-Type": "application/json"},
        "body": json.dumps(response_data)
    }

```

**on aws console:**
1. Go on lambda --> create function --> function name (first-gateway-lambda)
2. runtime as python 3.10
3. create function
once the function is created we can modify the code

4. as you can see in code section the pseudocode which they provided 
5. update the code above we can use that because it will also return the json payload and response code (we can also use the default code which they have  provided) --> click on deploy
6. so on the test tab we can validate our lambda functions is it working as expected (in event json we don't need to pass anything because we are performing GET request also our program does not require that)
![[Pasted image 20250914123123.png]]
as you can see the test we ran it got success and it send response code as 200

so lets move to api gateway to create api gateway: 

1. go on api gateway & click on create api
2. create HTTP API
![[Pasted image 20250914123439.png]]
3. Name of the API 'api-gateway-with-get'
4. we don't need to create routes for now move next
5. we don't need to create stages for now move next
6. next --> create

so we have created api gateway but its not yet integrated with our lambda function

to start the integrating lambda withing gateway we need to create routes on the routes we define the GET request/api.
	1. select route as GET on / and click on create
	2. click on atttach integration to inegrate the lambda 
	3. fill all the integration details like region, lambda function
	4. click on create
we have done with the routes part  we need URL to access this route/api
![[Pasted image 20250918111938.png]]

go back to the api you will see "Invoke URL" copy it paste on browser your hiting api

![[Pasted image 20250918111916.png]]

	imp: so when sending get reqeust its easy to send from the browser but when we need to send a post request to api with payload in it for that we need postman or burp

something like this 
![[Pasted image 20250918112332.png]]

response:
![[Pasted image 20250918112402.png]]


![[Pasted image 20250918112539.png]]


PART 2:

REST API: proxy & non-proxy integration
![[Pasted image 20250918113241.png]]

non-proxy integration: so we are creating a api gateway for called post and in the payload we will send it contain a "name" = "value" once the payload is send we can use the power of request modification (which is only possible with non-proxy integration) 
so in that request we gonna add a prefix basically we going to modify the value of the key "prefix_rahul" 
so this non-proxy integration which will allow you to modify the request which is comming & before it gets processed with lambda
	it can able to modify the request as well as response

	1. go to lambda click on create a function
	2. name as non-proxy-lambda-processor runtime as python 3.10
	3. once its created go on code section and edit it 
```python
		import json

		def lambda_handler(event, context):
	    name = event['queryStringParameters'].get('name', 'World')
	    return {
	        "statusCode": 200,
	        "headers": {"Content-Type": "application/json"},
	        "body": json.dumps({"message": f"Hello, {name}!"})
		    }
	``` 

	4. go to the test and create event and test it, and for the payload as
]![[Pasted image 20250920151939.png]]

5 . click on save and test

![[Pasted image 20250920152348.png]]


lets create api gateway for non-proxy integration
1. go on api gateway
2. choose REST API
3. use name "non-proxy-api-gateway" after that click on create api
4. click to create resource at / path and resource name is "non-proxy-demo-api"
5. click on create resource
6. after creating a resource we need to create a method (that method enables us to power of modification of integration request)
7. choose method type POST
8. integration type is lambda function, and choose the correct lambda function and click on create lambda function 
	1. go on method request, go on request validator and select validate body,  string parameter (its just another validator for our request)
	![[Pasted image 20250920153153.png]]

