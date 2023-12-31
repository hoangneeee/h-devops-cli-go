# H-devops

![Release](https://github.com/hoangneeee/h-devops-cli-go/actions/workflows/release_build.yml/badge.svg)

## How to install
`h-devops` is available for MacOS and Linux based operating systems.

```shell
curl -L https://raw.githubusercontent.com/hoangneeee/h-devops-cli-go/master/install.sh | bash
```
or
```shell
git clone https://github.com/hoangneeee/h-devops-cli-go
cd h-devops-cli-go
go build -o bin/h-devops main.go
sudo mv bin/h-devops /usr/local/bin
sudo chmod a+x /usr/local/bin/h-devops
```

---
## Feature
### Common commands
- See all available commands
```shell
h-devops cmd
```

### Get template services commands
- Backup postgres to S3
```shell
h-devops postgres-backup-s3
```
- Elastic snapshot to S3
```shell
h-devops ens
```

### Docker commands
- Setup docker env
```shell
h-devops docker i
```
- Add user to docker group
```shell
h-devops docker add <username>
```

### Linux commands
- Add username to sudoers
```shell
h-devops su <username>
```

### Node commands
- Install NVM (Node version manager)
```shell
h-devops nvm i
```

### Certbot commands
- Install Certbot
```shell
h-devops cert i
```
- Auto-renew Let's encrypt certificate for Nginx
```shell
h-devops cert a
```
- Check Certificates expiry date
```shell
h-devops cert ex
```

### PHP Helper commands
- Install PHP version (Default: 7.4)
```shell
h-devops php i
h-devops php i -v 8.0
```
- Remove PHP version
```shell
h-devops php r -v 8.*
```

### Fail2Ban commands
- Install Fail2Ban
```shell
h-devops f2b i
```
- Configure Fail2Ban
```shell
h-devops f2b c
```
---

## How to develop
Required docker-compose version 2.22.0 or higher
```shell
docker-compose watch  
```
or 
```shell
cd h-devops-cli-go
go get -d ./...
go run main.go
```
---

## Issue
Please open an issue [New issue](https://github.com/hoangneeee/h-devops-cli-go/issues)

## License

See [`LICENSE`](./LICENSE)
