# Backend

O código contido neste recorte do projeto representa o sistema de simulação e, futuramente, também abrangerá as APIs requeridas pelo sistema. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1].

## Dependências:

Antes de continuar, é necessário instalar as dependências para a execução dos comandos abaixo. Acesse o [link](https://docs.docker.com/desktop/install/ubuntu/).

## Como rodar o sistema:

Abaixo estão as possíveis interações e as instruções de como realizá-las.

#### Rodar testes:

Aqui, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L7).

###### Comando:

```shell
make test
```

###### Output:

```shell
[+] Running 1/1
 ✔ Container backend-broker-1  Started                                                                                                          0.0s 
Running the tests
?       github.com/Inteli-College/2024-T0002-EC09-G04/cmd/api-test      [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/cmd/simulation    [no test files]
ok      github.com/Inteli-College/2024-T0002-EC09-G04/internal/gas      0.003s  coverage: 100.0% of statements
ok      github.com/Inteli-College/2024-T0002-EC09-G04/internal/rad_lum  0.003s  coverage: 100.0% of statements
ok      github.com/Inteli-College/2024-T0002-EC09-G04/pkg/station       2.011s  coverage: 88.5% of statements
[+] Running 2/2
 ✔ Container backend-broker-1  Removed                                                                                                          0.2s 
 ✔ Network backend_default     Removed  
```

> [!NOTE]
> - No meio do processo, é necessário subir um broker local para realizar os testes de transmissão de mensagens entre os tópicos.

#### Rodar a simulação:

Mais uma vez, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L15C2-L15C7).

###### Comando:

```bash
make run
```

###### Output:

```shell
broker-1      | 1708301841: New client connected from 172.22.0.3:39954 as station-2805345190559206082 (p2, c1, k30).
broker-1      | 1708301841: New connection from 172.22.0.4:51230 on port 1891.
broker-1      | 1708301841: New client connected from 172.22.0.4:51230 as subscriber (p2, c1, k30).
api-test-1    | Received: {"location":"{\"latitude\":36.000000,\"longitude\":16.000000}","gas":"{\"CO2\":526.000000,\"CO\":7.000000,\"NO2\":343.000000,\"MP10\":404.000000,\"MP25\":179.000000}","rad_lum":"{\"ET\":267.000000,\"LI\":3.000000,\"SR\":241.000000}","timestamp":"2024-02-19 00:17:21.311444485 +0000 UTC m=+0.004634198"} from: /stations
```

> [!NOTE]
>  - Este comando está subindo todos os serviços presentes no arquivo compose.yml. São eles, o broker local, a simulação e a api-test que está sendo usada, por hora apenas para mostrar o log do que está sendo transmitido pela simulação.

#### Rodar a visualização da cobertura de testes:

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L21).

###### Comando:

```bash
make coverage 
```

###### Output:
![output_coverage](https://github.com/Inteli-College/2024-T0002-EC09-G04/assets/89201795/59e8654d-26bc-4e6c-990a-d4c823f38973)

> [!NOTE]
>  - Este comando está criando, a partir do arquivo `coverage_sheet.md`, uma visualização da cobertura de testes nos principais arquivos Go.

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.
