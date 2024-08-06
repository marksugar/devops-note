package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"etcd/active"
	"etcd/models"
	vts "etcd/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var Api = new(SystemCron)

type SystemCron struct {
	ac active.SystemCron
}

func (s SystemCron) Tables(c *gin.Context) {
	p := new(models.Tablist)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("TableHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			vts.ResponseError(c, vts.CodeInvalidParm)
			return
		}
		vts.ResponseErrorWitMsg(c, vts.CodeInvalidParm, vts.RemoveTopStruct(errs.Translate(vts.Trans)))
		return
	}

	data, err := s.ac.Tables(p.Page, p.PageSize, p.Name)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("SystemCron Cron.Tables select failed form", zap.Error(err))
		vts.ResponseError(c, vts.CodeSqlProblem)
		return
	}
	vts.ResponseSuccess(c, data)
}
func (s SystemCron) Updates(c *gin.Context) {
	var tasks models.Tasks
	tasks.ID = c.Param("id")
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.ac.Updates(&tasks)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("SystemCron Cron Tasks Updates failed form", zap.Error(err))
		vts.ResponseError(c, vts.CodeSqlProblem)
		return
	}
	vts.ResponseSuccess(c, struct {
		Msg string `json:"msg"`
	}{
		Msg: "更新成功",
	})
}
func (s SystemCron) Updates2(c *gin.Context) {
	var tasks models.Tasks
	tasks.ID = c.Param("id")
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.ac.UpdateIDs2(tasks, uuid.New())
	if err != nil {
		fmt.Println(err)
		zap.L().Error("SystemCron Cron Tasks Updates failed form", zap.Error(err))
		vts.ResponseError(c, vts.CodeSqlProblem)
		return
	}
	vts.ResponseSuccess(c, struct {
		Msg string `json:"msg"`
	}{
		Msg: "更新成功",
	})
}
func (s SystemCron) Adds(c *gin.Context) {
	var p models.Tasks
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SystemCron Adds Handler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			vts.ResponseError(c, vts.CodeInvalidParm)
			return
		}
		vts.ResponseErrorWitMsg(c, vts.CodeInvalidParm, vts.RemoveTopStruct(errs.Translate(vts.Trans)))
		return
	}
	iPasswd := strings.TrimSpace(p.Name)
	if iPasswd == "" {
		vts.ResponseError(c, vts.CodeCannotBeEmpty)
		return
	}
	if err := s.ac.Adds(&p); err != nil {
		zap.L().Error("SystemCron Add HandlerPage with invalid param from")
		vts.ResponseErrorActiveMsg(c, vts.CodeServerBusy, fmt.Sprintf("%s", err))
		return
	}
	vts.ResponseSuccess(c, struct {
		Msg string `json:"msg"`
	}{
		Msg: "添加成功",
	})
}
func (s SystemCron) Dels(c *gin.Context) {
	ids := struct {
		ID string `json:"id"`
	}{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(&ids); err != nil {
		zap.L().Error("SystemCron Dels Handler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			vts.ResponseError(c, vts.CodeInvalidParm)
			return
		}
		vts.ResponseErrorWitMsg(c, vts.CodeInvalidParm, vts.RemoveTopStruct(errs.Translate(vts.Trans)))
		return
	}
	if res := s.ac.Dels(ids.ID); res != nil {
		zap.L().Error("SystemCron Dels with invalid param form")
		vts.ResponseErrorActiveMsg(c, vts.CodeServerBusy, fmt.Sprintf("%s", res))
		return
	}
	vts.ResponseSuccess(c, struct {
		Msg string `json:"msg"`
	}{
		Msg: "删除成功",
	})
}
