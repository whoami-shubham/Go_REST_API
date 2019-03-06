INSTRUCTIONS

This REST API contains following four routes:
  localhost:4300/
  localhost:4300/users
  localhost:4300/users/userid
  localhost:4300/create/user


To run this REST API run following command
         go run main.go

To create new user use following command or you can use a tool called postman
         curl -d "id=someid&name=username&date_of_birth=dd-mm-yyyy" localhost:4300/create/user

 
