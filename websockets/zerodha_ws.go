package zerodha_ws
import(
	"fmt"
	"time"
	"os"
	"encoding/json"
	logger "trade/loggers"
	// "strconv"
	rd "trade/redis_cache"
	"github.com/joho/godotenv"
	instmt "trade/instrument"
	kitemodels "github.com/zerodha/gokiteconnect/v4/models"
	kiteticker "github.com/zerodha/gokiteconnect/v4/ticker"
	
)
var(
	apiKey string
	dataa1 map[string]int
)

// Creating Gobal Variable
var(
	ticker *kiteticker.Ticker
	instToken = []uint32{738561}
)

// OnError Func
func OnError(err error){
	logger.Logger.Error("Websocket Closed Due to : ", err)
}

// OnClose Func
func OnClose(code int, reason string){
	logger.Logger.Info("Websocket Close Due to Code %d And Reason is %s", code, reason)
}

// OnConnect Func
func OnConnect(){
	logger.Logger.Info("Connected...")
}

// OnMessage Func
func OnMessage(message kitemodels.Tick){
	
	data := map[string]interface{}{
		"Mode": message.Mode,
		"InstrumentToken": message.InstrumentToken,
		"IsTradable ": message.IsTradable,
		"IsIndex": message.IsIndex,
		"Timestamp":message.Timestamp,
		"LastTradeTime": message.LastTradeTime,
		"LastPrice": message.LastPrice,
		"LastTradedQuantity ": message.LastTradedQuantity,
		"TotalBuyQuantity" : message.TotalBuyQuantity,
		"TotalSellQuantity": message.TotalSellQuantity,
		"VolumeTraded" : message.VolumeTraded,
		"TotalBuy" : message.TotalBuy,
		"TotalSell" : message.TotalSell,
		"AverageTradePrice": message.AverageTradePrice,
		"OI": message.OI,
		"OIDayHigh": message.OIDayHigh,
		"OIDayLow": message.OIDayLow,
		"NetChange": message.NetChange,
		"OHLC": message.OHLC,
		"Depth": message.Depth,
		
	}

	if dataa1 == nil{
		err := godotenv.Load(".env")
		if err != nil{
			logger.Logger.Error("error loading .env ", err)
		}

		apiKey = os.Getenv("ZERODHA_API_KEY")
		if apiKey == ""{
			logger.Logger.Info("Error apikey not in env")
		}

		dataa1, err = instmt.GetAllInstrument(apiKey)
		if err != nil {
			logger.Logger.Error("error while calling GetAllInstrument ", err)
			fmt.Println()
		}
	}
	
	// Storing String into Json
	strData, err := json.Marshal(data)
	if err != nil{
		logger.Logger.Error(err)
	}

	var key string
	for keys, keyMap:= range dataa1{
		if int(message.InstrumentToken) == keyMap{
			key = string(keys)
		}
	}

	// Calling Insert func
	rd.Insert(key, strData)
	
}

// OnReconnect Func
func OnReconnect(atmp int, delay time.Duration){
	logger.Logger.Error("Reconnect Attempt : ", atmp , delay)
}