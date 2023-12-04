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
    const symbols = line.match(/[^a-zA-Z0-9\d.]/g);
    const symbolData = symbols?.map((symbol) => {
        const index = line.indexOf(symbol);
        line = line.replace(symbol, ".");
        return { symbol, index };
    });

    return { numData, symbolData };
});

let sum = 0;

for (let i = 0; i < lineData.length; i++) {
    const { numData, symbolData } = lineData[i];
    const sameLineSymbolData = symbolData || [];
    let previousLineSymbolData = [];
    let nextLineSymbolData = [];
    if (i > 0) {
        const { symbolData } = lineData[i - 1];
        previousLineSymbolData = symbolData || [];
    }
    if (i < lineData.length - 1) {
        const { symbolData } = lineData[i + 1];
        nextLineSymbolData = symbolData || [];
    }

    for (let numInfo of numData) {
        const { num, indexes } = numInfo;
        for (let symbolInfo of [...sameLineSymbolData, ...previousLineSymbolData, ...nextLineSymbolData]) {
            const { index } = symbolInfo;
            const contact = checkForSymbolContact(index, indexes);
            if (contact) {
                sum += num;
                break;
            }
        }
    }
}

function checkForSymbolContact(symbolIndex, numIndexes) {
    let contact = false;
    for (let numIndex of numIndexes) {
        if (numIndex - 1 === symbolIndex || numIndex + 1 === symbolIndex || numIndex === symbolIndex) {
            contact = true;
            break;
        }
    }

    return contact;
}

console.log(sum);
