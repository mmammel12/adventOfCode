const { getFileData } = require("../../readFile");

const data = getFileData("2023/day7/input.txt");

const cardStrengths = {
    2: 1,
    3: 2,
    4: 3,
    5: 4,
    6: 5,
    7: 6,
    8: 7,
    9: 8,
    T: 9,
    J: 10,
    Q: 11,
    K: 12,
    A: 13,
};

const HIGH_CARD = "HIGH_CARD";
const ONE_PAIR = "ONE_PAIR";
const TWO_PAIR = "TWO_PAIR";
const THREE_OF_A_KIND = "THREE_OF_A_KIND";
const FULL_HOUSE = "FULL_HOUSE";
const FOUR_OF_A_KIND = "FOUR_OF_A_KIND";
const FIVE_OF_A_KIND = "FIVE_OF_A_KIND";

const handStrengths = {
    HIGH_CARD: 1,
    ONE_PAIR: 2,
    TWO_PAIR: 3,
    THREE_OF_A_KIND: 4,
    FULL_HOUSE: 5,
    FOUR_OF_A_KIND: 6,
    FIVE_OF_A_KIND: 7,
};

const hands = data.map((hand) => {
    const handParts = hand.split(" ");
    const cards = handParts[0].split("");
    const bid = parseInt(handParts[1]);

    return {
        cards,
        bid,
    };
});

const getHandStrength = (hand) => {
    const cardCount = {};
    hand.cards.forEach((card) => {
        if (cardCount[card]) {
            cardCount[card]++;
        } else {
            cardCount[card] = 1;
        }
    });

    const cardCountValues = Object.values(cardCount);

    const cardCountValuesSorted = cardCountValues.sort((a, b) => b - a);

    if (cardCountValuesSorted[0] === 5) {
        return FIVE_OF_A_KIND;
    }

    if (cardCountValuesSorted[0] === 4) {
        return FOUR_OF_A_KIND;
    }

    if (cardCountValuesSorted[0] === 3) {
        if (cardCountValuesSorted[1] === 2) {
            return FULL_HOUSE;
        }

        return THREE_OF_A_KIND;
    }

    if (cardCountValuesSorted[0] === 2) {
        if (cardCountValuesSorted[1] === 2) {
            return TWO_PAIR;
        }

        return ONE_PAIR;
    }

    return HIGH_CARD;
};

const handTypes = hands.map((hand) => {
    return {
        hand,
        handStrength: getHandStrength(hand),
    };
});

const compareHands = (hand1, hand2) => {
    const hand1Strength = handStrengths[hand1.handStrength];
    const hand2Strength = handStrengths[hand2.handStrength];

    if (hand1Strength > hand2Strength) {
        return 1;
    }

    if (hand1Strength < hand2Strength) {
        return -1;
    }

    const hand1Cards = hand1.hand.cards;
    const hand2Cards = hand2.hand.cards;

    const hand1CardValues = hand1Cards.map((card) => cardStrengths[card]);
    const hand2CardValues = hand2Cards.map((card) => cardStrengths[card]);

    for (let i = 0; i < hand1CardValues.length; i++) {
        if (hand1CardValues[i] > hand2CardValues[i]) {
            return 1;
        }

        if (hand1CardValues[i] < hand2CardValues[i]) {
            return -1;
        }
    }

    return 0;
};

const sortedHands = handTypes.sort(compareHands);

let sum = 0;

sortedHands.forEach((hand, index) => {
    sum += hand.hand.bid * (index + 1);
});

console.log(sum);
