module github.com/assetsadapterstore/abbc-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/eosio-adapter v1.1.0
	github.com/blocktree/go-owcdrivers v1.1.22
	github.com/blocktree/openwallet v1.5.5
	github.com/eoscanada/eos-go v0.8.16
)

//replace github.com/blocktree/eosio-adapter => ../eosio-adapter
