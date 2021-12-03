# syntax=docker/dockerfile:1-labs
FROM busybox:latest

WORKDIR /workspace

# ヒアドキュメント機能は docker/dockerfile:1.3-labs から有効
RUN <<EOF
    echo 'hello' >> hello.txt
    echo 'world' >> world.txt
    cat hello.txt world.txt > helloworld.txt
EOF

CMD [ "cat", "helloworld.txt" ]