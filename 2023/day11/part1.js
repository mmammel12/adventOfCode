const { getFileData } = require('../../readFile');

const data = getFileData('2023/day11/input.txt');

const GALAXY_SYMBOL = '#';

const expandUniverseVertically = (universe) => {
    const expandedUniverse = [];
    universe.forEach((row) => {
        if (!row.includes(GALAXY_SYMBOL)) {
            expandedUniverse.push(row);
        }
        expandedUniverse.push(row);
    });

    return expandedUniverse;
};

const arrayColumn = (arr, n) => arr.map((x) => x[n]);

const expandUniverseHorizontally = (universe) => {
    const expandedUniverse = Array(universe.length).fill('');
    for (let i = 0; i < universe[0].length; i++) {
        const column = arrayColumn(universe, i);
        if (!column.includes(GALAXY_SYMBOL)) {
            expandedUniverse.forEach((_, j) => {
                expandedUniverse[j] += universe[j][i];
            });
        }
        expandedUniverse.forEach((_, j) => {
            expandedUniverse[j] += universe[j][i];
        });
    }

    return expandedUniverse;
};

const expandUniverse = (universe) => {
    const expandedUniverse = expandUniverseVertically(universe);
    return expandUniverseHorizontally(expandedUniverse);
};

const expandedUniverse = expandUniverse(data);

const distanceBetweenPoints = (x1, y1, x2, y2) => {
    let distance = 0;
    while (x1 !== x2 || y1 !== y2) {
        if (x1 < x2) {
            x1++;
            distance++;
        } else if (x1 > x2) {
            x1--;
            distance++;
        }

        if (y1 < y2) {
            y1++;
            distance++;
        } else if (y1 > y2) {
            y1--;
            distance++;
        }
    }

    return distance;
};

const galaxyPoints = expandedUniverse
    .map((row, y) => {
        return row.split('').reduce((acc, symbol, x) => {
            if (symbol === GALAXY_SYMBOL) {
                acc.push({ x, y });
            }
            return acc;
        }, []);
    })
    .flat();

let sum = 0;

for (let i = 0; i < galaxyPoints.length; i++) {
    const { x: x1, y: y1 } = galaxyPoints[i];
    for (let j = i + 1; j < galaxyPoints.length; j++) {
        const { x: x2, y: y2 } = galaxyPoints[j];
        const distance = distanceBetweenPoints(x1, y1, x2, y2);
        sum += distance;
    }
}

console.log(sum);
