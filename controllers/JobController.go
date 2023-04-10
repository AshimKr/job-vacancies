package controllers

import (
	"example/challenge/initializers"
	"example/challenge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @route: http://localhost:3000/job
// @Method: GET
func GetAllJob(c *gin.Context) {
	// Find jobs by Country

	// Look up requested jobs
	var allJob []models.AllJob

	result := initializers.DB.Find(&allJob)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, allJob)
}

// @route: http://localhost:3000/job
// @Method: POST
// @description: To create a single job in db
func JobCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		JobTitle           string
		SeniorityLevel     string
		SeniorityLevelRank int //1-junior, 2-middle, 3-senior
		Country            string
		City               string
		Salary             int
		Currency           string
		RequiredSkills     string
		CompanySize        string
		CompanyDomain      string
	}

	c.Bind(&body)

	// Create a job
	job := models.Job{JobTitle: body.JobTitle, SeniorityLevel: body.SeniorityLevel, SeniorityLevelRank: body.SeniorityLevelRank, Country: body.Country, City: body.City, Salary: body.Salary, Currency: body.Currency, RequiredSkills: body.RequiredSkills, CompanySize: body.CompanySize, CompanyDomain: body.CompanyDomain}

	result := initializers.DB.Create(&job)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"Job": job,
	})
}

// @route: http://localhost:3000/alljob
// @Method: POST
// @description: to create multiple at a time job in db
func AllJobCreate(c *gin.Context) {
	// Get data off req body
	// var body struct {
	// 	JobTitle           string
	// 	SeniorityLevel     string
	// 	SeniorityLevelRank int //1-junior, 2-middle, 3-senior
	// 	Country            string
	// 	City               string
	// 	Salary             int
	// 	Currency           string
	// 	RequiredSkills     string
	// 	CompanySize        string
	// 	CompanyDomain      string
	// }

	// c.Bind(&body)

	//
	// Junior PHP Developer	Junior	DE	Berlin	517500	SVU	PHP, LAMP, HTML, CSS, SQL	100-500	Automotive
	// Senior Java Developer	Senior	DE	Hamburg	897000	SVU	Java, Spring, REST, Microservices, Kafka, Hibernate	50-100	Real Estate
	// Senior Fullstack Developer	Senior	DE	Berlin	839500	SVU	PHP, JavaScript, CSS/SASS, PHPUnit, Karma, Jenkins	50-100	FinTech
	// Senior Golang Developer	Senior	DE	Hamburg	931500	SVU	GoLang, Microservices, GoRPC, Kafka, Jenkins, Docker	50-100	Real Estate
	// Middle Fullstack Developer	Middle	DE	Berlin	724500	SVU	PHP, JavaScript, CSS/SASS, Angular	50-100	FinTech
	// Fullstack TeamLead	Tech management	DE	Berlin	920000	SVU	PHP, JavaScript, SASS/LESS, SCRUM, ReactJS, Redux, NPM, Yarn	50-100	FinTech
	// Middle JavaScript developer	Middle	DE	Berlin	621000	SVU	JavaScript, TypeScript, SASS, React	10-50	Logistics
	// Senior Fullstack Developer	Senior	ES	Barcelona	506000	SVU	Node.js, JavaScript, CSS/SASS, PHPUnit, Karma, Jenkins	1000-5000	Mining
	// Middle Fullstack Developer	Middle	ES	Barcelona	598000	SVU	Node.js, JavaScript, CSS/SASS, Angular	1000-5000	Mining
	// Senior Java Developer	Senior	ES	Barcelona	644000	SVU	Java, JavaEE, jBoss, REST, Microservices, RabbitMQ, PostgresQL	1000-5000	Mining
	// Middle PHP Developer	Middle	PT	Lisbon	425500	SVU	PHP, Falcon, Unit-testing, SOLID, SQL, MongoDB	10-50	Communication
	// Junior PHP Developer	Junior	PT	Lisbon	322000	SVU	PHP, LAMP, SQL, bash, OOP, HTML, CSS	10-50	Communication

	// Create jobs
	allJob := []*models.AllJob{
		{JobTitle: "Senior Golang Developer", SeniorityLevel: "Senior", SeniorityLevelRank: 3, Country: "RE", City: "Bruges", Salary: 1035000, Currency: "SVU", RequiredSkills: "GoLang, Microservices, GoRPC, Bamboo, Docker", CompanySize: "50-100", CompanyDomain: "FinTech"},
		{JobTitle: "Node.js TeamLead", SeniorityLevel: "Tech management", SeniorityLevelRank: 3, Country: "NL", City: "Rotterdam", Salary: 897000, Currency: "SVU", RequiredSkills: "Node.js, JavaScript, JIRA, Bamboo, webpack", CompanySize: "10-50", CompanyDomain: "Real Estate"},
		{JobTitle: "Frontend TeamLead", SeniorityLevel: "Tech management", SeniorityLevelRank: 3, Country: "NL", City: "Rotterdam", Salary: 977500, Currency: "SVU", RequiredSkills: "TypeScript, JavaScript, SASS, JIRA, Bamboo", CompanySize: "10-50", CompanyDomain: "Real Estate"},
		{JobTitle: "PHP TeamLead", SeniorityLevel: "Tech management", SeniorityLevelRank: 3, Country: "NL", City: "Amsterdam", Salary: 920000, Currency: "SVU", RequiredSkills: "PHP, LAML, SCRUM, JIRA, Bamboo, MVC", CompanySize: "10-50", CompanyDomain: "Health"},
		{JobTitle: "Node.js Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 3, Country: "NL", City: "Amsterdam", Salary: 690000, Currency: "SVU", RequiredSkills: "Node.js, JavaScript, Git, Mongo, noSQL, NPM", CompanySize: "10-50", CompanyDomain: "Health"},
		{JobTitle: "Frontend Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 3, Country: "NL", City: "Amsterdam", Salary: 747500, Currency: "SVU", RequiredSkills: "TypeScript, JavaScript, SASS, LESS, Karma", CompanySize: "10-50", CompanyDomain: "Health"},
		{JobTitle: "PHP Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 3, Country: "NL", City: "Amsterdam", Salary: 782000, Currency: "SVU", RequiredSkills: "PHP, LAMP, MySQL, PHPUnit, OOP", CompanySize: "10-50", CompanyDomain: "Automotive"},

		{JobTitle: "Senior PHP Developer", SeniorityLevel: "Senior", SeniorityLevelRank: 3, Country: "PT", City: "Lisbon", Salary: 494500, Currency: "SVU", RequiredSkills: "PHP, Falcon, REST, SQL, MongoDB, Unit-testing, Behat, SOLID, Docker, AWS", CompanySize: "10-50", CompanyDomain: "Communication"},
	}

	result := initializers.DB.Create(&allJob)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"Job": allJob,
	})
}

// @route: http://localhost:3000/jobbycountryorcity
// @Method: POST
// @description: To find a job by using country name or by city name
func JobByCountryORCity(c *gin.Context) {
	// Find jobs by Country
	var body struct {
		Country string
		City    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up requested jobs
	var allJob []models.AllJob

	result := initializers.DB.Where("country = ?", body.Country).Or("city = ?", body.City).Find(&allJob)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, allJob)
}

// @route: http://localhost:3000/jobbycountry
// @Method: POST
// @description: To find a job by using country name and seniority level
func JobByCountry(c *gin.Context) {
	// Find jobs by Country
	var body struct {
		Country        string
		SeniorityLevel string
		salary         int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up requested jobs
	var allJob []models.AllJob

	result := initializers.DB.Where("country = ? AND seniority_level = ?", body.Country, body.SeniorityLevel).Find(&allJob)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, allJob)
}
