package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.GET("/members", listMembersHandler)
	r.POST("/members", createMembersHandler)
	r.DELETE("/members/:id", deleteMembersHandler)

	r.Run()

}

type Member struct {
	ID          string `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

var members = []Member{
	{ID: "001", Firstname: "Jetsupa", Lastname: "Kruetang", Age: 28, Nationality: "Thailand"},
	{ID: "002", Firstname: "Cherprang", Lastname: "Areekul", Age: 26, Nationality: "Thailand"},
	{ID: "003", Firstname: "Kana", Lastname: "Hanasawa", Age: 31, Nationality: "Japan"},
	{ID: "004", Firstname: "Min-young", Lastname: "Prak", Age: 35, Nationality: "Korea"},
	{ID: "005", Firstname: "Diana", Lastname: "Chongjintanakarn", Age: 40, Nationality: "Thai-China"},
}

func listMembersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, members)
}

func createMembersHandler(c *gin.Context) {
	var member Member

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	members = append(members, member)

	c.JSON(http.StatusCreated, member)
}

func deleteMembersHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range members {
		if a.ID == id {
			members = append(members[:i], members[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
