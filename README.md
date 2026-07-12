<h1 align="center">Controle Financeiro</h1>

<p align="center">
<img loading="lazy" src="http://img.shields.io/static/v1?label=STATUS&message=CONCLUIDO&color=green&style=for-the-badge"/>
</p>

> [!IMPORTANT]
> *Esse projeto está concluído.*

### Tópicos

- [Descrição do projeto](#descrição-do-projeto)
  - [Funcionalidades Principais](#funcionalidades-principais)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Projeto em funcionamento](#projeto-em-funcionamento)
- [Como utilizar o projeto](#como-utilizar-o-projeto)
- [Colaboradores](#colaboradores)

## Descrição do projeto

O **Controle Financeiro** é uma aplicação Full Stack desenvolvida para auxiliar usuários no gerenciamento de receitas e despesas pessoais.

O sistema oferece uma interface simples e intuitiva para cadastro de movimentações financeiras, além de disponibilizar um resumo em tempo real do fluxo de caixa, permitindo um acompanhamento eficiente da saúde financeira.

O projeto foi desenvolvido com foco em boas práticas de arquitetura, separação de responsabilidades, testes automatizados e documentação da API.

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

***Back-End:***
```
src
├── api
│   ├── routes
│   └── v1
│       ├── controllers
│       ├── dto
│       ├── validators
│       └── responses
├── services
├── repositories
├── models
├── config
└── shared
```

| Camada     | Responsabilidade                                                                       |
| ---------- | -------------------------------------------------------------------------------------- |
| Routes     | Define os endpoints da API e realiza o mapeamento das rotas                            |
| Controller | Recebe as requisições HTTP, valida a entrada e delega o processamento para os serviços |
| DTO        | Objetos de entrada e saída de dados da API                                             |
| Validators | Centraliza as regras de validação das requisições                                      |
| Responses  | Padroniza as respostas retornadas pela API                                             |
| Service    | Implementa regras de negócio da aplicação                                              |
| Repository | Responsável pelo acesso e persistência dos dados no banco                              |
| Models     | Representa as entidades do domínio                                                     |
| Config     | Configurações da aplicação, banco de dados e variáveis de ambiente                     |
| Shared     | Código compartilhado entre módulos, como utilitários, constantes, middlewares e helpers|

<br>

***Front-End:***
```
src
├── api
├── assets
├── components
├── hooks
├── pages
├── styles
├── types
├── utils
└── main.tsx
```

| Camada     | Responsabilidade                                                      |
| ---------- | --------------------------------------------------------------------- |
| API        | Comunicação com a API do backend                                      |
| Assets     | Imagens, ícones e arquivos estáticos                                  |
| Components | Componentes reutilizáveis da interface                                |
| Hooks      | Hooks customizados utilizados pela aplicação                          |
| Pages      | Páginas que representam as telas do sistema                           |
| Styles     | Estilos globais e específicos da aplicação                            |
| Types      | Tipagens compartilhadas do TypeScript                                 |
| Utils      | Funções utilitárias reutilizadas em diferentes partes do projeto      |

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