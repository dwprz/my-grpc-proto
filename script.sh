# initialisation git
git init
git branch 
git remote add origin https://github.com/dwprz/my-grpc-proto.git
git commit -am "Initial"  # digunakan untuk membuat commit yang mencakup semua perubahan pada file yang sudah dilacak dengan pesan komit "Initial".
git push -u origin master
git checkout -b "feat-hello-proto"

# create source code go proto
protoc --go_opt=module=github.com/dwprz/my-grpc-proto --go_out=. --go-grpc_opt=module=github.com/dwprz/my-grpc-proto --go-grpc_out=. ./proto/hello/*.proto ./proto/payment/*.proto ./proto/transaction/*.proto ./proto/bank/*.proto
