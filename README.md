# radicaster

## Warning!!

```bash
$ ulimit -n 16384
```

## 設定

ここで指定する曜日は、ラジオ番組で利用される曜日ではないので注意。例えば「火曜25時」なら水曜を指定する。

```yaml
programs:
  - cron: "0 4 * * *"
    title: "オールナイトニッポン"
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: LFR
    start: "0100"
    path: ann
  - cron: "0 5 * * *"
    title: "オールナイトニッポンZERO"
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: LFR
    start: "0300"
    path: ann_zero
  - cron: "0 4 * * *"
    titile: "TBS JUNK"
    weekdays:
      - tue
      - wed
      - thu
      - fri
      - sat
    station: TBS
    start: "0100"
    path: junk
```

## Usage

```bash
$ radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml -targetdir ./output
```

## Usage with BasicAuth

```bash
$ radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml  -targetdir ./output -basicauth user:password
```
