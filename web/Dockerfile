FROM golang:latest

WORKDIR /go/src

#COPY . .

ENV PATH="/go/bin:${PATH}"

RUN apt-get update

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go get github.com/go-sql-driver/mysql

RUN go get github.com/gorilla/mux

# Inicializando o módulo
#RUN go mod init celsopires/learning-golang 

#Gera o binário - somente para mandar para produção
#RUN go install 

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]