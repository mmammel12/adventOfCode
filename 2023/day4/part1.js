const { getFileData } = require("../../readFile");

const data = getFileData("2023/day4/input.txt");

const lineData = data.map((line) => {
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
    return { winningNumbers, cardNumbers };
});

let sum = 0;

lineData.forEach((line) => {
    const { winningNumbers, cardNumbers } = line;
    const matchingNumbers = winningNumbers.filter((num) => cardNumbers.includes(num));
    if (matchingNumbers.length > 0) {
        const points = 2 ** (matchingNumbers.length - 1);
        sum += points;
    }
});

console.log(sum);
