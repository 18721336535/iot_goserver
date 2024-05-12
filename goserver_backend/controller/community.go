package controller

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"go.uber.org/zap"
)

var PingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

// CommunityHandler 获取社区列表
func CommunityHandler(c *gin.Context) {
	PingCounter.Inc()
	data, err := logic.GetCommunityList(c)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(c, id)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err), zap.Int64("id", id))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
