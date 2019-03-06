**INSTRUCTIONS**

This REST API contains following four routes: <br/>
  localhost:4300/  <br/>
  localhost:4300/users   <br/>
  localhost:4300/users/userid   <br/>
  localhost:4300/create/user     <br/>


To run this REST API run following command  <br/>
go run main.go                     <br/>

To create new user use following command or you can use a tool called postman <br/>
curl -d "id=someid&name=username&date_of_birth=dd-mm-yyyy" localhost:4300/create/user <br/>

 
