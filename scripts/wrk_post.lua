-- example HTTP POST script which demonstraes setting the
-- HTTP method, body, and adding a header

wrk.method = "POST"
wrk.body = '[{"title": "title", "post": "post"}]'
wrk.headers["Content-Type"] = "application/json"