# Issue 1: missing blst_src from vendored github.com/onflow/crypto
Running `go build` produces:
```
$ go build
# github.com/onflow/crypto
In file included from vendor/github.com/onflow/crypto/bls12381_utils.h:26,
                 from vendor/github.com/onflow/crypto/bls_include.h:24,
                 from vendor/github.com/onflow/crypto/bls.go:44:
vendor/github.com/onflow/crypto/blst_include.h:5:10: fatal error: consts.h: No such file or directory
    5 | #include "consts.h"
      |          ^~~~~~~~~~
compilation terminated.
```

This seems to be because `blst_src` directory is missing from the vendored deps.

Trying to manually copy this directory to vendor/github.com/onflow/crypto works fine.
However, when trying to compile the code as a Go plugin, there's this issue:

```
$ go build --mod=vendor --trimpath --buildmode=plugin -o test.o
# flow-go-plugin
/home/syhpoon/UNIX/soft/go/pkg/tool/linux_amd64/link: running gcc failed: exit status 1
/usr/bin/ld: /tmp/go-link-3804225422/000013.o: warning: relocation against `__blst_platform_cap' in read-only section `.text'
/usr/bin/ld: /tmp/go-link-3804225422/000013.o: relocation R_X86_64_PC32 against symbol `__blst_platform_cap' can not be used when making a shared object; recompile with -fPIC
/usr/bin/ld: final link failed: bad value
collect2: error: ld returned 1 exit status
```
