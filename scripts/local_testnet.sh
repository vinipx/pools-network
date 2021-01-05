#!/bin/sh

working_dir=./.tmp_local_testnet
binary=$working_dir/build/pools
testnet_name="local_test_net"
testnet_moniker="testnet"

cleanup() {
  #  key dir
  rm -r $working_dir
}

init() {
  mkdir $working_dir

  # build
  go build -o $binary github.com/bloxapp/pools-network/cmd/pools-networkd
  #init chain
  $binary init --chain-id=$testnet_name $testnet_moniker --home $working_dir
}

generate_keys() {
  $binary unsafe-reset-all
  params="--keyring-backend test --keyring-dir $working_dir/keys"

  for i in 1 2 3 4
  do
    val_name="validator_$i"
    $binary keys add $val_name $params
    $binary add-genesis-account $($binary keys show $params $val_name -a) 1000000000stake  --home $working_dir

    # gentx
    mkdir $working_dir/config/gentx
    pubkey=$($binary keys show --keyring-backend test --keyring-dir $working_dir/keys $val_name --pubkey --bech cons)
    echo $pubkey
    $binary gentx $val_name --pubkey $pubkey --chain-id=$testnet_name --home $working_dir --output-document $working_dir/config/gentx/gentx_$val_name.json $params
  done

  $binary collect-gentxs --home $working_dir
}

#cleanup
#init
#generate_keys

$binary start --home $working_dir
