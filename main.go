package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//Block - Declaração da estrutura básica
type Block struct {
	Timestamp     int64  //Carimbo da data/hora atual
	Data          []byte // informação valiosa real que contem o bloco ***deve ser separada aqui está junto para facilitar o entendimento
	PrevBlockHash []byte // armazena o hash do bloco anterior
	Hash          []byte // armazena o hash do bloco atual
}

/*
Então, como calculamos os hashes? A maneira como os hashes são calculados é um recurso muito importante da blockchain,
e é esse recurso que torna a blockchain segura.
O fato é que calcular um hash é uma operação computacionalmente difícil, leva algum tempo, mesmo em computadores velozes
(é por isso que as pessoas compram GPUs poderosas para minerar Bitcoin).
Esse é um projeto arquitetônico intencional, que dificulta a adição de novos blocos, impedindo sua modificação após a adição.
*/

//SetHash pegando os campos de blocos, concatena-los e calcular um hash SHA-256 na combinação concatenada
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//NewBlock seguindo uma convenção do Golang implementando uma função que simplifica a criação de um bloco
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{}}

	block.SetHash()
	return block
}

//Blockchain - strutura do bloco :)
type Blockchain struct {
	blocks []*Block
}

//AddBlock - Função que permite adicionar um blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//NewGenesisBlock - Cria um novo bloco inicial pois ainda nao temo nenhum
func NewGenesisBlock() *Block {
	return NewBlock("William Keylon", []byte{})
}

//NewBlockchain - função que cria uma blockchain com o bloco incial
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

/*func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewBlockchain()}}
}*/
func main() {
	bc := NewBlockchain()

	bc.AddBlock("Transferindo 1 BTC para Ivan")
	bc.AddBlock("Transferindo mais 2 BTC para Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Hash Anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Informação: %s\n", block.Data)
		fmt.Printf("Hash Atual: %x\n", block.Hash)
		fmt.Println()
	}
}
