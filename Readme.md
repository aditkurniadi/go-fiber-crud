Project Task List: Go-Fiber CRUD with Docker

Status: implemented

Implemented stack:
- Go Fiber for HTTP server and logging middleware
- GORM with PostgreSQL
- `.env` loading with `godotenv`
- CRUD API for `Task`
- Auto migration on startup
- Dockerfile and Postman collection

Run local:
1. Start database with `docker-compose up -d`
2. Start app with `go run .`
3. Open `http://localhost:3000/api/tasks`

API endpoints:
- `POST /api/tasks`
- `GET /api/tasks`
- `GET /api/tasks/:id`
- `PUT /api/tasks/:id`
- `DELETE /api/tasks/:id`

Docker deploy:
- Use [Dockerfile](Dockerfile) for the app image
- Use [docker-compose.prod.yml](docker-compose.prod.yml) if you want Dokploy to run app + Postgres as one stack
- In Dokploy, set the app service port to `3000`
- If port `80` is already occupied by another project, do not publish this app directly to `80`; route it through Dokploy's domain/proxy to the internal port `3000`
- Set database host to `db` in production so the app connects to the Postgres container instead of `localhost`

Postman environment:
- Import [postman_environment.json](postman_environment.json) into Postman
- Import [postman_collection.json](postman_collection.json) after that
- Select environment `Go Fiber CRUD Local`
- Edit `base_url`, `task_id`, `task_title`, `task_description`, and `task_completed` from the environment UI

Payload example:
```json
{
	"title": "Belajar Fiber",
	"description": "Membuat CRUD sederhana",
	"completed": false
}
```