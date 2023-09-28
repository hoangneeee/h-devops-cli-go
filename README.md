# H-devops

### How to install
```shell
mv bin/h-devops /usr/local/bin/h-devops
chmod a+x /usr/local/bin/h-devops
```

---
### Feature
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
---

### How to develop
Required docker-compose version 2.22.0 or higher
```shell
docker-compose watch  
```
