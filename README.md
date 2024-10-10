# sCram-go

SCAM SHA256 Password hashing using stdlib

Almost first iteration. Needs more love :)

## Usage

```shell
./sCram-go <password>

Missing password to hash.
Usage: sCram-go <password>
exit status 1
```

with a given password:

```shell
./sCram-go test                                                                
Hashed Password:
SCRAM-SHA-256$4096:HlMwkDOMAVteCQN....
```