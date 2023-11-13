# radicaster

## Warning!!

```bash
$ ulimit -n 16384
```

## 設定

ここで指定する曜日は、ラジオ番組で利用される曜日ではないので注意。例えば「火曜25時」なら水曜を指定する。

```yaml
programs:
- title: オールナイトニッポン
  weekdays:
    - Tuesday
    - Wednesday
    - Thursday
    - Friday
    - Saturday
    - Sunday
  cron: 10 3 * * 0-6
  station: LFR
  start: "0100"
  encoding: aac
  image_url: http://example/image.png
  path: ann
- title: オールナイトニッポン(ZERO)
  weekdays:
    - Tuesday
    - Wednesday
    - Thursday
    - Friday
    - Saturday
    - Sunday
  cron: 30 5 * * 0-6
  station: LFR
  start: "0300"
  encoding: aac
  image_url: http://example/image.png
  path: ann_zero
- title: TBS JUNK
  weekdays:
    - Tuesday
    - Wednesday
    - Thursday
    - Friday
    - Saturday
  cron: 10 3 * * 1-6
  station: TBS
  start: "0100"
  encoding: aac
  image_url: http://example/image.png
  path: junk
- title: MBSヤングタウン
  weekdays:
    - Monday
    - Tuesday
    - Wednesday
    - Thursday
    - Friday
    - Saturday
    - Sunday
  cron: 40 23 * * *
  area: JP27 # エリア外の番組はエリアIDの指定が必要(JP27=大阪)
  station: MBS
  start: "2200"
  encoding: aac
  image_url: http://example/image.png
  path: yantan
zenroku:
  enable: true
  cron: 0 3 * * *
  encoding: aac
  stations:
    jorf:
      image_url: http://example/image.png
    lfr:
      image_url: http://example/image.png
    qrr:
      image_url: http://example/image.png
    tbs:
      image_url: http://example/image.png
  enable_stations:
    - tbs
    - qrr
    - lfr
    - jorf
```

## Usage

```bash
$ radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml -targetdir ./output
```

## Usage with BasicAuth

```bash
$ radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml  -targetdir ./output -basicauth user:password
```

## Usage with Premium

```bash
$ radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml  -targetdir ./output -basicauth user:password -radikoemail "${RADIKO_EMAIL}" -radikopassword "${RADIKO_PASSWORD}"
```
