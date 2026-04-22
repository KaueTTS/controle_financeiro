<h1 align="center">Padrões Docker</h1>

### Tópicos

- [Principais comandos do Docker](#principais-comandos-do-docker)

### Principais comandos do Docker

- `docker compose up --build` - Construir imagem do container.

- `docker compose down` - Parar e remover containers, redes e volumes criados.

- `docker compose stop` - Parar container.

- `docker compose start` - Iniciar container.

- `docker compose restart` - Reiniciar container.

- `docker compose ps` - Listar containers que estão rodando no projeto.

- `docker compose exec <servico_container> <comando>` - Roda um comando dentro do container.

- `docker compose logs` - Mostra os logs dos containers (se colocar `-f` no final, ele mostra os logs em tempo real).

- `docker compose config` - Configuração do docker.