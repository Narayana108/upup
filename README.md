# UPUP
### Ubuntu Package Updates Parser

Depends on ansible playbook `check_updates.yml` from https://github.com/narayana-das/ansible

## Quick start:
Run the playbook `check_updates.yml`
```bash
ansible-playbook -f 20 --ask-vault-pass -i hosts_local check_updates.yml
```

Run script:
```bash
go run main.go
```

Create executable binary:
```bash
go build
```
