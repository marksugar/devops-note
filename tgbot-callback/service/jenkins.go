package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bndr/gojenkins"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var jenkins *gojenkins.Jenkins

type jenkinsService struct {
}

var JenkinsService = new(jenkinsService)

// ctx context.Context,
func (j *jenkinsService) initJenkinsClient(jenkinsURL, username, password string) (err error) {
	ctx := context.Background()
	jenkins = gojenkins.CreateJenkins(nil, jenkinsURL, username, password)
	_, err = jenkins.Init(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize Jenkins client: %w", err)
	}
	return nil
}
func (j *jenkinsService) CheckProjectExists(ctx context.Context, jobname string) bool {
	job, err := jenkins.GetJob(ctx, jobname)
	// job, err := j.GetJob(ctx, joname)
	fmt.Println("job:", job, err)
	if err != nil {
		return true
	}
	return false
}
func (j *jenkinsService) triggerJenkinsBuild(ctx context.Context, jobName string, params map[string]string) (int64, error) {
	queueID, err := jenkins.BuildJob(ctx, jobName, params)
	if err != nil {
		return 0, fmt.Errorf("failed to trigger Jenkins build: %w", err)
	}
	return queueID, nil
}
func (j *jenkinsService) waitForBuild(ctx context.Context, queueID int64) (*gojenkins.Build, error) {
	for {
		build, err := jenkins.GetBuildFromQueueID(ctx, queueID)
		if err == nil && build != nil && build.IsRunning(ctx) {
			log.Printf("Build %d is running...", build.GetBuildNumber())
			zap.L().Info(fmt.Sprintf("Build %d is running...", build.GetBuildNumber()))
			time.Sleep(10 * time.Second)
		} else if build != nil && !build.IsRunning(ctx) {
			return build, nil
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}
func (j *jenkinsService) printBuildResult(ctx context.Context, build *gojenkins.Build) {
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

// "测试", firsttext, twotext, manytext
// /dev env type arys1 arys1
func (j *jenkinsService) maintrigger(itmes, firsttext, twotext, manytext string) (result string, buildnum int64, buildtime string, err error) {
	ctx := context.Background()
	err = j.initJenkinsClient(viper.GetString("jenkins.url"), viper.GetString("jenkins.user"), viper.GetString("jenkins.password"))
	if err != nil {
		zap.L().Fatal("init jenkins failed,err: ", zap.Error(err))
		return "", 0, "", errors.New(fmt.Sprintf("init jenkins failed,err:%v\n", err))
	}

	resut := j.CheckProjectExists(ctx, itmes)
	if resut {
		zap.L().Debug(fmt.Sprintf("%s项目不存在", itmes))
		return "", 0, "", errors.New(fmt.Sprintf("%s项目不存在", itmes))
	}

	queueid, err := j.triggerJenkinsBuild(ctx, itmes, map[string]string{"Code_Type": twotext, "WEB_ID": manytext, "Platform": firsttext, "GIT_TAG": ""})
	if err != nil {
		zap.L().Debug(fmt.Sprintf("构建失败:%s", err))
		return "", 0, "", errors.New(fmt.Sprintf("构建失败:%s", err))
	}

	build, err := j.waitForBuild(ctx, queueid)
	if err != nil {
		panic(err)
	}

	if build.GetResult() != "SUCCESS" {
		return "", 0, "", errors.New("更新失败")
	}

	// 获取构建的控制台输出
	// output, err := getConsoleOutput(ctx, build)
	// if err != nil {
	// 	log.Fatalf("Error getting console output: %s", err)
	// }
	// fmt.Println(output)

	// fmt.Println("更新成功,id: %d", build.GetBuildNumber())
	// fmt.Printf("用时%f秒钟", build.Raw.Duration/1000)

	return build.GetResult(), build.GetBuildNumber(), fmt.Sprintf("用时%f秒", build.Raw.Duration/1000), nil
}

func (j *jenkinsService) timeConversion(tis string) (data string, err error) {
	gtimestamp, err := strconv.ParseInt(tis, 10, 64)
	if err != nil {
		return "", err
	}
	stime := time.Unix(gtimestamp, 0).Format("15:04:05")
	return stime, nil
}
func (j *jenkinsService) getConsoleOutput(ctx context.Context, build *gojenkins.Build) (string, error) {
	return build.GetConsoleOutput(ctx), nil
}
