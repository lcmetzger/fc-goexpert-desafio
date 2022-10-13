# Descrição do Desafio Full Cycle - Go Expert

Olá dev, tudo bem?

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

Você precisará nos entregar dois sistemas em Go:

- **`client.go`**
- **`server.go`**

## Os requisitos para cumprir este desafio são:

- O **`client.go`** deverá realizar uma requisição HTTP no **`server.go`** solicitando a cotação do dólar.
- O **`server.go`** deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
- Usando o package "context", o **`server.go`** deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
- O **`client.go`** precisará receber do **`server.go`** apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o **`client.go`** terá um timeout máximo de 300ms para receber o resultado do **`server.go`**.
- O **`client.go`** terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
- O endpoint necessário gerado pelo **`server.go`** para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

Ao finalizar, envie o link do repositório para correção.

# Organização dos Requisitos e Restrições

## `server.go`

- ### Requisitos funcionais

  - Requisitar os dados de cotação no site remoto (vide link acima);
  - Deve retornar os dados em formato JSON;
  - Deve enviar somente o valor atual do cambio (campo bid do retorno);
  - Gravar a cotação em tabela no banco de dados SQLite;

- ### Restrições
  - Requisição para obter a cotação no site não pode ultrapassar os 200ms;
  - Persistência no banco de dados deve ser de até 10ms;

## `client.go`

- ### Requisitos funcionais

  - Deve fazer uma requisição ao endereço: `http://localhost:8080/cotacao`, a fim de obter os dados;
  - Deverá persistir a cotação atual em um arquivo chamado "cotacao.txt" contendo o seguinte formato: "Dólar: {valor}";

- ### Restrições
  - Requisição para o Server não deve ultrapassar os 300ms

# Detalhes sobre a implementação

Nos requisitos do desafio, não constavam detalhes como por exemplo:

- nome e estrutura de tabela;
  - assumido o nome da tabela como cambio com os campos data e cotacao
- estrutura do json retornado pelo server
  - assumido que será retornado a chave "bid" com o valor da cotação
- se deveria sempre estar inserindo os dados na tabela e cada requisição, ou atualizando os dados existentes
  - não será verificado se o registro já existe na base, e sempre será inserido um novo registro
  - a data será a data em que foi executado a inclusão do registro
- qual deveria ser o indice primário da tabela
  - não foi estabelecido índice na tabela
