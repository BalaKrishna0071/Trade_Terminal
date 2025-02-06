package instrument_mgmt
import(
	logger "trade/loggers"
	"slices"
	zerodhaconnect	"github.com/zerodha/gokiteconnect/v4"
)

// NiftyContract Func
func NiftyContract(apiKey string)(map[string]int, error){
	kiteConnect := zerodhaconnect.New(apiKey)

	// Creating Map
	var instrumentMap = make(map[string]int)

	// Calling GetInstrument Inbuild Func
	data, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}

	// Filtering Instrument Through loop and Based on Condition
	for _, val := range data{
		
		// This will  Filter only Option Contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Name == "NIFTY" && val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType )&& expiry == "2025-02-06" {
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
			
		}
	
		// This will Filter Indices
		if val.Name == "NIFTY 50" && val.Segment == "INDICES" && val.InstrumentType == "EQ" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}
	
	return instrumentMap, nil
	
}

// BankNiftyContract Func
func BankNiftyContract(apiKey string)(map[string]int, error){
	kiteConnect := zerodhaconnect.New(apiKey)

	// Creating Map
	var instrumentMap = make(map[string]int)

	// Calling GetInstrument Inbuild Func
	data, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}

	// Filtering Instrument Through loop and Based on Condition
	for _, val := range data{

		// This will Filter only Option Contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType) && val.Name == "BANKNIFTY" && expiry == "2025-02-27" && val.Exchange == "NFO"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}

		// This will Filter Indices
		if val.Segment == "INDICES" && val.Name == "NIFTY BANK" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}
	
	return instrumentMap, nil
}

// FinNiftyContract Func
func FinNiftyContract(apiKey string)(map[string]int, error){
	kiteConnect := zerodhaconnect.New(apiKey)

	// Creating Map
	var instrumentMap = make(map[string]int)

	// Calling GetInstrument Inbuild Func
	data, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}

	// Filtering Instrument Through loop and Based on Condition
	for _, val := range data{

		// this filter only option contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType) && val.Name == "FINNIFTY" && expiry == "2025-02-27" && val.Exchange == "NFO"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}

		// this will filter indices
		if val.Segment == "INDICES" && val.Name == "NIFTY FIN SERVICE" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}

	return instrumentMap, nil
}

// GetAllInstrument Func
func GetAllInstrument(apiKey string)(map[string]int, error){
	kiteConnect := zerodhaconnect.New(apiKey)
	var instrumentMap = make(map[string]int)

	// nifty
	data1, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}
	for _, val := range data1{
		
		// this filter only option contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Name == "NIFTY" && val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType )&& expiry == "2025-02-06" {
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
			
		}
	
		// this will filter indices
		if val.Name == "NIFTY 50" && val.Segment == "INDICES" && val.InstrumentType == "EQ" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}

	// banknifty
	data2, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}

	for _, val := range data2{

		// this filter only option contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType) && val.Name == "BANKNIFTY" && expiry == "2025-02-27" && val.Exchange == "NFO"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}

		// this will filter indices
		if val.Segment == "INDICES" && val.Name == "NIFTY BANK" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}

	// finnifty
	data3, err := kiteConnect.GetInstruments()
	if err != nil{
		logger.Logger.Error(err)
	}
	for _, val := range data3{

		// this filter only option contract
		expiry := val.Expiry.Time.Format("2006-01-02")
		if val.Segment == "NFO-OPT" && slices.Contains([]string{"CE","PE"}, val.InstrumentType) && val.Name == "FINNIFTY" && expiry == "2025-02-27" && val.Exchange == "NFO"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}

		// this will filter indices
		if val.Segment == "INDICES" && val.Name == "NIFTY FIN SERVICE" && val.Exchange == "NSE"{
			instrumentMap[val.Tradingsymbol] = val.InstrumentToken
		}
	}

	return instrumentMap, nil
}