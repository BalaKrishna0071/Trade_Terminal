package main
import(
	"os"
	"github.com/joho/godotenv"
	inst "trade/instrument"
	ws "trade/websockets"
	logger "trade/loggers"
	kiteticker "github.com/zerodha/gokiteconnect/v4/ticker"
)

// Creating Global Variable
var(
	ticker *kiteticker.Ticker
	instToken = []uint32{738561}
)

// Let it be here will explain it later
func OnConnect(){
	logger.Logger.Info("Websocket Connected....")

	// Accessing .env Credentials
	err := godotenv.Load(".env")
	if err != nil{
		logger.Logger.Error("Error : ", err)
	}
	apiKey := os.Getenv("ZERODHA_API_KEY")
	logger.Logger.Info("API Key : " , apiKey)
	 var arr []int

	// NiftyContract
	data, err := inst.NiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
	}
	for _, val:= range data{
		arr = append(arr, val)
	}

	// BankNiftyContract
	data2 , err := inst.BankNiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
	}
	for _, val:= range data2{
		arr = append(arr, val)
	}
	
	// FinNiftyContract
	data3 , err := inst.FinNiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
		
	}
	for _, val:= range data3{
		arr = append(arr, val)
	}
	
	// Appending all InstruementToken in one Array
	var combArr []uint32
	for _, v := range arr{
		combArr = append(combArr, uint32(v))
	}

	logger.Logger.Info("Length : ",len(combArr))

	err = ticker.Subscribe(combArr)
	if err != nil {
		logger.Logger.Error("err : ", err)
	}

	// Set Subscription mode for the Subscribed Token
	err = ticker.SetMode(kiteticker.ModeFull, combArr)
	if err != nil {
		logger.Logger.Error("err : ", err)
	}

}


func main(){
	
	// Accessing .env Credentials
	err := godotenv.Load(".env")
	if err != nil{
		logger.Logger.Error(err)
	}
	
	// Accessing API Key
	apiKey := os.Getenv("ZERODHA_API_KEY")
	logger.Logger.Info("API Key : " , apiKey)

	// Accessing AccessToken
	accToken := os.Getenv("ZERODHA_ACCESS_TOKEN")
	logger.Logger.Info("Access Token : ", accToken)

	// Websocket Methods
	ticker = kiteticker.New(apiKey, accToken)
	ticker.OnError(ws.OnError)
	ticker.OnClose(ws.OnClose)
	ticker.OnConnect(OnConnect)
	ticker.OnTick(ws.OnMessage)
	ticker.OnReconnect(ws.OnReconnect)

	ticker.Serve()
}

