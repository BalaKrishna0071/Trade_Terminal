Here is a well-structured `README.md` for your project:

```markdown
# Trade Terminal

## Project Overview

**Trade Terminal** is a real-time financial data terminal that connects to the Zerodha KiteConnect API for streaming data. It handles instrument token generation, data subscription, storage in Redis, and WebSocket connections to provide live market data. The system also includes logging capabilities, and data is organized and stored for easy access and monitoring.

## Tech Stack

- **Go**: The primary programming language used for the implementation of the project.
- **Redis**: Used as a cache to store instrument data and quotes.
- **KiteConnect API**: Utilized for accessing market data from Zerodha.

## Modules & Packages

1. **instrument**: Manages the instruments and their tokens.
2. **redis_cache**: Handles Redis connections and data storage.
3. **token_management**: Manages instrument token generation and filtering.
4. **websocket**: Subscribes and listens to live market data.
5. **loggers**: Handles logging and log file management.

## Libraries

- **godotenv**: Loads environment variables from `.env` file.
- **zerodha models**: Data models from Zerodha for handling market data.
- **zerodha ticker**: For subscribing to and receiving live market tick data.
- **go-redis**: Redis client for Go.
- **Air**: For environment-based configuration and management.
- **Logrus**: Structured logger for Go.
- **rotateslog**: Handles log file rotation.

## Files

1. **instrument_mgmt**: Logic related to managing and generating instrument tokens.
2. **loggers**: Contains logging configurations and log generation.
3. **redis_cache**: Implements Redis connection and operations.
4. **build-error**: Logs build-related errors.
5. **token_mgmt**: Manages token generation and filtering for instruments.
6. **zerodha_ws**: WebSocket connection and management for Zerodha.
7. **Readme**: This file.
8. **test_instrument**: Unit tests for the instrument module.
9. **test_redis**: Unit tests for Redis caching.
10. **test_websocket**: Unit tests for WebSocket communication.
11. **.env**: Stores sensitive credentials (API keys, tokens, etc.).

## Key Steps for Implementation

1. **Creating `.env` File**: Store sensitive credentials such as API keys, secret keys, and access tokens.
2. **Instrument Token Generation**: Generate tokens for market instruments such as NIFTY, BANKNIFTY, and FINNIFTY.
3. **Instrument Filtering**: Filter instruments based on specific criteria (e.g., NIFTY, BANKNIFTY, FINNIFTY).
4. **Instrument Array Creation**: Create an array with filtered instruments, including options and indices.
5. **Subscription to Zerodha API**: Subscribe to Zerodha’s live market data for the filtered instrument array.
6. **Redis Connection**: Establish a connection to Redis for storing incoming market quotes.
7. **WebSocket Connection**: Set up WebSocket connection to listen for live data from Zerodha.
8. **Data Storage in Redis**: Store instrument data in Redis, where the key is the instrument token and the value is the last price.
9. **Storing Detailed Instrument Data**: Convert data to JSON format and store it in Redis for easy retrieval.
10. **Retrieving Stock Data**: Retrieve data for all instruments (stocks, options, and indices), merge them into an array, and update Redis.
11. **Instrument Contract Retrieval**: Create functions to retrieve data for:
    - Nifty Contract
    - BankNifty Contract
    - FinNifty Contract (including options and indices).
12. **Redis Key Management**: Use instrument symbols as keys in Redis, storing related details.
13. **Logging**: Implement logging for information, errors, and warnings.
14. **Logrus Integration**: Integrate Logrus for structured logging.
15. **File Rotation**: Implement file rotation using `rotateslog`:
    - Log files will be deleted after 7 days.
    - A new log file will be created every 5 minutes for storing log data.

## Installation & Setup

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd trade_terminal
   ```

2. Install dependencies:

   ```bash
   go get -u github.com/joho/godotenv
   go get -u github.com/zerodhatech/gokiteconnect
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/sirupsen/logrus
   go get -u github.com/natefinch/atomic
   ```

3. Set up the `.env` file with your credentials (API keys, access token, etc.):

   ```
   KITE_API_KEY=<your_api_key>
   KITE_API_SECRET=<your_api_secret>
   KITE_ACCESS_TOKEN=<your_access_token>
   REDIS_HOST=<your_redis_host>
   REDIS_PORT=<your_redis_port>
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

## Testing

The project includes unit tests for different modules. To run the tests, use the following command:

```bash
go test -v
```

## Logs & Monitoring

- Logs will be stored in log files, with rotation every 5 minutes and deletion after 7 days.
- Log levels include INFO, ERROR, and WARNING for different types of events.

## Contributing

Feel free to submit a pull request or open an issue if you find a bug or want to contribute new features!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This should provide clear documentation for your project. Let me know if you need any adjustments or additions!