# Kiki

Basic configuration:

```
# $HOME/.config/kiki/config.yaml
---
database:
  user: kiki
  name: kiki
  password: kiki
  host: localhost
  port: 5432
```


Optionally, you can specify with `--config` parameter.

# Usage

## Add User

```
# Misskey example
./kiki add-account --name="happy_bot" \
                   --api-token="user token from misskey" \
                   --base-url="https://slippy.xyz" \
                   --publisher="misskey"
```


## Add feed to a user

```
./kiki add-feed --name="Happy News" --url="https://i-am-happy/newletter.atom" --user="happy_bot" --provider="atom"
```

## Fetch

```
./kiki fetch
```

## Create a post

Kiki is designed the way she is not able to post all the new items in one batch.
Kiki will publish only the oldest one and mark it as posted.

```
./kiki publish
```

## Cron job can be

```
*/15 * * * * /home/user/kiki fetch && /home/user/kiki publish
```
