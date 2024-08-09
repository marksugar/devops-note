package main

import (
	"context"
	"fmt"
	"gjenkins/settings"
	"log"
	"strconv"
	"time"

	"github.com/bndr/gojenkins"
	"github.com/spf13/viper"
)

var jenkins *gojenkins.Jenkins

// ctx context.Context,
func initJenkinsClient(jenkinsURL, username, password string) (err error) {
	ctx := context.Background()
	jenkins = gojenkins.CreateJenkins(nil, jenkinsURL, username, password)
	_, err = jenkins.Init(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize Jenkins client: %w", err)
	}
	return nil
}
func CheckProjectExists(ctx context.Context, jobname string) bool {
	job, err := jenkins.GetJob(ctx, jobname)
	// job, err := j.GetJob(ctx, joname)
	fmt.Println("job:", job, err)
	if err != nil {
		return true
	}
	return false
}
func triggerJenkinsBuild(ctx context.Context, jobName string, params map[string]string) (int64, error) {
	queueID, err := jenkins.BuildJob(ctx, jobName, params)
	if err != nil {
		return 0, fmt.Errorf("failed to trigger Jenkins build: %w", err)
	}
	return queueID, nil
}
func waitForBuild(ctx context.Context, queueID int64) (*gojenkins.Build, error) {
	for {
		build, err := jenkins.GetBuildFromQueueID(ctx, queueID)
		if err == nil && build != nil && build.IsRunning(ctx) {
			log.Printf("Build %d is running...", build.GetBuildNumber())
			time.Sleep(10 * time.Second)
		} else if build != nil && !build.IsRunning(ctx) {
			return build, nil
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}
func printBuildResult(ctx context.Context, build *gojenkins.Build) {
	for build.IsRunning(ctx) {
		// 10秒获取一次状态，并且打印ConsoleText
		time.Sleep(10 * time.Millisecond)
		build.GetConsoleOutput(ctx)
		build.Poll(ctx)
	}
	if build != nil {
		log.Printf("Build %d result: %s", build.GetBuildNumber(), build.GetResult())
		log.Printf("Build %d URL: %s", build.GetBuildNumber(), build.GetUrl())
	} else {
		log.Println("No build information found.")
	}
}
func main() {
	ctx := context.Background()
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}
	err := initJenkinsClient(viper.GetString("jenkins.url"), viper.GetString("jenkins.user"), viper.GetString("jenkins.password"))
	if err != nil {
		fmt.Printf("init jenkins failed,err:%v\n", err)
		return
	}

	resut := CheckProjectExists(ctx, "测试2")
	if resut {
		fmt.Println("没有")
		return
	}
	fmt.Println("有")
	queueid, err := triggerJenkinsBuild(ctx, "测试2", map[string]string{"WEB_ID": "6000", "Code_Type": "vue_pc", "Platform": "Grayscale", "GIT_TAG": "123"})
	if err != nil {
		panic(err)
	}

	build, err := waitForBuild(ctx, queueid)
	if err != nil {
		panic(err)
	}

	if build.GetResult() != "SUCCESS" {
		fmt.Println("更新失败")
		return
	}

	// 获取构建的控制台输出
	output, err := getConsoleOutput(ctx, build)
	if err != nil {
		log.Fatalf("Error getting console output: %s", err)
	}
	fmt.Println(output)
	// printBuildResult(ctx, build)
	fmt.Println("更新成功,id: %d", build.GetBuildNumber())
	fmt.Printf("用时%f秒钟", build.Raw.Duration/1000)
}

func timeConversion(tis string) (data string, err error) {
	gtimestamp, err := strconv.ParseInt(tis, 10, 64)
	if err != nil {
		return "", err
	}
	stime := time.Unix(gtimestamp, 0).Format("15:04:05")
	return stime, nil
}
func getConsoleOutput(ctx context.Context, build *gojenkins.Build) (string, error) {
	return build.GetConsoleOutput(ctx), nil
}
