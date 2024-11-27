# SimpleChat

## Backend written on Golang

### [Swagger Docs](https://150.241.82.68/api/swagger/index.html)

### [WebSocket Docs](./app_messanger/websocket_docs.md)


### `.env` file must contain:

```dotenv
GO_PORT=8000

SECRET=sample_secret_key_for_jwt

# comma-separated allowed origins for CORS
CORS_ALLOWED_ORIGINS="http://localhost:5173"
# comma-separated allowed methods for CORS
CORS_ALLOWED_METHODS="GET,HEAD,POST,OPTIONS"

# true OR false
COOKIES_SECURE=true
# string (DefaultMode || LaxMode || StrictMode || NoneMode)
COOKIES_SAME_SITE=DefaultMode

```

### Before deploy project you must have SSL cert on you server:

1. Save cert and private key to dir `/etc/ssl/my_certs`
2. Name cert like `certificate.crt`
3. Name private key like `privateKey.key`

> \* ***HINT*** *
> <br>
> For dev you may obtain a `self-signed certificate` with command:
> <br>
> ```shell
> mkdir /etc/ssl/my_certs
> cd /etc/ssl/my_certs
> openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout privateKey.key -out certificate.crt
> ```
