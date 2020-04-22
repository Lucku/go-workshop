package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type API struct {
	repository Repository
}

func NewAPI(repository Repository) *API {
	return &API{repository: repository}
}

func (a *API) GetAllTodos(c *gin.Context) {

	todos := a.repository.FindAll()

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (a *API) GetTodoByID(c *gin.Context) {

	idx := c.Param("id")

	if idx == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID parameter provided"})
		return
	}

	idxInt, err := strconv.Atoi(idx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided as a number"})
		return
	}

	todo := a.repository.FindByID(idxInt)

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func (a *API) CreateTodo(c *gin.Context) {

	newTodo := new(Todo)

	if err := c.BindJSON(newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid Todo object provided"})
		return
	}

	fmt.Println(newTodo)

	inserted := a.repository.Save(newTodo)

	fmt.Println(inserted)

	c.JSON(http.StatusOK, gin.H{"inserted": inserted})
}

func (a *API) DeleteTodo(c *gin.Context) {

	idx := c.Param("id")

	if idx == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID parameter provided"})
		return
	}

	idxInt, err := strconv.Atoi(idx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Index not provided as a number"})
		return
	}

	deleted := a.repository.DeleteByID(idxInt)

	if !deleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Object with ID not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "Object was deleted"})
	}
}
