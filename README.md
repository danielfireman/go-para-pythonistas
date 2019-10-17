# Markdown Server Python

Esse é um tutorial que visa apresentar coinceitos básicos da linguagem de programação [Go](https://golang.org/) para quem já sabe [Python](https://www.python.org/). O objetivo é ser "mão-na-massa", ou seja, novos conceitos serão apresentados num contexto específico de implementação. Nessa versão do tutorial, implementaremos um servidor web que renderiza páginas escritas em [Markdown](https://en.wikipedia.org/wiki/Markdown). Será uma espécie de [Wiki](https://pt.wikipedia.org/wiki/Wiki), porém com a sintaxe Markdown.

As funcionalidades do servidor serão as seguintes:

- Nova página (`/edita/:nome_pagina`): abre um formulário que permite que a pessoa usuária insira ou edite o conteúdo Markdown da página especificada
- Remover página (`/remove/:nome_pagina`): Remove a página especificada
- Renderizar página (`/:página`): Dado o nome da página na URL o servidor deve renderizá-la
- Listar páginas (`/lista`): Lista as páginas armazenadas

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
cd go-para-pythonistas
$GOPATH/bin/present
```

Pronto, agora é só apontar o navegador para https://127.0.0.1:3999.

## Mais materiais sobre Go

### Em português

- [Um tour por Go](https://go-tour-br.appspot.com/welcome/1)
- [Go efetivo](http://www.golangbr.org/doc/go_efetivo)

### Em inglês

- [Guia da Google sobre Revisões de Código](https://github.com/golang/go/wiki/CodeReviewComments)
- [Guia da Uber](https://github.com/uber-go/guide/blob/master/style.md)

Conhece alguma material bacana que não está aqui? Ficaremos muito felizes com sua contribuição/PR!
