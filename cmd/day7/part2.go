/*
--- Part Two ---

To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

    32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
    KK677 is now the only two pair, making it the second-weakest hand.
    T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?
*/

package day7

import (
	"aoc2023/internal/util"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)


var CardStrengthJoker = map[rune]int {
    'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, 
    '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 
    'Q': 11, 'K': 12, 'A': 13,
}

func (h *Hand) set_hand_type_with_joker () {
    counts := make (map[rune]int)
    for _, card := range h.cards {
        counts[card]++
    }

    var count_slice []int
    for _, count := range counts {
        count_slice = append (count_slice, count)
    }

    sort.Sort (sort.IntSlice (count_slice))

    // Count the number of jokers
    joker_count := 0
    for _, card := range h.cards {
        if card == 'J' { joker_count++ }
    }

    if joker_count < 5 {
        for i := 0; i < len (count_slice); i++ {
            if count_slice[i] == joker_count {
                // Remove the joker from slice
                count_slice = append (count_slice[:i], count_slice[i+1:]...)
                // Increment last item by joker_count
                count_slice[len (count_slice) - 1] += joker_count
                break
            }
        }
    }

    if reflect.DeepEqual (count_slice, []int{1, 1, 1, 1, 1}) {
        h.hand_type= HIGH_CARD
    } else if reflect.DeepEqual (count_slice, []int{1, 1, 1, 2}) {
        h.hand_type = ONE_PAIR
    } else if reflect.DeepEqual (count_slice, []int{1, 2, 2}) {
        h.hand_type = TWO_PAIR
    } else if reflect.DeepEqual (count_slice, []int{1, 1, 3}) {
        h.hand_type = THREE_OF_A_KIND
    } else if reflect.DeepEqual (count_slice, []int{2, 3}) {
        h.hand_type = FULL_HOUSE
    } else if reflect.DeepEqual (count_slice, []int{1, 4}) {
        h.hand_type = FOUR_OF_A_KIND
    } else if reflect.DeepEqual (count_slice, []int{5}) {
        h.hand_type = FIVE_OF_A_KIND
    }
}


func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    var hands []Hand;

    for buffer.Scan() {
        input := strings.Split (buffer.Text(), " ")
        cards := []rune (input[0])
        bid, _ := strconv.Atoi (input[1])
        hand := NewHand (cards, bid)
        hand.set_hand_type_with_joker()
        hands = append (hands, hand)
    }

    // Sort the hands by their type and strength
    sort.Slice (hands, func (i, j int) bool {
        if hands[i].hand_type == hands[j].hand_type {
            for k := 0; k < len (hands[i].cards); k++ {
                strength_i := CardStrengthJoker[hands[i].cards[k]]
                strength_j := CardStrengthJoker[hands[j].cards[k]]
                if strength_i != strength_j  {
                    return strength_i < strength_j
                } else { continue }
            }
        }
        return hands[i].hand_type < hands[j].hand_type
    })

    for i, hand := range hands {
        sum += hand.bid * (i + 1)
    }

    fmt.Println (sum)
}

