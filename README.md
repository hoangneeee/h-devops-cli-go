# H-devops

## How to install
`h-devops` is available for MacOS and Linux based operating systems.

```shell
curl -L https://raw.githubusercontent.com/hoangneeee/h-devops-cli-go/master/install.sh | bash
```

---
## Feature
- Backup postgres to S3
```shell
h-devops postgres-backup-s3
```
- Setup docker env
```shell
h-devops docker i
```
- Add user to docker group
```shell
h-devops docker add <username>
```
- Add username to sudoers
```shell
h-devops su <username>
```
- Install NVM (Node version manager)
```shell
h-devops nvm i
```
- Elastic snapshot to S3
```shell
h-devops ens
```
---

## How to develop
Required docker-compose version 2.22.0 or higher
```shell
docker-compose watch  
```
---

## Issue
Please open an issue [New issue](https://github.com/hoangneeee/h-devops-cli-go/issues)
