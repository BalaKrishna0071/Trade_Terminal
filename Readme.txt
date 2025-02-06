Project Name = Trade_Terminal 

TeckStack := 

			1) Go
			2) Redis 
			3) KiteConnect API 
			 

Module Package :=

	            1) instrument 
            	2) redis_cache
                3) token_management
            	4) websocket
	            5) loggers

Libraries :=

			1) godotenv
			2) zerodha models
			3) zerodha ticker
			4) go-redis
			5) Air 
			6) Logrus
			7) rotateslog

Files :=

		1) instrument_mgmt
		2) loggers
		3) redis_cache
		4) build-error
		5) token_mgmt
		6) zerodha_ws
		7) Readme
		8) test_instrument
		9) test_redis
		10) test_websocket
		11) .env

Steps :=

   1 ) creating .env file for storing Credentials api keys , secret key, access token .
   2 ) instrument token generation .
   3 ) filtering instrument according to the NIFTY, BANKNIFTY, FINNIFTY .
   4 ) Filter all options and indices and store in a array, including all three instruments .
   5 ) subscribe to the zerodha of this array .
   6 ) making redis connection for storing incoming Quotes .
   7 ) making websocket connection .
   8 ) storing instrument data in redis based key n value where key is instrumentToken and value is Lastprice .
   9 ) converted string into json and stored in redis , full details of particular stock .
   10) retrieving all stocks along with options and indices and merging into array and updating into redis .
   11) creating a func for getting all the a) NiftContract b) BankNiftyContract c) FinNiftyContract including Options and Indices .
   12) in redis key is instrumentSymbol and other details of it .
   13) adding loggers for login info, error, warning etc .  
   14) implementing loggers using logrus package  .
   15) implementing file rotation using rotation using rotateslog package .
   16) file rotation like after 7 days log file will be deleted and within 5 min new file will be created for storing log .