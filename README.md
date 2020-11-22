# 02-tarbine-lib

Librairie partagée pour tous les microservices

C'est un module golang ; donc une libraire (pas un binaire)

Ce module est décomposé en différents packages.

Ces packages contiennent tout le boilerplate "technico-stack-fonctionnel". 
L'idée sous jacente est que chaque microservice se comporte de la même manière. Par exemple, tous les µs doivent exposer un endpoint /health, utiliser le même format de logs, ...

## Build golang

Ce projet se comporte comme un module go classique.

## Démo

### En cli
```
# Définition d'un GOPATH clean pour voir ce qu'il s'y passe
$ export GOPATH=/home/nicolas/go-02-tartine
$ sudo rm -rf $GOPATH

# Compilation du module
$ cd pkg/core/
$ cat go.mod
$ cat logger/logrus.go
$ go build ./...
go: downloading github.com/sirupsen/logrus v1.7.0
go: downloading github.com/spf13/pflag v1.0.5
go: downloading golang.org/x/sys v0.0.0-20191026070338-33540a1f6037

# Quel impact sur le GOPATH ?
$ ls $GOPATH/pkg/

```

## Vendoring or not ?

Ici pas de dossier `vendor`. Ce sont les commandes `go build`, `go run`, `go test`  qui ont
fait le taf d'aller récupérer les modules.
L'IDE a fait le même taf.
