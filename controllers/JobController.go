package controllers

import (
	"example/challenge/initializers"
	"example/challenge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// Trying to create bunch records
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
	// Senior PHP Developer	Senior	PT	Lisbon	494500	SVU	PHP, Falcon, REST, SQL, MongoDB, Unit-testing, Behat, SOLID, Docker, AWS	10-50	Communication
	// PHP Developer	Middle	NL	Amsterdam	782000	SVU	PHP, LAMP, MySQL, PHPUnit, OOP	10-50	Automotive
	// Frontend Developer	Middle	NL	Amsterdam	747500	SVU	TypeScript, JavaScript, SASS, LESS, Karma	10-50	Health
	// Node.js Developer	Middle	NL	Amsterdam	690000	SVU	Node.js, JavaScript, Git, Mongo, noSQL, NPM	10-50	Health
	// PHP TeamLead	Tech management	NL	Amsterdam	920000	SVU	PHP, LAML, SCRUM, JIRA, Bamboo, MVC	10-50	Health
	// Frontend TeamLead	Tech management	NL	Rotterdam	977500	SVU	TypeScript, JavaScript, SASS, JIRA, Bamboo	10-50	Real Estate
	// Node.js TeamLead	Tech management	NL	Rotterdam	897000	SVU	Node.js, JavaScript, JIRA, Bamboo, webpack	10-50	Real Estate

	// Senior Golang Developer	Senior	BE	Bruges	1035000	SVU	GoLang, Microservices, GoRPC, Bamboo, Docker	50-100	FinTech

	// Create jobs
	allJob := []*models.AllJob{
		{JobTitle: "Fullstack TeamLead Tech management", SeniorityLevel: "Senior", SeniorityLevelRank: 3, Country: "IE", City: "Dublin", Salary: 897000, Currency: "SVU", RequiredSkills: "SVU	PHP, JavaScript, CSS/SASS, JIRA, SCRUM, AWS, SNS/SQS, Kinesis, NPM", CompanySize: "10-50", CompanyDomain: "Logistics"},
		{JobTitle: "Middle Fullstack Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 2, Country: "IE", City: "Dublin", Salary: 862500, Currency: "SVU", RequiredSkills: "PHP, JavaScript, CSS/SASS, SQL, AWS, Docker", CompanySize: "10-50", CompanyDomain: "Logistics"},
		{JobTitle: "Senior Java Developer", SeniorityLevel: "Senior", SeniorityLevelRank: 3, Country: "BE", City: "Bruges", Salary: 977500, Currency: "SVU", RequiredSkills: "Java, J2SE, Spring, Microservices, Bamboo, Docker", CompanySize: "10-50", CompanyDomain: "FinTech"},
		{JobTitle: "Middle Java Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 3, Country: "BE", City: "Antwerp", Salary: 920000, Currency: "SVU", RequiredSkills: "Java, Spring, Microservices, Kinesis, Junit, SOAP/RPC, Hibernate", CompanySize: "50-100", CompanyDomain: "Automotive"},
		{JobTitle: "Middle Golang Developer", SeniorityLevel: "Middle", SeniorityLevelRank: 3, Country: "BE", City: "Antwerp", Salary: 770500, Currency: "SVU", RequiredSkills: "GoLang, Microservices, GoRPC", CompanySize: "50-100", CompanyDomain: "Automotive"},
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

// @desc: Find job by country name OR city name
// Protected route
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

// @desc: Find job by country name and seniority level
// Protected route
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
