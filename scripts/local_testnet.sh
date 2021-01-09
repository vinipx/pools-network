#!/bin/sh

working_dir=./.tmp_local_testnet
binary=$working_dir/build/pools
testnet_name="local_test_net"
generated_nodes=6

cleanup() {
  #  key dir
  rm -r $working_dir
}

build() {
  # install dependencies
  if [[ "$OSTYPE" =~ ^darwin ]]; then
    for dependency in tmux reattach-to-user-namespace; do
        if brew ls --versions $dependency > /dev/null; then
          echo "$dependency instelled, moving on..."
        else
          echo "Installing $dependency"
          brew install $dependency
        fi
    done
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
  echo "node$idx"
}

generate_app_folders() {
  $binary testnet  --v $generated_nodes --output-dir ./.tmp_local_testnet --chain-id $testnet_name
}

kill_on_press() {
  read
  tmux kill-session
}

run() {
  cmd="tmux new-session \"$binary start --home $working_dir/$(node_name 0)/poolsd\" \;"
  for (( i=1; i<$generated_nodes; i++ ))
  do
    cmd="${cmd} split-window \"$binary start --home $working_dir/$(node_name $i)/poolsd\" \;"
  done
  cmd="${cmd} select-layout tiled"
  echo $cmd
  eval $cmd
}

cleanup
build
generate_app_folders
run
