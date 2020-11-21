# 02-tarbine-lib

Librairie partagée pour tous les microservices

C'est un module golang ; donc une libraire (pas un binaire)

Ce module est décomposé en différents packages.

Ces packages contiennent tout le boilerplate "technico-stack-fonctionnel". 
L'idée sous jacente est que chaque microservice se comporte de la même manière. Par exemple, tous les µs doivent exposer un endpoint /health, utiliser le même format de logs, ...

## Build golang

Ce projet se comporte comme un module go classique.
