

export GO111MODULE=on
go get -u -v github.com/dgraph-io/dgo/v2
  


WORKING:
--------
curl -H "Content-Type: application/rdf" -X POST localhost:8080/mutate?commitNow=true -d $'
{set {
  <2993> <genre.name> "Test Genre" .
  <2993> <genre.id> "2993" .
}}' | jq


http://localhost:8089/movierecom/0/100/0





docker run:
-----------
	cd /data/GeekBooks/dbs/dgraphdocs
	sudo docker-compose up -d


	EXPORT:  
	curl localhost:8080/admin/export
	Automatic export folder created in alpha container.
	
	container login command
	sudo docker exec -it -u 0 dgraphdocs_alpha_1 bash

	sudo docker cp dgraphdocs_alpha_1:/dgraph/export/dgraph.r44.u0320.0721/ .
	
	sudo apt install docker-compose
	sudo apt install docker.io
	
	
	sudo apt install docker-compose
	sudo apt install docker.io
	sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
	sudo chmod +x /usr/local/bin/docker-compose
	docker-compose --version
	docker-compose up -d
	sudo /etc/init.d/docker restart
	docker-compose up -d
	sudo netstat -ntpl
	sudo /etc/init.d/docker stop
	sudo /etc/init.d/docker status
	sudo /etc/init.d/docker start
	sudo /etc/init.d/docker status
	sudo /etc/init.d/docker status -l
	sudo docker-compose up -d
	sudo docker ps -a
