FROM golang

COPY /src .

ENTRYPOINT [ "go", "run" ]