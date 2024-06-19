# Projeto de Estudos: GraphQL com Go usando gqlgen

Bem-vindo ao repositório do projeto de estudos do curso Full Cycle sobre GraphQL com Go! Este projeto tem como objetivo fornecer um ambiente prático para aprender e experimentar com a implementação de APIs GraphQL usando a linguagem Go e a biblioteca `gqlgen`.

## Índice

- [Descrição](#descrição)
- [Funcionalidades](#funcionalidades)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Descrição

Este projeto é uma aplicação básica de GraphQL implementada em Go utilizando a biblioteca `gqlgen`. Através deste projeto, você aprenderá a criar esquemas GraphQL, definir resolvers e executar queries e mutations em um servidor GraphQL.

## Funcionalidades

- Criação de um servidor GraphQL em Go.
- Definição de esquemas GraphQL.
- Implementação de resolvers para queries e mutations.
- Integração com um banco de dados simples (opcional).

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [gqlgen](https://github.com/99designs/gqlgen)
- [Docker](https://www.docker.com/) (opcional para ambiente de desenvolvimento)

## Instalação

### Pré-requisitos

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/) (opcional)

### Passos para Instalação

1. Clone este repositório:

    ```sh
    git clone https://github.com/seu-usuario/projeto-estudos-graphql-go.git
    cd projeto-estudos-graphql-go
    ```

2. Instale as dependências:

    ```sh
    go mod tidy
    ```

3. Gere o código GraphQL:

    ```sh
    go run github.com/99designs/gqlgen generate
    ```

## Como Usar

1. Execute o servidor GraphQL:

    ```sh
    go run server.go
    ```

2. Abra seu navegador e acesse `http://localhost:8080` para visualizar a interface do GraphQL Playground e começar a executar queries e mutations.

## Estrutura do Projeto

├── gqlgen.yml # Configuração do gqlgen
├── graph
│ ├── model # Definições dos modelos
│ ├── resolver.go # Implementação dos resolvers
│ ├── schema.graphqls # Definição do esquema GraphQL
├── server.go # Configuração e inicialização do servidor
├── go.mod # Dependências do Go
├── go.sum # Hashes das dependências


## Contribuição

Se você deseja contribuir com este projeto, siga os passos abaixo:

1. Faça um fork deste repositório.
2. Crie uma branch com sua feature ou correção de bug: `git checkout -b minha-feature`.
3. Faça commit das suas alterações: `git commit -m 'Minha nova feature'`.
4. Envie para o branch original: `git push origin minha-feature`.
5. Abra um pull request.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Se você tiver alguma dúvida ou sugestão, sinta-se à vontade para abrir uma issue ou entrar em contato. Boa aprendizagem e codificação! 🚀
