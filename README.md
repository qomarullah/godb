#### Synopsis

GoDB provides convenient API to query from DB with custom SQL defined based on parameter, built with golang (beego)

#### API Example
  ```
  curl -X GET "http://localhost:8080/v1/query/select?ds=localhost&sqlid=user" -H "accept: application/json"
  ```
  ##### Configuration App.conf
  ```
  #ds
  ds.localhost= "apps:aplikasi@tcp(localhost:3306)/test|mysql|10|10|120000"
  sqlid.user="SELECT * from `user` limit 10"
  ```
  
  ```
  result:
 {
  "count": 10,
  "data": [
    {
      "email": "slene",
      "id": "1",
      "name": "testing"
    },
    {
      "email": "someemail@someemailprovider.com",
      "id": "2",
      "name": "First"
    },
    {
      "email": "someemail@someemailprovider.com",
      "id": "3",
      "name": "First"
    }
    ....
  ],
  "desc": "-",
  "status": false
}
```
```
curl -X GET "http://localhost:8080/v1/query/select?ds=localhost2&sqlid=user2&id=1" -H "accept: application/json"
```
#### Configuration App.conf - flexible parameter eg.[id]
```
ds.localhost2= "apps:aplikasi@tcp(localhost:3306)/mfs|mysql|10|10|120000"
sqlid.user2="SELECT * from `user` where id=[id]"
```
```
Result
{
  "count": 1,
  "data": [
    {
      "email": "slene",
      "id": "1",
      "name": "testing"
    }
  ],
  "desc": "-",
  "success": true
}
```

#### Installation or Development

1. install golang
2. clone this project
3. go get github.com/astaxie/beego 
4. update config in 'conf/App.conf'
5. run bee run -downdoc=true -gendoc=true 
    or bee run
    or go run godb.go
6. test API http://localhost:8080/swagger/
7. dashboard & monitoring http://localhost:8088

#### API Reference
- https://beego.me
- https://github.com/elgs/gosqljson


#### Tests

Swagger
http://localhost:8080/swagger/

Load Test with ab 3x times faster to java based

ab -c10 -n5000 "http://localhost:8080/v1/query/select?ds=localhost&sqlid=user"
```
Concurrency Level:      10
Time taken for tests:   1.386 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      5030000 bytes
HTML transferred:       4275000 bytes
<b>Requests per second:    3608.79 [#/sec] (mean)</b>
Time per request:       2.771 [ms] (mean)
Time per request:       0.277 [ms] (mean, across all concurrent requests)
Transfer rate:          3545.35 [Kbytes/sec] received
```

compare to java apidb 

ab -c10 -n5000 "http://localhost:9002/getlist?jndi=localhost&sqlid=user"
```
Concurrency Level:      10
Time taken for tests:   5.160 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      3295000 bytes
HTML transferred:       2660000 bytes
<b>Requests per second:    968.93 [#/sec] (mean)</b>
Time per request:       10.321 [ms] (mean)
Time per request:       1.032 [ms] (mean, across all concurrent requests)
Transfer rate:          623.56 [Kbytes/sec] received
```


#### Contributors

Let people know how they can dive into the project, include important links to things like issue trackers, irc, twitter accounts if applicable.

#### License

A short snippet describing the license (MIT, Apache, etc.)
