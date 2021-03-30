package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
	Nonce    int
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{prevHash, []byte{}, []byte(data), 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("GenesisBlock", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
