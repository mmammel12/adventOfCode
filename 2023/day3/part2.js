const { getFileData } = require("../../readFile");

const data = getFileData("2023/day3/input.txt");

const lineData = data.map((line, i) => {
    const nums = line.match(/\d+/g);
    const numData = nums?.map((num) => {
        const number = parseInt(num);
        const indexes = num.split("").map((_, i) => line.indexOf(num) + i);
        line = line.replace(num, ".".repeat(num.length));
        return { num: number, indexes };
    });
    const asterisks = line.match(/\*/g);
    const asteriskIndexes = asterisks?.map((asterisk) => {
        const index = line.indexOf(asterisk);
        line = line.replace(asterisk, ".");
        return index;
    });

    return { numData, asteriskIndexes };
});

let sum = 0;
for (let i = 0; i < lineData.length; i++) {
    const { numData, asteriskIndexes } = lineData[i];
    if (!asteriskIndexes) {
        continue;
    }
    const sameLineNumData = numData || [];
    let previousLineNumData = [];
    let nextLineNumData = [];
    if (i > 0) {
        const { numData } = lineData[i - 1];
        previousLineNumData = numData || [];
    }
    if (i < lineData.length - 1) {
        const { numData } = lineData[i + 1];
        nextLineNumData = numData || [];
    }

    for (let asteriskIndex of asteriskIndexes) {
        const numbersToCheck = [...previousLineNumData, ...sameLineNumData, ...nextLineNumData];
        const { isGear, contactNumbers } = checkForGear(asteriskIndex, numbersToCheck);
        const currentSum = sum;
        if (isGear) {
            sum += contactNumbers[0] * contactNumbers[1];
            continue;
        }
    }
}

function checkForGear(asteriskIndex, numbers) {
    let isGear = false;
    let contactNumbers = [];
    for (let numberInfo of numbers) {
        const { indexes } = numberInfo;
        for (let index of indexes) {
            if (index - 1 === asteriskIndex || index + 1 === asteriskIndex || index === asteriskIndex) {
                contactNumbers.push(numberInfo.num);
                break;
            }
        }
        if (contactNumbers.length === 2) {
            isGear = true;
            break;
        }
    }

    return { isGear, contactNumbers };
}

console.log(sum);
