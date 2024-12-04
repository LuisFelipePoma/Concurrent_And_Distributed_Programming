#IMAGEN BASE
FROM golang:alpine
#Copiado del fuente a la nueva imagen
COPY ./codeRedP2PHP.go ./codenodehp-003.go
#exponer los puertos
EXPOSE 9002
#poner en ejecución el algoritmo
CMD ["go","run","codenodehp-003.go"]