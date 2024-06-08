nginx access logfile(access.log) Parser example

nginx log_format example:

```
    log_format upstream2  '[$proxy_add_x_forwarded_for]-[$geoip2_data_country_code]-[$geoip2_data_province_name]-[$geoip2_data_city]'
        ' $remote_user [$time_local] "$request" $http_host'
        ' [$body_bytes_sent] $request_body "$http_referer" "$http_user_agent" [$ssl_protocol] [$ssl_cipher]'
        ' [$request_time] [$status] [$upstream_status] [$upstream_response_time] [$upstream_addr]';
```
