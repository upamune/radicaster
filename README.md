# podcast-server

## Warning!!

```bash
$ ulimit -n 16384
```

## 設定

ここで指定する曜日は、ラジオ番組で利用される曜日ではないので注意。例えば「火曜25時」なら水曜を指定する。

```yaml
programs:
  - cron: "0 4 * * *" # オールナイトニッポン
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: LFR
    start: "0100"
  - cron: "0 5 * * *" # オールナイトニッポンZERO
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: LFR
    start: "0300"
  - cron: "0 4 * * *" # TBS JUNK
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: TBS
    start: "0100"
```

## Usage

```bash
$ rec_radiko_ts.sh -s TBS -f 202104210100 -d 120 -o output/爆笑問題カーボーイ`date +%Y年%m月%d日`_`date +%Y%m%d%H%M` # https://github.com/uru2/rec_radiko_ts
$ podcastserver -baseurl http://localhost:3333 -targetdir ./output
```

## Usage with BasicAuth

```bash
$ podcastserver -baseurl http://localhost:3333 -targetdir ./output -basicauth user:password
```
