## Synopsis

GoDB provides convinent API to query from any DB and with custom SQL , built using golang (beego)

## Code Example
  ```
  curl -X GET "http://localhost:8080/v1/query/select?ds=localhost&sqlid=user" -H "accept: application/json"
  ```
  ```
  result:
 {
  "count": 0,
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
  ],
  "desc": "-",
  "status": false
}

 ```
## Configuration
```
#ds
ds.localhost= "apps:aplikasi@tcp(localhost:3306)/test|mysql|10|10|120000"
sqlid.user="SELECT * from `user` limit 10"

ds.localhost2= "apps:aplikasi@tcp(localhost:3306)/mfs|mysql|10|10|120000"
sqlid.user2="SELECT * from `config` limit 10"

```

## Installation

1. install golang
2. clone this project
3. update config in 'conf/App.conf'
4. run go run godb

## API Reference
- https://beego.me
- https://github.com/elgs/gosqljson


## Tests

Describe and show how to run the tests with code examples.

## Contributors

Let people know how they can dive into the project, include important links to things like issue trackers, irc, twitter accounts if applicable.

## License

A short snippet describing the license (MIT, Apache, etc.)
