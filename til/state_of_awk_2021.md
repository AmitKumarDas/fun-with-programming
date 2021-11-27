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
```
