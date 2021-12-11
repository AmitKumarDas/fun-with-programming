```yaml
- https://arslan.io/2019/07/03/how-to-write-idempotent-bash-scripts/
```

```yaml
- https://benhoyt.com/writings/goawk
- interpreter
- a web server log file
- line format "timestamp method path ip status time"

- 2018-11-07T07:56:34Z GET /about 1.2.3.4 200 0.013
- 2018-11-07T07:56:35Z GET /contact 1.2.3.4 200 0.020
- 2018-11-07T07:57:00Z GET /about 2.3.4.5 200 0.014
- 2018-11-07T08:00:02Z HEAD / 4.5.6.7 200 0.008
- 2018-11-07T08:05:57Z GET /robots.txt 201.12.34.56 404 0.004

- get IP addresses i.e. field 4 
- of all hits to the /about page

- $ awk '/about/ { print $4 }' server.log 
- 1.2.3.4
- 2.3.4.5

- determine the average response time (field 6) of all GET requests
- sum the response time 
- count the number of GET requests
- print the average in the END block – 18 milliseconds

- $ awk '/GET/ { total += $6; n++ } END { print total/n }' server.log 
- 0.0186667

- AWK supports hash tables 
- i.e. “associative arrays”
- so you can print the count of each request method
– remember the regex pattern is optional and omitted here

- $ awk '{ num[$2]++ } END { for (m in num) print m, num[m] }' server.log 
- GET 9
- POST 1
- HEAD 2

- AWK has two scalar types, string and number
- but it’s been described as “stringly typed”

- comparison operators like == and < 
- do numeric comparisons if number
- else string comparisons
```

### dig to yaml via AWK & jq
```sh
dig acme-test.jimdo-platform-eks-stage.net +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'
```
