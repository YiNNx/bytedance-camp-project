/*
 * @Author: Go不浪队
 * @Date: 2023-02-06 21:05:56
 * @LastEditTime: 2023-02-07 20:29:36
 * @Description:
 */

package db

import (
	"time"

	"github.com/go-redis/redis/v8"

	"tiktok/pkg/config"
	"tiktok/pkg/log"
)

 var (
	 redisClient *redis.Client
 )
 
 const (
	 redisTimeoutTime = time.Millisecond * 500 // 设置一个较短的timeout时间，可以在redis故障时还能够以较快的速度穿透到数据库获得数据
	 redisExpireTime  = time.Hour
	 redisNoExpire    = 0

 )
 
 func init() {
	 redisClient = redis.NewClient(
		 &redis.Options{
			 Addr:        config.C.Redis.Addr,
			 Password:    config.C.Redis.Password,
			 DB:          config.C.Redis.DB,
			 DialTimeout: redisTimeoutTime,
		 },
	 )
	 err := redisClient.Ping(redisClient.Context()).Err()
	 if err != nil {
		 log.Logger.Debug(err)
		 log.Logger.Error(err)
	 } else {
		 log.Logger.Info("Redis server connected.")
	 }
 
 }
 
 