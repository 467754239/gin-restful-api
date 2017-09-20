package apis

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/467754239/db-api/models"

	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	p := Person{}
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": persons,
	})

}

func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatal(err)
	}
	p := Person{Id: id}
	ra, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": ra,
	})

}

func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatal(err)
	}
	p := Person{Id: id}
	ra, err := p.DelPerson()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": ra,
	})
}

func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{Id: id, FirstName: firstName, LastName: lastName}
	//fmt.Println(c.Request.PostForm)
	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
	})
}
