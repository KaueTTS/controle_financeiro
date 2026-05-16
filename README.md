<h1 align="center">Controle Financeiro</h1>

<p align="center">
<img loading="lazy" src="https://img.shields.io/static/v1?label=STATUS&message=EM%20ANDAMENTO&color=blue&style=for-the-badge"/>
</p>

> [!IMPORTANT]
> *Esse projeto está em andamento.*

### Tópicos

- [Descrição do projeto](#descrição-do-projeto)
  - [Funcionalidades Principais](#funcionalidades-principais)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Projeto em funcionamento](#projeto-em-funcionamento)
- [Como utilizar o projeto](#como-utilizar-o-projeto)
- [Colaboradores](#colaboradores)

## Descrição do projeto

O Controle Financeiro é uma aplicação web para gerenciamento de receitas e despesas pessoais.

A aplicação permite que usuários registrem, visualizem e analisem suas transações financeiras, facilitando o controle do fluxo de caixa e a tomada de decisões financeiras.

### Funcionalidades Principais
> ***Listar transações:*** Visualização completa das movimentações financeiras cadastradas. <br>
> ***Criar transações:*** Cadastro de receitas e despesas. <br>
> ***Deletar transações:*** Remoção de transações cadastradas. <br>
> ***Editar transações:*** Atualização de informações das transações. <br>
> ***Resumo financeiro:*** Cálculo automático de entradas, saídas e saldo total. <br>
> ***Paginação e filtros:*** Busca e navegação otimizada entre transações.

O design do site é responsivo e intuitivo, permitindo que usuários de qualquer dispositivo acessem as informações de maneira rápida e eficiente.

## Tecnologias

<details closed>
<summary>Front-End</summary>
  <div width="140px">
      <img src="https://skillicons.dev/icons?i=react,css,typescript" />
  </div>
</details>

<details closed>
<summary>Back-End</summary>
  <div width="140px">
      <img src="https://skillicons.dev/icons?i=go,sqlite" />
  </div>
</details>

<details closed>
<summary>Infra</summary>
  <div width="140px">
      <img src="https://skillicons.dev/icons?i=docker" />
  </div>
</details>

<details closed>
<summary>Ferramentas</summary>
  <div width="140px">
      <img src="https://skillicons.dev/icons?i=vscode,vite" />
  </div>
</details>

## Arquitetura

Back-End:

- Controller -> recebe requisições HTTP 
- Service -> regras de negócio
- Repository -> acesso a dados e APIs
- DTOs -> transporte de dados entre camadas
- Models -> representação das entidades
- Config ->

Front-End:

- Components -> componentes reutilizáveis
- Hooks -> 
- Styles -> estilizações


## Projeto em funcionamento

Clique na imagem abaixo para assistir ao tutorial em vídeo!

[![Assista ao tutorial](image.png "Como utilizar esse projeto na sua máquina")](semvideo.com)

**Descrição**: Este vídeo cobre todo o processo para visualizar o projeto em funcionamento, do início ao fim.

## Como utilizar o projeto

```
< INSTALADORES >

Back-End:
cd ./server
go mod tidy

Front-End:
cd ./web
npm install


< INICIADORES >

docker compose up --build


< TESTES DE COVERAGE >

No git bash, rode:

go test ./... -v -coverprofile=coverage.out
go tool cover -func=coverage.out

```

## Colaboradores

| [<img src="https://avatars.githubusercontent.com/u/69527468?v=4" width=115><br><sub>Kauê Bertaze de Oliveira</sub>](https://github.com/KaueTTS)<br><sub>Developer Full Stack</sub> |
| :---: