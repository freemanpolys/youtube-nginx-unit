# Sample
https://unit.nginx.org/howto/samples/#go
https://www.nginx.com/products/nginx-unit
https://itnext.io/performance-comparison-between-nginx-unit-and-uwsgi-python3-4511fc172a4c
https://unit.nginx.org/

# Get configuration
curl    --unix-socket /usr/local/var/run/unit/control.sock http://localhost/config/ 
# Enable nginx conf
curl -X PUT -d @config.json   --unix-socket /usr/local/var/run/unit/control.sock http://localhost/config/ 

# Test
curl localhost:1180

# Delete conf
curl -X DELETE -d @config.json   --unix-socket /usr/local/var/run/unit/control.sock http://localhost/config/ 