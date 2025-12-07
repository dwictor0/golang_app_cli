# go_app - CLI

## Flags
- ip: Obtém os endereços IP de um domínio ou servidor.
- server: Lista os servidores/domínios associados.
- help: Exibe a lista de comandos disponíveis e opções de ajuda.

## Instalação
Clone o repositório e compile o projeto:

git clone https://github.com/seu-usuario/go_app.git<br>
cd go_app<br>
go build -o go_app main.go

## Uso
Sintaxe geral:
go_app [command] --flags (host) 

### Exemplo obter IP:
go_app ip --host google.com

### Exemplo listar servidores/domínios:
go_app server --host example.com
