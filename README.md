# Template Prototyping

Este projeto é uma ferramenta simples em Go que gera um PDF a partir de um arquivo JSON e um template HTML. Ele é útil para criar extratos ou relatórios de forma dinâmica.

## Pré-requisitos

Antes de executar o projeto, certifique-se de que você tem o seguinte instalado:

- [Go](https://golang.org/dl/) (versão 1.16 ou superior)
- [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html) - Uma ferramenta para converter HTML em PDF.
- [Make](https://www.gnu.org/software/make/) - Para usar o Makefile.

### Instalando dependencias

```
$ make install-deps
```

### Executando projeto

```
$ make clean run ; xdg-open output.pdf
```
