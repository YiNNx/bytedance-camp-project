echo -e "===== start docker-compose ===== \n"
sudo docker-compose up -d

echo -e "===== start api server ===== \n"
go run cmd/api/main.go &
sleep 2

echo -e "===== start services ===== \n"

cd cmd/user/
./build.sh
cd ../..
./cmd/user/output/bin/user