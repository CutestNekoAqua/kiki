# Kiki

This project is still in a WIP state.

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
misskey:
  baseurl: https://slippy.xyz
```


Optionally, you can specify with `--config` parameter.

# Usage

## Add User

```
./kiki add-account --name="happy_bot" --api-token="user token from misskey"
```


## Add feed to a user

```
./kiki add-feed --name="Happy News" --url="https://i-am-happy/newletter.atom" --user="happy_bot"
```

## Fetch

```
./kiki fetch
```

## Create a post

Kiki is designed the way she is not able to post all the new items in one batch.
Kiki will publish only the oldest one and mark it as posted.

```
./kiki send
```

## Cron job can be

```
*/15 * * * * /home/user/kiki fetch && /home/user/kiki send
```
