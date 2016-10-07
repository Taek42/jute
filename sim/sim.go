package main

// Note - the random number generator on play.golang.org is bad. Also the
// computational limits are highly restricted, which means you can't run a
// meaningful simulation from this playground anyway.

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

var counter uint64

// Given a probability of success, output a bool indicating whether or not
// success was achieved.
func tryHash(num, denom int) bool {
	// Grab the value of the global counter, hash it. (This is faster than
	// using rand.Reader for whatever reason).
	newValue := atomic.AddUint64(&counter, 1)
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

func main() {
	// Set the counter to a random value.
	r, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		panic(err)
	}
	counter = uint64(r.Int64())

	// Tunable Variables
	simulationIterations := int(1e5) // run locally for higher values
	attackerHashrate := 40           // percent
	honestPropagationTime := 15      // seconds
	blockTime := 6                   // seconds
	maximumGap := 200                // number of blocks

	// Variables derived from the tunable variables
	honestHashrate := 100 - attackerHashrate
	attackerAdvantage := honestPropagationTime / blockTime

	// Variables that affect simulation accuracy
	difficulty := 20 // chance of finding a block is hashrate / difficulty.

	// Parallelism Variables.
	threads := 2
	var wg sync.WaitGroup

	// Launch the threads.
	attackerWins := uint64(0)
	honestWins := uint64(0)
	for t := 0; t < threads; t++ {
		wg.Add(1)
		go func() {
			// In each thread, run 'simulationIterations/threads iterations.
			for i := 0; i < simulationIterations/threads; i++ {
				// Some output to help track simulation progress.
				if i%(1e4/threads) == 0 {
					fmt.Println(float64(i*100*threads)/float64(simulationIterations), "%")
				}

				// Keep track of the number of blocks each party has found.
				attackerBlocks := 0
				honestBlocks := 0

				// Simulate an attacker and an honest network mining until one
				// of them hits the 'maximum gap'. The attacker can use slow
				// network propatation to their advantage, so they actually win
				// at a lower height.
				for attackerBlocks < (maximumGap-attackerAdvantage) && honestBlocks < maximumGap {
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
				if attackerBlocks >= (maximumGap - attackerAdvantage) {
					atomic.AddUint64(&attackerWins, 1)
				} else {
					atomic.AddUint64(&honestWins, 1)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Print the result of the simulation.
	fmt.Printf("Attacker Wins: %v (%v%%)\n", attackerWins, 100*int(attackerWins)/simulationIterations)
	fmt.Printf("Honest Wins: %v (%v%%)\n", honestWins, 100*int(honestWins)/simulationIterations)
}
