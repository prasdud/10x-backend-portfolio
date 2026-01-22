# 10x Backend Dev Portfolio

I am sick of these cool looking Portfolio sites these frontend and fullstack devs have, partly because im a backend dev.

So this is the Portfolio of a backend dev.
APIS.
Gin.
Go.
Fast.
Goodbye UI/UX, hello curl.

## Usage

### Run with Docker (Recommended)
This starts the server with **hot reloading** using `air`. Any changes you make to your code will automatically recompile and restart the server.

```bash
docker compose up --build
```
The API will be available at `http://localhost:7741/api/v1/ping`.

### Run Locally
If you have Go installed locally:

```bash
go run cmd/server/main.go
```

## API Endpoints
- `GET /api/v1/ping` - Health check
- `GET /api/v1/about` - About me
- `GET /api/v1/resume` - Download resume
- `GET /api/v1/skills` - List skills (filters: `level`, `category`)
- `GET /api/v1/projects` - List projects
- `GET /api/v1/help` - List all commands

# Status, Dev Notes

Actively working on this

## TODO

- setup query params for other stuff
- do some cool backend stuff, thinking of bouncing a packet around the globe
- setup docker and set up as systemd on vps, setup a subdomain for this


## Notes
HTTP -> handler -> store -> DB -> response
