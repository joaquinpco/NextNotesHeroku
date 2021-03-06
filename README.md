# NextNotesHeroku

## Install dependencies

```bash
go mod tidy
```

## Run project

```bash
go run main.go
```

### Credentials to run in local

```json
You need the file credentials.json with credentials of your project firebase.
Example of file:
{
  "type": "service_account",
  "project_id": "...",
  "private_key_id": "...",
  "private_key": "...",
  "client_email": "...",
  "client_id": "...",
  "auth_uri": "...",
  "token_uri": "...",
  "auth_provider_x509_cert_url": "...",
  "client_x509_cert_url": "..."
}
```

## Test in local

```bash
code ext install humao.rest-client
go run main.go
code nextnotes-firebase/notes/note.http
#enjoy!
```

## Upload to Heroku

```bash
  heroku login
  heroku git:remote -a nextnotes
  heroku container:login
  heroku stack:set container
  heroku labs:enable runtime-new-layer-extract
  heroku container:push web
  heroku container:release web
```

## Up with Docker

```bash
docker-compose up -d --build
```
