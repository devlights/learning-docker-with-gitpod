# syntax=docker/dockerfile:1-labs
FROM busybox:latest

WORKDIR /workspace

# ヒアドキュメント機能は docker/dockerfile:1.3-labs から有効
RUN <<EOF
    echo 'hello' >> helloworld.txt
    echo 'world' >> helloworld.txt
EOF

CMD [ "cat", "helloworld.txt" ]