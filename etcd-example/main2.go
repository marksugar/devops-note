package main

import (
	"etcd/etcdtask"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := etcdtask.NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 启动监视任务变化的协程
	go client.WatchTasks()

	r := gin.Default()

	r.POST("/tasks", func(c *gin.Context) {
		var task etcdtask.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := client.AddTask(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		fmt.Println(c.Params)
		// id, err := strconv.Atoi(c.Param("id"))
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		// 	return
		// }
		fmt.Println(c.Param("id"))
		task, err := client.GetTask(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task etcdtask.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// task.ID, _ = strconv.Atoi(c.Param("id"))
		task.ID = c.Param("id")
		if err := client.UpdateTask(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}
		if err := client.DeleteTask(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Task deleted"})
	})

	r.GET("/tasks/list", func(c *gin.Context) {
		tasks, err := client.ListTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})

	r.Run(":8080")
}
