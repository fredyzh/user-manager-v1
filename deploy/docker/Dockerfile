FROM ubuntu:latest

WORKDIR /app

RUN mkdir -p /bin

COPY ./bin/user_manager_lnx.exe ./bin/user_manager_lnx.exe
COPY ./bin/config.yaml ./bin/config.yaml

RUN [ "chmod", "+x", "/app/bin/user_manager_lnx.exe"]

EXPOSE 7778
EXPOSE 7779
EXPOSE 17778
EXPOSE 17779

CMD ["./bin/user_manager_lnx.exe", "-conf", "./bin/config.yaml"]