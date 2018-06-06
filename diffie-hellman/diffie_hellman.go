package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey doc here
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	k, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, two))
	return k.Add(k, two)
}

// PublicKey doc here
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair doc here
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey doc here
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
