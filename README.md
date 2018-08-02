# Build a simple prepaid card

## How to build

`make build`

## Example of use following a basic flow to simulate normal operation

### Create a new card
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/cards/new -d '{"name":"christian espinoza"}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:49:29 GMT
Content-Length: 76

{"id":1,"number":"4916199871272266","name":"christian espinoza","balance":0}%  
```
### Make a deposit
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/transaction/deposit -d '{"card_id":1, "amount":3000}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:50:08 GMT
Content-Length: 31

{"total":3000,"operation_id":1}% 
```
### Make an Authorization Request to block an amount of £
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/claim/block_auth -d '{"card_id":1, "amount":1000}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:50:38 GMT
Content-Length: 22

{"blocking_auth_id":1}% 
```

### Cancel Authorization Request to unblock blocked amount of £
`curl -i -H "Content-type: application/json" -X DELETE http://localhost:8080/apiv1/claim/block_auth -d '{"card_id":1, "blocking_auth_id":1}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:50:57 GMT
Content-Length: 15

{"result":"ok"}%   
```
### Request Balance
`curl -i -H "Content-type: application/json" -X GET http://localhost:8080/apiv1/cards/balance -d '{"card_id":1}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:51:19 GMT
Content-Length: 28

{"balance":3000,"blocked":0}%  
```
### Make a second Authorization Request to block an amount of £
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/claim/block_auth -d '{"card_id":1, "amount":1000}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:51:38 GMT
Content-Length: 22

{"blocking_auth_id":1}%  
```

### Capture an amount of £
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/transaction/capture -d '{"card_id":1, "blocking_auth_id":1}'`

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:51:57 GMT
Content-Length: 34

{"captured":1000,"operation_id":2}%     
```
### Request Balance again
`curl -i -H "Content-type: application/json" -X GET http://localhost:8080/apiv1/cards/balance -d '{"card_id":1}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:52:23 GMT
Content-Length: 28

{"balance":2000,"blocked":0}% 
```

### Refund an amount of £
`curl -i -H "Content-type: application/json" -X POST http://localhost:8080/apiv1/transaction/refund -d '{"card_id":1, "operation_id":2, "amount":1000}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:52:43 GMT
Content-Length: 41

{"refunded_amount":1000,"operation_id":4}%   
```
### Request Balance again
`curl -i -H "Content-type: application/json" -X GET http://localhost:8080/apiv1/cards/balance -d '{"card_id":1}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:53:03 GMT
Content-Length: 28

{"balance":3000,"blocked":0}%   
```
### Request statement for card with id 1
`curl -i -H "Content-type: application/json" -X GET http://localhost:8080/apiv1/cards/statement -d '{"card_id":1}'`
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 02 Aug 2018 02:54:36 GMT
Content-Length: 529

[{"id":1,"kind":"deposit","datetime":"2018-08-02T03:50:08.440228727+01:00","amount":3000,"status":"done","card_id":1,"merchant_id":0},{"id":2,"kind":"capture","datetime":"2018-08-02T03:51:57.4107068+01:00","amount":1000,"status":"done","card_id":1,"merchant_id":0},{"id":3,"kind":"deposit","datetime":"2018-08-02T03:52:43.210231003+01:00","amount":1000,"status":"done","card_id":1,"merchant_id":0},{"id":4,"kind":"refund","datetime":"2018-08-02T03:52:43.21026022+01:00","amount":1000,"status":"done","card_id":1,"merchant_id":0}]% 
```