# order-book

Send a Bid:
```
url -L -X POST -H "Content-Type: application/json" 'http://localhost:3001/api/v3/order/?symbol=LTCBTC' -d '{"id":1,"price":10.00,"quantity":20.00,"side":"BUY","timestamp":123}'
{"orderId":0}
```
See: https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#new-order-trade

Send an Ask:
```
curl -L -X POST -H "Content-Type: application/json" 'http://localhost:3001/api/v3/order/?symbol=LTCBTC' -d '{"id":1,"price":15.00,"quantity":30.00,"side":"SELL","timestamp":123}'
```

Get the "Depth" of the order book:
```
curl -L -X GET -H "Content-Type:application/json" 'http://localhost:3001/api/v3/depth?symbol=LTCBTC'                             ✔  took 30s   at 15:01:03 
```
See: https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#order-book
