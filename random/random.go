package random

import (
	//"fmt"
	"errors"
	"time"
	"crypto/rand"
	"math/big"
	mathrand "math/rand"
)

// Returns a crypto grade string of len(n).
// Because of the crypto grade it's quite slow
// so don't use if crypto grade is not an requirement
// and speed is important.
func RandString(n int) (string, error) {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#%&?+-_*<>;:"
    symbols := big.NewInt(int64(len(alphanum)))
    states := big.NewInt(0)
    states.Exp(symbols, big.NewInt(int64(n)), nil)
    r, err := rand.Int(rand.Reader, states)
    if err != nil {
        return "", err
    }
    var bytes = make([]byte, n)
    r2 := big.NewInt(0)
    symbol := big.NewInt(0)
    for i := range bytes {
        r2.DivMod(r, symbols, symbol)
        r, r2 = r2, r
        bytes[i] = alphanum[symbol.Int64()]
    }
    return string(bytes), err
}

// Returns a crypto grade random value [0, max]. If max == 10 -> returns value between 0 and 9.
// Because of the crypto grade it's quite slow
// so don't use if crypto grade is not an requirement
// and speed is important.
func RandInt(max int64) (int64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(max))
    if err != nil {
		return 0, err
    }

    return r.Int64(), err
}

// Returns a crypto grade random value [min, max]. If min == 5 && max == 10 -> returns value between 5 and 9.
// Because of the crypto grade it's quite slow
// so don't use if crypto grade is not an requirement
// and speed is important.
func RandIntMinMax(min int64, max int64) (r64 int64, err error) {
	for i := -1; i < 0; i-- {
    	r, err := rand.Int(rand.Reader, big.NewInt(max))
    	if err != nil {
			return r64, err
    	}
    	r64 = r.Int64()
    	//fmt.Println(r64)
    	if r64 >= min {
    		break
    	}
	}
	return
}

// Returns non crypto grade random int [0, max]. If max == 10 -> returns value between 0 and 9.
func PseudoRandInt(max int) (int, error) {
	if max < 0 {
		return 0, errors.New("random: max cannot be smaller than 0")
	}
	mathrand.Seed(time.Now().UnixNano())
	return mathrand.Intn(max), nil
}