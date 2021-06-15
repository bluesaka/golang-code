FROM loads/alpine:3.8
LABEL maintainer="user@example.com"

ENV WORKDIR /app
ADD ./main $WORKDIR/main
WORKDIR $WORKDIR
RUN chmod +x ./main

CMD ["./main"]

