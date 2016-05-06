# griffon

## Usage

add `nameserver 127.0.0.1` in /etc/resolv.conf

```
$ griffon -p 53
```

Add entries 

```
$ curl -H "Content-Type: application/json" -X POST -d '{"name":"mysql","ip":"127.0.0.1", "port": 3306 }' http://localhost:3000/api/v1/entries
```

now use dig 

```
$ dig mysql.service.consul

; <<>> DiG 9.8.3-P1 <<>> mysql.service.consul
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 47942
;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;mysql.service.consul.        IN    A

;; ANSWER SECTION:
mysql.service.consul.    0    IN    A    127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Sat May  7 00:40:03 2016
;; MSG SIZE  rcvd: 74
```

or

```
$ dig mysql.service.consul SRV

; <<>> DiG 9.8.3-P1 <<>> mysql.service.consul SRV
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 48708
;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;mysql.service.consul.        IN    SRV

;; ANSWER SECTION:
mysql.service.consul.    0    IN    SRV    1 1 3360 mysql.service.consul.

;; ADDITIONAL SECTION:
mysql.service.consul.    0    IN    A    127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Sat May  7 00:34:00 2016
;; MSG SIZE  rcvd: 134
```
