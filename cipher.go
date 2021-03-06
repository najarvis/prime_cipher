package main

import (
    "fmt"
    "sort"
    "os"
    "strconv"
)

func main() {
    var num int

    if (len(os.Args) == 1) {
        fmt.Print("Enter max number: ")
        fmt.Scan(&num)
    } else {
        var err error
        num, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Println("Usage: ./prime_cipher.exe [max_number]")
            return
        }
    }

    fmt.Println(sortAndPrintSet(GetPrimes(num)))

}

// GetPrimes is an implementation of the Sieve of Eratosthenes to generate primes in Golang.
// It creates a set of all (odd) numbers up to an input num, then goes through each number
// in the set in ascending order and removes all multiples from the set. A set of primes is returned.
func GetPrimes(num int) map[int]struct{} {

    // There isn't a set datastructure, so this int->empty struct map will do. (Basically the same thing).
    prime_map := make(map[int]struct{})

    prime_map[2] = struct{}{}
    for i := 3; i < num; i+=2 {
        prime_map[i] = struct{}{}
    }

    // 2 is the min prime_num, but we didn't add in any of its multiples at the start so we start with 3.
    index := 3
    for index * index <= num {
        // If the index isn't a prime number, all of it's multiples will already be out of the set.
        if _, ok := prime_map[index]; !ok {
            index++
            continue
        }
        // Go through all the multiples of the number we are on and remove them from the set.
        for mul := 2; index * mul <= num; mul++ {
            delete(prime_map, mul * index)
        }
        index++
    }

    return prime_map;
}

func sortAndPrintSet(m map[int]struct{}) (primes []int) {

    // loop through all int->empty struct pairs and if the key is present then add it to our slice of primes.
    // These lines would be something like [k for k in prime_map] in python
    for k, _ := range m {
        primes = append(primes, k)
    }

    // Sort them so they are a bit nicer to look at.
    return_primes := sort.IntSlice(primes)
    sort.Sort(return_primes)
    return return_primes
}
