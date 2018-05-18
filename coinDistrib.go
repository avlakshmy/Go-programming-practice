/*You have 50 coins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth.
The coins will be distributed based on the vowels contained in each name where: a: 1 coin, e: 1 coin, i: 2 coins, o: 3 coins, u: 4 coins.
and a user can’t get more than 10 coins. 
Print a map with each user’s name and the amount of coins distributed.*/

package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))	
)

func valVowel (c byte) int {
	if c == 'a' || c == 'A' || c == 'e' || c == 'E' {
		return 1
	} else if c == 'i' || c == 'I' {
		return 2
	} else if c == 'o' || c == 'O' {
		return 3
	} else if c == 'u' || c == 'U' {
		return 4
	}	
	return 0
}

func countVowel (s string) int {
	b := []byte(s)
	c := 0
	for i := 0; i < len(b); i++ {
		c = c + valVowel(b[i]) 
		if c > 10 {
			return 10
		}
	}
	return c
}
	
func main() {
	fmt.Println("Initially:")
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
	for _, v := range users {
		vowels := countVowel(v)
		distribution[v] = vowels
		coins = coins - vowels
	}	
	fmt.Println()
	fmt.Println("After distribution of coins:")
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
