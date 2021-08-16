solc --optimize --abi ./contracts/forum.sol -o contracts/build
solc --optimize --bin ./contracts/forum.sol -o contracts/build

abigen --abi=contracts/build/Forum.abi --bin=contracts/build/Forum.bin --pkg=forum --out=./contracts/Forum.go
   Need to update the package name

### Template model
https://www.bootdey.com/snippets/view/bs4-forum
### Auto-refresh
go get github.com/pilu/fresh
fresh