# CIDR2HOSTSDENY

[![Gobuild Download](http://gobuild.io/badge/github.com/jbuchbinder/cidr2hostsdeny/downloads.svg)](http://gobuild.io/github.com/jbuchbinder/cidr2hostsdeny)

Convert CIDR formatted netblock lists to /etc/hosts.deny format.

## USAGE

```
Usage of ./cidr2hostsdeny:
  -services="sshd": Service(s) to block
```

cidr2hostsdeny takes a list of CIDR-formatted netblock files, and outputs
the reformatted files to stdout. It expects Unix-style EOL characters at
the moment, but will probably work just fine with DOS-style CRLF ones.

