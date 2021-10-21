# メモ

コンテナ ```ansible001``` に入った後、最初に コンテナ ```deb001``` の公開キーを known_hosts に入れないと駄目。

```sh
gitpod$ docker-compose exec ansible001 bash
ansible001# cd
ansible001# mkdir -m 700 .ssh
ansible001# ssh-keyscan deb001 >> .ssh/known_hosts
```

その後、 ```ansible``` 経由で ping できるか確認

```sh
ansible001# cd /app
ansible001# ansible -i hosts deb001 -m ping
deb001 | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python3"
    },
    "changed": false,
    "ping": "pong"
}
```