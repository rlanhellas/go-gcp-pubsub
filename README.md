# Sobre
Este projeto possui dois módulos: `pub` e `sub`. O módulo `pub` é responsável por publicar mensagens em um tópico. O módulo `sub` é responsável por assinar um tópico e receber mensagens. Tudo usando o `Google Pub/Sub`. 

# Como Usar
Primeiro você precisa fazer build de ambos os módulos: 

```go
go build -C pub
go build -C sub
```

Agora podemos produzir mensagens:

```go
./pub/pub --projectid {NOME DO SEU PROJETO GCP} --topic test --msg "minha primeira mensagem"
```

E para consumir: 

```go
./sub/sub --projectid {NOME DO SEU PROJETO GCP} --topic test --subid subtest
```

Lembre de substituir {NOME DO SEU PROJETO GCP} pelo nome do seu projeto GCP e executar os comandos acima em terminais separados para que você consiga ver as mensagens sendo produzidas e consumidas ao mesmo tempo.
