# erk

a inline issue manager inspired by [watson-ruby](https://github.com/nhmood/watson-ruby)

![gif](http://i.gyazo.com/e9f49623a437a6c9611ecbe8c53981ab.gif)

## Configuration

erk supports to configure default settings by editing `erkconfig.json` file.

here is an example

```
{
  // "[todo] - some issue title" will find.
  "label": "todo",
  // whether sync local issue comment to remote.
  "remote": true,
  // remtoe configuration. now erk supports only github.
  "remote_config": {
    "repo": "takashi/erk",
    "api_token": "hogehogehogheogoehgoe"
  }
}
```
