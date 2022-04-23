

## Exemplo Utilizando gRPC 

    - Exemplos Basicos utilizando gRPC + protobuf 
    - Exemplo na Linguagem Go (google)

## Configuração do ambiente de desenvolvimento

    - Utilizado Visual Studio Code 
        - Plugin Remote Container (https://code.visualstudio.com/docs/remote/containers) 
    
    - go version go1.17.8 linux/amd64
    - Depencias do go está no arquivo go.mod
    - protoc
  
## Comandos

    - Rodar Servidor
        go run cmd/client/server.go
    - Rodar Cliente
        go run cmd/client/client.go
    - Gerar stubs
        protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

## Estrutura proto

    - Usuario (id, nome, email)
        - interface de implementação para adicionar usuario(AddUser())

## Pendencias

    - poder buildar protoc no ambiente vscode 
      - sudo apt install protobuf-compiler-grpc
    - 
# Referencias 

    - [https://grpc.io/](https://grpc.io/)
    - [https://developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers)
