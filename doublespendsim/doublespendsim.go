package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
)

// A global nonce used to seed the rng.
var nonce uint64

// Given a probability of success, output a bool indicating whether or not
// success was achieved.
func tryHash(num, denom int) bool {
	// Grab the value of the global nonce, hash it. (This is faster than
	// using /dev/urandom)
	newValue := atomic.AddUint64(&nonce, 1)
	hashBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(hashBytes, newValue)
	sum := sha256.Sum256(hashBytes)

	// Pull the top 8 bytes of the result and convert to a uint64, modulus the
	// denominator.
	randVal := binary.LittleEndian.Uint64(sum[:8])
	randVal = randVal % uint64(denom)

	if randVal < uint64(num) {
		return true
	}
	return false
}

// Run a simulation where an attacker tries to commit a double-spend. The
// attacker is given zero latency. The hashrate of the attacker, the latency of
// the honest miners, the block time, and the number of confirmations can all
// be configured. The result is the probability that the attacker's double
// spend will be successful.
//
// Simulation is more accurate when 'honestPropagationTime / blockTime' is an
// integer (result will be rounded to an integer).
func main() {
	// Confiration.
	simulationIterations := int(100e3)
	attackerHashrate := 40             // percent
	blockTime := 6 * 1000              // milliseconds
	honestPropagationTime := 18 * 1000 // milliseconds
	confirmations := 150               // number of blocks
	threads := 2                       // set to # of CPU cores. Works best in factors of 2 and 5 (2, 4, 5, 8, 10, etc.)
	difficulty := 10                   // chance of finding a block is hashrate / difficulty. Increase for slower, more accurate simulation.

	// Variables derived from the tunable variables
	honestHashrate := 100 - attackerHashrate
	attackerAdvantage := honestPropagationTime / blockTime

	// Set the global nonce to a random value, so that different simulations
	// produce different results. There is a small chance of overlap, as the
	// nonce-space is only 64 bits, and each run consumes 16-32 bits of
	// nonce-space.
	r, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		panic(err)
	}
	nonce = uint64(r.Int64())

	// Record what nonce was used to perform the simulation. By setting the
	// nonce manually, a simulation can be replicated exactly.
	nonceStart := nonce

	// Run the simluation.
	var wg sync.WaitGroup
	attackerWins := uint64(0)
	honestWins := uint64(0)
	for t := 0; t < threads; t++ {
		wg.Add(1)
		go func(t int) {
			// In each thread, run 'simulationIterations/threads' iterations.
			for i := 0; i < simulationIterations/threads; i++ {
				// Some output to help track simulation progress. Only tracks
				// progress of first thread, other threads may be ahead or
				// behind.
				if t == 0 && i%(10e3/threads) == 0 {
					fmt.Println(float64(i*100*threads)/float64(simulationIterations), "%")
				}

				// Keep track of the number of blocks each party has found.
				attackerBlocks := 0
				honestBlocks := 0

				// Simulate an attacker and an honest network mining until one
				// of them hits the desired number of confirmations. The
				// attacker can use slow network propatation to their
				// advantage, so they win at a lower height.
				for attackerBlocks < (confirmations-attackerAdvantage) && honestBlocks < confirmations {
					if tryHash(attackerHashrate, 100*difficulty) {
						attackerBlocks++
					}
					if tryHash(honestHashrate, 100*difficulty) {
						honestBlocks++
					}
				}

				// Determine who the winner was. If it's a tie, favor the
				// attacker (assume the attacker network advantage is slightly
				// better than what was simulated via the
				// 'honestPropagationTime' head start).
				if attackerBlocks >= (confirmations - attackerAdvantage) {
					atomic.AddUint64(&attackerWins, 1)
				} else {
					atomic.AddUint64(&honestWins, 1)
				}
			}
			wg.Done()
		}(t)
	}
	wg.Wait()

	// Print the result of the simulation.
	fmt.Print("\nResults of Double-Spend Simulation:")
	fmt.Printf(`
	Iterations:                          %v
	Attacker Hashrate:                   %v%%
	Block Time:                          %v (milliseconds)
	Honest Miner Block Propagation Time: %v (milliseconds)
	Confirmations:                       %v
	Threads:                             %v
	Difficulty:                          %v
	Nonce Starting Point:                %v

	Attacker Wins: %v (%.6f%%)

`, simulationIterations, attackerHashrate, blockTime, honestPropagationTime, confirmations, threads, difficulty, nonceStart, attackerWins, 100*float64(attackerWins)/float64(simulationIterations))
}
