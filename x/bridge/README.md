[<img src="../../internal/images/bloxstaking_header_image.png" >](https://www.bloxstaking.com/)

<br>
<br>

# Bridge Module
The bridge module is responsible for creating, signing and managing messages from ethereum <-> Pools network.<br/>
The bridge model is heavily inspired by the Cosmos [Peggy](https://github.com/cosmos/peggy) project.

## Operator set changes
The basic security assumption for the ethereum bridge contract is that only finalized events from the pools network to the ethereum get committed. 
A committed event is an event that has 2/3 of the current pools operator set signed on it.<br/>
A key part of "counting" signatures is for the ethereum bridge to be aware of the updated current operator set, this is done by having the operator set of time X sign on the next operator set (changes) at time X+1.<br/>

The operator set update is done by initiating an operator set change request on the pools chain, then each operator signs the request. 
Once it reaches 2/3 votes it can be broadcasted to the ethereum chain, if the signatures are valid, the operator set in the contract will get updated.

## ethereum -> pools
Individual operators will index transactions sent to the ethereum bridge contract and will submit claims to the pools network. 
A claim is just an event detected on ethereum that an operator would like to commit.

Ethereum has [probabilistic finality](https://medium.com/mechanism-labs/finality-in-blockchain-consensus-d1f83c120a9a) while tendermint has an absolute (deterministic) finality as a pBFT protocol. 
This can cause issues when trying to commit evets from ethereum to a pBFT potocol as ethereum can experience forks while tendermint's finalized checkpoints can't.
To solve this issue, events from ethereum (in the forms of claims) will only be communicated to the pools network if [enough](https://github.com/ethereum/annotated-spec/blob/master/phase0/beacon-chain.md#misc) confirmations have occured.<br/>
The eth2 beacon chain solves this issue by having validators "attest" to blocks including events (deposits) from eth1, including an event in a block is made possible because the beacon chain finalizes (happy flow) every 32 blocks (epoch). 
Tendermint finalizes every epoch, why not just include a claim in a block and wait for it to finalize? That could halt the tendermint chain in case of an ethereum fork, we need to simulate the voting mechanism as in eth2. 
  

Each operator submits its locally seen events as claims, each claim is uniquely identified.
Internally, a node that saw votes on claim X will aggregate the power of the operators that voted on it.
Claims that receive 2/3 votes (out of the operator set) will get committed.

## pools -> ethereum
Operators broadcast transactions they want to relay to ethereum, anyone can submit such a transaction which gets included in a block.<br/>
Anyone can also broadcast a request batch transaction which builds a batch of transactions for the operator set to sign.  
The blockchain now becomes a coordination platform for creating batches and collecting individual signatures from operators.<br\>
A batch that collected 2/3 signatures from the operator set can be relayed to ethereum and processed.



