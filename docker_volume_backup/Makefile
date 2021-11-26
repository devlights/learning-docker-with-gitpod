default: \
	create-backup-target-volume \
	write-backup-volume-data \
	show-backup-volume-data \
	backup-volume-data \
	confirm-backup-contents \
	remove-backup-target-volume \
	create-restore-target-volume \
	restore-volume-data \
	show-restore-volume-data \
	remove-restore-target-volume

create-backup-target-volume:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップ対象 Docker ボリューム作成'
	@echo '+----------------------------------------------+'
	docker volume create vol-backup-target

create-restore-target-volume:
	@echo '+----------------------------------------------+'
	@echo '+ レストア対象 Docker ボリューム作成'
	@echo '+----------------------------------------------+'
	docker volume create vol-restore-target

remove-backup-target-volume:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップ対象 Docker ボリューム削除'
	@echo '+----------------------------------------------+'
	docker volume rm vol-backup-target

remove-restore-target-volume:
	@echo '+----------------------------------------------+'
	@echo '+ レストア対象 Docker ボリューム削除'
	@echo '+----------------------------------------------+'
	docker volume rm vol-restore-target

show-backup-volume-data:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップ対象 Docker ボリューム データ確認'
	@echo '+----------------------------------------------+'
	docker container run \
		--init \
		--rm \
		--mount type=volume,src=vol-backup-target,dst=/data \
		--workdir /data \
		busybox \
		sh -c 'pwd ; ls -lh'

show-restore-volume-data:
	@echo '+----------------------------------------------+'
	@echo '+ レストア対象 Docker ボリューム データ確認'
	@echo '+----------------------------------------------+'
	docker container run \
		--init \
		--rm \
		--mount type=volume,src=vol-restore-target,dst=/data \
		--workdir /data \
		busybox \
		sh -c 'pwd ; ls -lh'

write-backup-volume-data:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップ対象 Docker ボリューム データ書き込み'
	@echo '+----------------------------------------------+'
	docker container run \
		--init \
		--rm \
		--mount type=volume,src=vol-backup-target,dst=/data \
		--workdir /data \
		busybox \
		sh -c 'pwd ; echo "test data" > test.txt'

backup-volume-data:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップ対象 Docker ボリューム バックアップ'
	@echo '+----------------------------------------------+'
	sudo rm -rf ./backup.tar.gz
	docker container run \
		--init \
		--rm \
		--mount type=volume,src=vol-backup-target,dst=/data \
		--mount type=bind,src=${PWD},dst=/backup \
		--workdir / \
		busybox \
		sh -c 'tar czf /backup/backup.tar.gz /data'

confirm-backup-contents:
	@echo '+----------------------------------------------+'
	@echo '+ バックアップファイルの中身を確認'
	@echo '+----------------------------------------------+'
	sudo tar ztf ${PWD}/backup.tar.gz

restore-volume-data:
	@echo '+----------------------------------------------+'
	@echo '+ レストア対象 Docker ボリューム レストア'
	@echo '+----------------------------------------------+'
	docker container run \
		--init \
		--rm \
		--mount type=volume,src=vol-restore-target,dst=/data \
		--mount type=bind,src=${PWD},dst=/backup \
		--workdir / \
		busybox \
		sh -c 'tar zxf /backup/backup.tar.gz'
