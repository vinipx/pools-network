[<img src="../../internal/images/bloxstaking_header_image.png" >](https://www.bloxstaking.com/)

<br>
<br>

# Bridge Module
The bridge module is reponsible for creating, signing and managing messages from ethereum <-> Pools network.

## ethereum -> pools
Individual operators will index transactions sent to the bridge ethereum smart contract and will submit claims to the pools network. 
A claim is just an event detected on ethereum that an operator would like to commit.

Ethereum has [probabilistic finality](https://medium.com/mechanism-labs/finality-in-blockchain-consensus-d1f83c120a9a) while tendermint has an absolute (deterministic) finality as a pBFT potocol. This can cause issues when trying to commit evets from ethereum to a pBFT potocol as ethereum can experience forks while tendermint's finalized checkpoints can't.
To solve this issue, events from ethereum (in the forms of claims) will only be communicated to the pools network if [enough](https://github.com/ethereum/annotated-spec/blob/master/phase0/beacon-chain.md#misc) confirmations have occured.  

Each operator in the operator set wil submit its locally seen events as claims, each claim is uniquely identified.
Claims that receive 2/3 votes (out of the operator set) will get committed.

There are 4 claim types:
- Delegate
- Undelegate
- Create pool
- Create operator

