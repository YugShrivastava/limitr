# limitr
A collection of rate-limiters made using different techniques(token bucket, per client, tollbooth) in go

## token-bucket
- package used "golang.org/x/time/rate"
- uses a token bucket to limit the api calls
- once the bucket overflows it starts rejecting api calls

## per-client
- package used "golang.org/x/time/rate"
- uses a per client model, which is bound by client's ip
- a client has x amount of api calls per minute/second

## tollbooth
- package used "github.com/didip/tollbooth/v7"
- great package for simple rate limiter
- implemented the same limiter using tollbooth

> API is kept simple for clear understanding and duped in all instances for better comparison