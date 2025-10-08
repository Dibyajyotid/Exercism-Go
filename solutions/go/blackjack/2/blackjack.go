package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    cardValues := map[string]int {
        "ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
    }

    if val, ok := cardValues[card]; ok {
        return val;
    }

    return 0;
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

    card1Value, card2Value := ParseCard(card1), ParseCard(card2);
    dealerCardValue := ParseCard(dealerCard);
    cardSumValue := card1Value + card2Value;
	switch {
        case card1Value == 11 && card2Value == 11:
        	return "P";

        case cardSumValue == 21 && dealerCardValue < 10:
        	return "W";

        case cardSumValue == 21 && dealerCardValue >= 10:
        	return "S";
		
        case cardSumValue >= 17 && cardSumValue <= 20:
        	return "S";

        case cardSumValue >= 12 && cardSumValue <= 16 && dealerCardValue < 7:
        	return "S";

        case cardSumValue >= 12 && cardSumValue <= 16 && dealerCardValue >= 7:
        	return "H";

        case cardSumValue <= 11:
        	return "H";
    }
    return "";
}
