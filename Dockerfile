FROM golang:1.10

# Open container port
EXPOSE 9000

WORKDIR /go/src/app
COPY QuestionBankAPI /go/src/app/
COPY . /go/src/app/
RUN chmod +x /go/src/app/QuestionBankAPI

ENV QB_REDISHOST redis
ENV QB_REDISPORT 6379
ENV QB_DBUSERNAME user
ENV QB_DBPASSWORD pass
ENV QB_DBDEFAULTHOST mysql
ENV QB_DBDEFAULTPORT 3306
ENV QB_DBNAME db

ENTRYPOINT ["./QuestionBankAPI"]
