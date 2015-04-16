## Meetup Leave Group UI

Go server with minimal UI for leaving meetup.com groups.

## Use

The server uses your meetup api key to list and leave groups.

```bash
export MEETUP_API_KEY=your-api-key
go build server.go
./server
open http://localhost:8080/groups
```
