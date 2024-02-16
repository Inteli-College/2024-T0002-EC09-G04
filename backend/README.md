# Backend

O código contido nesse recorte do projeto permite a criação de um sistema de publisher e subscriber utilizando um broker local para isso. O objetivo aqui é simular a emissão de dados por parte dos sensores ( publisher ) e a recepção destes por um agregador ( subscriber ).

> [!IMPORTANT]
> Antes de seguir os comando abaixo, confira se a sua máquina contém os requerimentos [necessários](https://docs.docker.com/desktop/install/ubuntu/). Para uma melhor interação entre no workspace do diretório backend.

## Para rodar todo o sistema:

```bash
docker compose up
```