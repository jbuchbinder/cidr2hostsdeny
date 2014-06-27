# CIDR2HOSTSDENY

Convert CIDR formatted netblock lists to /etc/hosts.deny format.

## USAGE

```
Usage of ./cidr2hostsdeny:
  -services="sshd": Service(s) to block
```

cidr2hostsdeny takes a list of CIDR-formatted netblock files, and outputs
the reformatted files to stdout. It expects Unix-style EOL characters at
the moment, but will probably work just fine with DOS-style CRLF ones.

