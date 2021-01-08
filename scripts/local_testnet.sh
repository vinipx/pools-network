#!/bin/sh

working_dir=./.tmp_local_testnet
binary=$working_dir/build/pools
testnet_name="local_test_net"
generated_nodes=4

cleanup() {
  #  key dir
  rm -r $working_dir
}

build() {
  # install dependencies
  if [[ "$OSTYPE" =~ ^darwin ]]; then
    echo "Installing tmux"
    brew install tmux
  else
    echo "Supports only OSX"
    exit
  fi

  # make temp dir for testnets
  mkdir $working_dir

  # build
  echo "building pools"
  go build -o $binary github.com/bloxapp/pools-network/cmd/pools-networkd
}

node_name() {
  idx=$1
  echo "node_$idx"
}

generate_app_folders() {
  for ((i=1 ; i <= generated_nodes ; ++i))
  do
    node_name=$(node_name $i)
    node_folder=$working_dir/$node_name
    $binary init node_name --chain-id=$testnet_name --home $node_folder
  done
}

generate_state() {
  ref_node_folder=$working_dir/ref_node
  $binary init ref_node --chain-id=$testnet_name --home $ref_node_folder
  params="--keyring-backend test --keyring-dir $ref_node_folder/keys"

  for ((i=1 ; i <= generated_nodes ; ++i))
  do
    val_name="validator_$i"
    $binary keys add $val_name $params
    $binary add-genesis-account $($binary keys show $params $val_name -a) 1000000000stake  --home $ref_node_folder

    # gentx
    mkdir $ref_node_folder/config/gentx
    pubkey=$($binary tendermint show-validator --home $working_dir/$(node_name $i))
    ip="127.0.0.1:6000$i"
    $binary gentx $val_name --ip $ip --pubkey $pubkey --chain-id=$testnet_name --home $ref_node_folder --output-document $ref_node_folder/config/gentx/gentx_$val_name.json $params
  done

  $binary collect-gentxs --home $ref_node_folder
}

copy_state_to_nodes() {
  for ((i=1 ; i <= generated_nodes ; ++i))
  do
    source=$working_dir/ref_node/config/genesis.json
    dest=$working_dir/$(node_name $i)/config
    cp $source $dest
  done
}

run() {
  tmux \
    new-session  "$binary start --p2p.laddr 0.0.0.0:25555 --home $working_dir/$(node_name 1)" \; \
    split-window "$binary start --p2p.laddr 0.0.0.0:25556 --home $working_dir/$(node_name 2)" \; \
    split-window "$binary start --p2p.laddr 0.0.0.0:25557 --home $working_dir/$(node_name 3)" \; \
    split-window "$binary start --p2p.laddr 0.0.0.0:25558 --home $working_dir/$(node_name 4)" \; \
    select-layout tiled
}

cleanup
build
generate_app_folders
generate_state
copy_state_to_nodes

#run
$binary gentx --help