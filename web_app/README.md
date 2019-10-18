# Servidor REST de Lista de Coisas a Fazer (TODO List)

Esse é um tutorial que visa apresentar coinceitos básicos da linguagem de programação [Go](https://golang.org/) para quem já sabe [Python](https://www.python.org/). O objetivo é ser "mão-na-massa", ou seja, novos conceitos serão apresentados num contexto específico de implementação. Nessa versão do tutorial, implementaremos um servidor REST que provê uma API para uma lista de coisas a fazer (TODO list)

## Dependências

- Conta no [Github](github.com)
- [VSCode](https://code.visualstudio.com/)
- [Go](http://www.golangbr.org/doc/instalacao)
- [Plugin Go para VSCode](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

A IDE+Plugin e a conta no github são opcionais, porém recomendadas para facilitar o acompanhamento do tutorial.

## Apresentação

Para iniciar a apresentação digite o comando abaixo:

```sh
go get -v golang.org/x/tools/cmd/present
cd go-para-pythonistas/web_app
$GOPATH/bin/present
```

Pronto, agora é só apontar o navegador para https://127.0.0.1:3999.