package main
import(
	 instmt "trade/instrument"
	 logger"trade/loggers"
	"os"
	"github.com/joho/godotenv"
)

func main(){

	// Accessing .env Credentials
	err := godotenv.Load(".env")
	if err != nil{
		logger.Logger.Error(err)
		
	}
	
	// Accessing API Key
	apiKey := os.Getenv("ZERODHA_API_KEY")
	logger.Logger.Info("API Key : ", apiKey)
	

	// Accessing API Secret
	apiSecret := os.Getenv("ZERODHA_API_SECRET")
	logger.Logger.Info("API Secret Key : ", apiSecret)
	

	// Calling niftyContract
	instrumentList1, err := instmt.NiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
	}
	logger.Logger.Info("Nifty     length : ",len(instrumentList1))
	

	// Calling BankniftyContract
	instrumentList2, err := instmt.BankNiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
	}
	logger.Logger.Info("Nifty     length : ",len(instrumentList2))


	// Calling FinniftyContract
	instrumentList3, err := instmt.FinNiftyContract(apiKey)
	if err != nil{
		logger.Logger.Error(err)
	}
	logger.Logger.Info("Nifty     length : ",len(instrumentList3))
	

	// Calling GetAllInstrument
	allInstrumentList, err := instmt.GetAllInstrument(apiKey)
	if err != nil{
		logger.Logger.Error(err)

	}
	logger.Logger.Info("Nifty     length : ",len(allInstrumentList))
	
}