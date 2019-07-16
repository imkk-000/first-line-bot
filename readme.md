# My First Line Bot

## How to Build in My Local

```sh
# build and push docker image with 'latest' tag
$   docker build -t anonyfz/first-line-bot .
$   docker push anonyfz/first-line-bot
```

## How to Deploy on My Server

```sh
$   docker stop line-bot
$   docker rm line-bot
$   docker pull anonyfz/first-line-bot
$   docker run --name line-bot -d \
      -e "CHANNEL_SECRET=this_is_line_channel_secret" \
      -e "CHANNEL_TOKEN=this_is_line_channel_token" \
      -e "VIRTUAL_HOST=domain" \
      -e "LETSENCRYPT_HOST=domain" \
      anonyfz/first-line-bot
```
