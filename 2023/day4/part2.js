const { getFileData } = require("../../readFile");

const data = getFileData("2023/day4/input.txt");

const lineData = data.map((line, i) => {
    const splitLine = line.split("|");
    const winningNumbers = splitLine[0]
        .replace(/(card\s+\d+: )/i, "")
        .split(" ")
        .map((num) => parseInt(num.trim()))
        .filter((num) => !isNaN(num));
    const cardNumbers = splitLine[1]
        .split(" ")
        .map((num) => parseInt(num.trim()))
        .filter((num) => !isNaN(num));
    return { winningNumbers, cardNumbers, index: i };
});

let sum = lineData.length;

const addedCards = [];

for (let i = 0; i < lineData.length; i++) {
    const { winningNumbers, cardNumbers } = lineData[i];
    const matchingNumbers = winningNumbers.filter((num) => cardNumbers.includes(num));
    if (matchingNumbers.length > 0) {
        for (let j = 0; j < matchingNumbers.length; j++) {
            addedCards.push(lineData[j + i + 1].index);
        }
    }
    sum += matchingNumbers.length;
}

let i = 0;
while (i < addedCards.length) {
    const { winningNumbers, cardNumbers, index } = lineData[addedCards[i]];
    const matchingNumbers = winningNumbers.filter((num) => cardNumbers.includes(num));
    if (matchingNumbers.length > 0) {
        for (let j = 0; j < matchingNumbers.length; j++) {
            addedCards.push(index + j + 1);
        }
    }
    sum += matchingNumbers.length;
    i++;
}

console.log(sum);
