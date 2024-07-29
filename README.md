# postgres
psql "postgres://user:password@127.0.0.1:5432/database

# helm
kubectl create secret generic quree-env-secrets --from-env-file=.env.prod -n quree

# nats port-forward
k port-forward svc/nats 4222:4222 -n nats

# pub-test
nats pub messages.tg123 '{"chat_id": "", "bot_token": "", "text": "hello {{.Count}}"}' --count=10

# ngrok

ngrok http --domain=cunning-communal-mudfish.ngrok-free.app 80
