const { getFileData } = require('../../readFile');

const data = getFileData('2023/day11/input.txt');

const GALAXY_SYMBOL = '#';

const UNIVERSE_EXPANSION_FACTOR = 1000000;

const galaxyPoints = data
    .map((row, y) => {
        return row.split('').reduce((acc, symbol, x) => {
            if (symbol === GALAXY_SYMBOL) {
                acc.push({ x, y });
            }
            return acc;
        }, []);
    })
    .flat();

const emptyRows = data.reduce((acc, row, i) => {
    if (!row.includes(GALAXY_SYMBOL)) {
        acc.push(i);
    }
    return acc;
}, []);

const emptyColumns = data[0].split('').reduce((acc, _, i) => {
    if (!data.some((row) => row[i] === GALAXY_SYMBOL)) {
        acc.push(i);
    }
    return acc;
}, []);

let sum = 0;

for (let galaxy1Point of galaxyPoints) {
    for (let galaxy2Point of galaxyPoints) {
        const x1 = Math.min(galaxy1Point.x, galaxy2Point.x);
        const x2 = Math.max(galaxy1Point.x, galaxy2Point.x);
        const y1 = Math.min(galaxy1Point.y, galaxy2Point.y);
        const y2 = Math.max(galaxy1Point.y, galaxy2Point.y);
        let xDistance = 0;
        let yDistance = 0;

        for (let y = y1; y < y2; y++) {
            yDistance++;
            if (emptyRows.includes(y)) {
                yDistance += UNIVERSE_EXPANSION_FACTOR - 1;
            }
        }

        for (let x = x1; x < x2; x++) {
            xDistance++;
            if (emptyColumns.includes(x)) {
                xDistance += UNIVERSE_EXPANSION_FACTOR - 1;
            }
        }
        sum += xDistance + yDistance;
    }
}

sum = sum / 2;

console.log(sum);
