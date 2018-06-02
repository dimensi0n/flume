# How to use the File storage

For each http requests you'll have to add a `token` field inside your http request header (If you don't choose a token by using `-t yourtoken` while launching flume,, the default is 1234)

## Post a new file
```
POST /fs/post
```
Parameters to give :

* `file` : The file you want to post

The body must have *multipart* type.

## Get a file

*you don't need the token*

```
GET /public/{name}
```

* `{name}` : The file you want to get

`/public/` is the equivalent to a CDN