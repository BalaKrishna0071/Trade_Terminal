package redis_util
import(
	"context"
	"time"
	logger "trade/loggers"
	"github.com/redis/go-redis/v9"
	
)

// Creating Global Variable
var (
	ctx = context.Background()
    redisC *redis.Client
	
)

// Connect Func
func Connect(){
	if redisC == nil{
		redisC = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		PoolSize: 100,
		MinIdleConns: 100, 
		DialTimeout:  5 * time.Second,
	   })
    }
	
	_, err := redisC.Ping(ctx).Result()
	if err != nil{
		logger.Logger.Error("Could not Conect to Redis : %v", err)
		
	}
	logger.Logger.Info("Redis Connected... ")

}

// Insert Func
func Insert(key string, value []byte){
	
	if(!IsConnected()){
		Connect()
	} 
		
	if redisC == nil {
		logger.Logger.Info("Error : Redis client is not initialized")
		return
	}
	err := redisC.Set(ctx, key, value, 0).Err()
      	if err != nil{
			logger.Logger.Error("Error : ", err)
		}
		
	
}

// IsConnected Func
func IsConnected() bool {
	if redisC == nil{
		return false
	}
	_, err := redisC.Ping(ctx).Result()

	return err == nil
}