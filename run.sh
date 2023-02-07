echo -e "===== start docker-compose ===== \n\n"
sudo docker-compose up -d

echo -e "===== start api server ===== \n\n"
go run cmd/api/main.go &
sleep 2

echo -e "===== start services ===== \n\n"
cd cmd/user/
./build.sh
cd ../..
./cmd/user/output/bin/user

# for file in *
# do
#     if [ "${file}" != "api" ];
#     then
#         cd ${file}/
#         ./build.sh
#         cd ../..
#         ./cmd/${file}/output/bin/${file}
#     fi
# done