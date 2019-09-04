module github.com/assetsadapterstore/abbc-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.11.1
	github.com/blocktree/eosio-adapter v1.0.14
	github.com/blocktree/go-owcdrivers v1.0.33
	github.com/blocktree/openwallet v1.4.5
	github.com/eoscanada/eos-go v0.8.11
)

//replace github.com/blocktree/eosio-adapter => ../eosio-adapter
