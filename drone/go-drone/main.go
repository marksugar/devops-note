package main

import (
	"fmt"
	"log"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

const (
	token = "06fpxwl9PCUArz2DUQNhKuAtNHBT50hk"
	host  = "http://172.25.200.22:8000"
)

func main() {
	// create an http client with oauth authentication.
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	// 创建一个新的 Drone 客户端
	client := drone.NewClient(host, auther)

	// 获取仓库的构建信息
	owner := "test"
	repo := "backend"
	// builds, err := client.BuildList(owner, repo, drone.ListOptions{})
	// if err != nil {
	// 	log.Fatalf("Failed to list builds: %v", err)
	// }

	// 打印构建信息
	// for _, build := range builds {
	// 	fmt.Printf("Build #%d - Status: %s\n", build.Number, build.Status)
	// 	fmt.Println(build)
	// }

	// dr, err := client.BuildLast(owner, repo, "master")
	// if err != nil {
	// 	log.Fatalf("Failed to list builds: %v", err)
	// }
	// fmt.Println(dr)
	// params := map[string]string{
	// 	"password": "ZTQwMTg2ZWViODA1NTMwYmYx",
	// }

	// bu, err := client.BuildCreate(owner, repo, "", "master", params)
	// if err != nil {
	// 	log.Fatalf("Failed to list builds: %v", err)
	// }
	// fmt.Println(bu)

	line, err := client.Logs(owner, repo, 10, 1, 3)
	if err != nil {
		log.Fatalf("Failed to list line: %v", err)
	}

	for _, v := range line {
		fmt.Println(v.Message)
	}
}
