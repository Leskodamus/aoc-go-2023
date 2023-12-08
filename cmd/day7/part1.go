/*
--- Day 7: Camel Cards ---

Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship. (At least it's a cool airship!) It drops you off at the edge of a vast desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles, and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like very large rocks covering half of the horizon. The Elf explains that the rocks are all along the part of Desert Island that is directly above Island Island, making it hard to even get there. Normally, they use big machines to move the rocks and filter the sand, but the machines have broken down because Desert Island recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the game of Camel Cards. Camel Cards is sort of similar to poker except it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

    Five of a kind, where all five cards have the same label: AAAAA
    Four of a kind, where four cards have the same label and one card has a different label: AA8AA
    Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
    Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
    Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
    One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
    High card, where all cards' labels are distinct: 23456

Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

    32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
    KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
    T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?
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


type Part1 struct { util.Part }
type Part2 struct { util.Part }

var Challenge util.Challenge = util.Challenge {
    Part1: Part1{}, Part2: Part2{},
}


var CardStrength = map[rune]int {
    '2': 1, '3': 2, '4': 3, '5': 4, 
    '6': 5, '7': 6, '8': 7, '9': 8, 
    'T': 9, 'J': 10, 'Q': 11, 'K': 12,
    'A': 13,
}


type HandType int

const (
    HIGH_CARD HandType = iota + 1
    ONE_PAIR
    TWO_PAIR
    THREE_OF_A_KIND
    FULL_HOUSE
    FOUR_OF_A_KIND
    FIVE_OF_A_KIND
)


type Hand struct {
    cards []rune
    bid int
    hand_type HandType
}


func NewHand (cards []rune, bid int) Hand {
    var hand Hand
    hand.cards = cards
    hand.bid = bid
    hand.hand_type= hand.get_hand_type()
    return hand
}


func (h *Hand) get_hand_type () HandType {
    var hand_type HandType = 0
    
    counts := make (map[rune]int)
    for _, card := range h.cards {
        counts[card]++
    }

    var count_slice []int
    for _, count := range counts {
        count_slice = append (count_slice, count)
    }

    sort.Sort (sort.IntSlice (count_slice))

    if reflect.DeepEqual (count_slice, []int{1, 1, 1, 1, 1}) {
        hand_type = HIGH_CARD
    } else if reflect.DeepEqual (count_slice, []int{1, 1, 1, 2}) {
        hand_type = ONE_PAIR
    } else if reflect.DeepEqual (count_slice, []int{1, 2, 2}) {
        hand_type = TWO_PAIR
    } else if reflect.DeepEqual (count_slice, []int{1, 1, 3}) {
        hand_type = THREE_OF_A_KIND
    } else if reflect.DeepEqual (count_slice, []int{2, 3}) {
        hand_type = FULL_HOUSE
    } else if reflect.DeepEqual (count_slice, []int{1, 4}) {
        hand_type = FOUR_OF_A_KIND
    } else if reflect.DeepEqual (count_slice, []int{5}) {
        hand_type = FIVE_OF_A_KIND
    }

    return hand_type
}


func (p Part1) Run (input string) {
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
        hands = append (hands, NewHand (cards, bid))
    }

    // Sort the hands by their type and strength
    sort.Slice (hands, func (i, j int) bool {
        if hands[i].hand_type == hands[j].hand_type {
            for k := 0; k < len (hands[i].cards); k++ {
                strength_i := CardStrength[hands[i].cards[k]]
                strength_j := CardStrength[hands[j].cards[k]]
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

