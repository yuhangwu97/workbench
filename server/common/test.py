import requests
import json

url_base = 'https://api.eu.xdr.trendmicro.com'
url_path = '/v2.0/xdr/portal/accounts/roles'
token = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJjaWQiOiJiMTEyODVjZS1hZDIzLTQ4NWQtYmNlMi02MjNkODliOTFhOGUiLCJjcGlkIjoic3ZwIiwicHBpZCI6ImFhayIsIml0IjoxNjkwMTg3MTk1LCJ1aWQiOiIiLCJwbCI6IiIsImV0IjoxNjkwMzU5OTkwLCJ1Y3QiOiJlcmljX3lhbyJ9.V7UePidQUUgsKqG4Cna5Uh1VswAXWu_yK-wqAHClDysPEHqLHrRW4pahTHgRGZHU21SCjVvjta-x-x-Z9V2fEJhebhi-PJv72GCWI8dd8_NBu_M2-2Bi0OqgGWiBvyRuMIADAaNYddOh1qHLQfhqgZEf7veNStfC7HfPWnGbAUruCKsXbFRokRrLbN_qxwEhUYqrOl0SkgAAHZ5j0TEU2BQPvqBvMibffUIkppUsUCNlluy7ayRhIMnPF7B6B8xqYOdXrimdwnPkKVpTj0Alt4pRd0BxEzJT9uRiUZIXCjxB-oQou7lAA2D4wfggGJcqSgJX9h-Ri5WmVH3bWvY1c4-ndKwMImRcAOpzJP3_SBbgN20xVXmovw_Pe-vB5YaXqZphzXyA7IYr5BNdY5OnTynlHSvbiECI3xRBARvLGlonZ8AoE7HXUABeH4dEPHE1mRC_Riqc9D7qA_80VRZdK2pOZi45tHc1VvFAcL65Um7grlMYZbc53INk8EOdCFMKVLy985plJsQCdi3t5U9O9AxLvBSZNll4AsS6jsT7H4Fu_krHG77e6UGc8NO-vuRbRT32RtBcoNub3LJM-ii_sp8FNlQCnMqYSyRfx00NXOpHwZBTUzKNTd3m5AKye6KyRQXkQxAnBsnsQIksS3E1wp1WubP1WgmEqX7CxCOZk-M'

query_params = {}
headers = {'Authorization': 'Bearer ' + token}

r = requests.get(url_base + url_path, params=query_params, headers=headers)

print(r.status_code)
if 'application/json' in r.headers.get('Content-Type', '') and len(r.content):
    print(json.dumps(r.json(), indent=4))
else:
    print(r.text)