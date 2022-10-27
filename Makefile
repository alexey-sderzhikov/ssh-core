db-up:
	docker start mysql

db-con:
	mysql --host 172.17.0.2 -u root -p

.PHONY:
	db-up