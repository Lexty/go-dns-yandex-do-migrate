# go-dns-yandex-do-migrate

Simple tool for migrating DNS records from Yandex.Connect to DigitalOcean Networking.

## Export from Yandex.Connect

```bash
Extract yandex DNS records

Usage:
  dns-yandex-do-migrate yandex [flags]

Flags:
  -d, --domain string   Domain name
  -h, --help            help for yandex
  -p, --pretty-print    Pretty print JSON
  -t, --token string    Admin token
```

You need to get admin token.

Go to the link: https://pddimp.yandex.ru/api2/admin/get_token

### Example

```bash
./dns-yandex-do-migrate yandex \
  -t FSDSDGFARDFK9SDK09SAFGMU0923RGUMN09EFDM09UMRG98U3WMR \
  -d example.org \
  -p > yandex-dns-records.json
```

## Import to DigitalOcean Networking

```bash
Loads DNS records to DigitalOcean

Usage:
  dns-yandex-do-migrate do [flags]

Flags:
  -d, --domain string   Domain name
  -h, --help            help for do
  -t, --token string    Access token
```

You need to get access token. 

Go to the link: https://cloud.digitalocean.com/settings/applications

### Example

```bash
cat yandex-dns-records.json | ./dns-yandex-do-migrate do \
  -t 3483284952cm348tc23m984tm39c854tymnc9485gymn9485fymnc98435ymnvf4 \
  -d example.org
```

## One command migrate example Yandex -> DigitalOcean

```bash
./dns-yandex-do-migrate yandex \
  -t FSDSDGFARDFK9SDK09SAFGMU0923RGUMN09EFDM09UMRG98U3WMR \
  -d example.org | ./dns-yandex-do-migrate do \
  -t 3483284952cm348tc23m984tm39c854tymnc9485gymn9485fymnc98435ymnvf4 \
  -d example.org
```
