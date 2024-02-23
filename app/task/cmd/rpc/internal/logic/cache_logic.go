package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// DeleteListCache 删除 list 缓存
func DeleteListCache(r *redis.Redis, uid int64) {
	prefix := "cache:gtodolist:task:list:"
	listKey := fmt.Sprintf("%s%v:*", prefix, uid)
	//keys, _, _ := r.Scan(0, listKey, 0)
	keys, _ := r.Keys(listKey)
	logx.Info("DeleteListCache, keys: ", keys)

	for _, key := range keys {
		_, _ = r.Del(key)
	}
}
