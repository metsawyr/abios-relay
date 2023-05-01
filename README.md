# Needed improvements
## Scope to be production ready
1. Application endpoints are returning internal errors (as in Redis errors). In production these errors should be logged with correlation id and return something like "server error" message.
2. E2E tests

## Other improvements
1. Overall code structure is suboptimal. Business logic should be extracted into "domain" package and operate with its' own domain entities instead of passing through Abios entities from AbiosClient.
2. AbiosClient can be extended to fetch /games endpoint and cache the response for an hour (or more). All existing endpoints will be able to match against this values for the riched payloads without any call to the actual API
# Run locally
1. Copy .env.example and fill in respective values onto .env
2. Run `docker-compose -f local.yaml up`