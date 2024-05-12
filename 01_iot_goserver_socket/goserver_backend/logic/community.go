package logic

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/dao/redis"
)

// GetCommunityList 获取社区列表
func GetCommunityList(c *gin.Context) ([]mysql.Community, error) {
	defer redis.Close()
	id, err := strconv.Atoi(redis.Rdb.Get("id").Val())
	if err == nil {
		community := mysql.Community{
			Id:            uint32(id),
			CommunityId:   redis.Rdb.Get("community_id").Val(),
			CommunityName: redis.Rdb.Get("community_name").Val(),
			Introduction:  redis.Rdb.Get("introduction").Val(),
		}
		communities := []mysql.Community{community}
		return communities, nil
	}

	communities, err := mysql.GetCommunityList()

	if err != nil {
		return nil, err
	}

	return communities, nil

}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(c *gin.Context, id int64) (*mysql.Community, error) {
	community, err := mysql.GetCommunityDetailById(id)
	if err != nil {
		return nil, err
	}

	return community, nil
}
