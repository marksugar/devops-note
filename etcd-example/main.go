package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdClient *clientv3.Client

func main() {
	var err error
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.25.110.31:2379"}, // etcd 服务器地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer etcdClient.Close()

	r := gin.Default()

	r.POST("/create", createValue)
	r.GET("/get/:key", getValue)
	r.PUT("/update", updateValue)
	r.DELETE("/delete/:key", deleteValue)

	r.Run(":8081")
}

func createValue(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := etcdClient.Put(ctx, req.Key, req.Value)
	cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "value created"})
}

func getValue(c *gin.Context) {
	key := c.Param("key")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(resp.Kvs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": string(resp.Kvs[0].Key), "value": string(resp.Kvs[0].Value)})
}

func updateValue(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := etcdClient.Put(ctx, req.Key, req.Value)
	cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "value updated"})
}

func deleteValue(c *gin.Context) {
	key := c.Param("key")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := etcdClient.Delete(ctx, key)
	cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "key deleted"})
}
