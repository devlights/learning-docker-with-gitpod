mkdir -m 700 $HOME/.ssh
ssh-keyscan demo001 >> $HOME/.ssh/known_hosts

tail -f /dev/null
