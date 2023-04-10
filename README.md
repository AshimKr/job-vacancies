Go Lang API
This is a RESTful API built using the Go programming language, with the Gin framework for routing and GORM for interacting with a Postgres database.

1. Clone the repository to your local machine:
git clone https://github.com/Sustiknow-Ashim/job-vacancies

2. Navigate to the project directory:
cd job-vacancies

3. Install the project dependencies:

4. Set up the database:
Create a Postgres database and update the database credentials in the .env file.

5. Start the server:
go run main.go or you can use compiledaemon --command="./dir_name"

The API will now be running at http://localhost:3000.


already created user details(you can check any details by using this credentials):= {
    "email": "ashim@test.com",
    "password": "123456"
}

Endpoints
The API has the following endpoints:

POST /signup
To Create a user to the database
request body: {
    "email": "johndoe@example.com",
    "password" : "12874541"
}


POST /login
To login and generate jwt token
request body: {
    "email": "johndoe@example.com",
    "password" : "12874541"
}


<!-- all the below api is protected by jwt token -->

GET /validate
To get the logged in user details

POST /job
To Create a job vacancy to the database
request body: {
        JobTitle           
		SeniorityLevel     
		SeniorityLevelRank 
		Country            
		City               
		Salary             
		Currency           
		RequiredSkills     
		CompanySize        
		CompanyDomain      
	}


GET /job
TO get all the job vacancies

POST /jobbycountry
Get details of job vacancy using country name and SeniorityLevel
request body: {
    Country        
	SeniorityLevel 
}


POST /jobbycountryorcity
Get details of job vacancy using country name OR SeniorityLevel
request body: {
    Country        
	City 
}
