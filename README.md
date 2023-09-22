# podcast-server

## Usage

```bash
$ rec_radiko_ts.sh -s TBS -f 202104210100 -d 120 -o output/爆笑問題カーボーイ`date +%Y年%m月%d日`_`date +%Y%m%d%H%M` # https://github.com/uru2/rec_radiko_ts
$ podcastserver -baseurl http://localhost:3333 -targetdir ./output
```

## Usage with BasicAuth

```bash
$ podcastserver -baseurl http://localhost:3333 -targetdir ./output -basicauth user:password
```
