const { getFileData } = require('../../readFile');

const data = getFileData('2023/day10/input.txt');

const STARTING_POINT = 'S';

const DIRECTIONS = {
    '|': { N: true, E: false, S: true, W: false },
    '-': { N: false, E: true, S: false, W: true },
    L: { N: true, E: true, S: false, W: false },
    J: { N: true, E: false, S: false, W: true },
    7: { N: false, E: false, S: true, W: true },
    F: { N: false, E: true, S: true, W: false },
    '.': { N: false, E: false, S: false, W: false },
};

const getDirection = (x, y, map) => {
    const directions = DIRECTIONS[map[y][x]];
    return Object.keys(directions).filter((direction) => directions[direction]);
};

const getStartingPoint = (map) => {
    for (let y = 0; y < map.length; y++) {
        const x = map[y].indexOf(STARTING_POINT);
        if (x !== -1) {
            return { x, y };
        }
    }
};

const startingPoint = getStartingPoint(data);

// determine which 2 directions connect to the starting point
const northPointSymbol = startingPoint.y > 0 ? data[startingPoint.y - 1][startingPoint.x] : null;
const eastPointSymbol = startingPoint.x < data[0].length - 1 ? data[startingPoint.y][startingPoint.x + 1] : null;
const southPointSymbol = startingPoint.y < data.length - 1 ? data[startingPoint.y + 1][startingPoint.x] : null;
const westPointSymbol = startingPoint.x > 0 ? data[startingPoint.y][startingPoint.x - 1] : null;

const getPreviousDirection = (selectedDirection) => {
    switch (selectedDirection) {
        case 'N':
            return 'S';
        case 'S':
            return 'N';
        case 'E':
            return 'W';
        case 'W':
            return 'E';
    }
};

const findLoop = (startingPoint, map) => {
    let directions;
    let previousDirection;
    let currentPoint;
    if (northPointSymbol !== '.' && northPointSymbol !== null) {
        currentPoint = { x: startingPoint.x, y: startingPoint.y - 1 };
        directions = getDirection(startingPoint.x, startingPoint.y - 1, map);
        previousDirection = 'S';
    } else if (eastPointSymbol !== '.' && eastPointSymbol !== null) {
        currentPoint = { x: startingPoint.x + 1, y: startingPoint.y };
        directions = getDirection(startingPoint.x + 1, startingPoint.y, map);
        previousDirection = 'W';
    } else if (southPointSymbol !== '.' && southPointSymbol !== null) {
        currentPoint = { x: startingPoint.x, y: startingPoint.y + 1 };
        directions = getDirection(startingPoint.x, startingPoint.y + 1, map);
        previousDirection = 'N';
    } else if (westPointSymbol !== '.' && westPointSymbol !== null) {
        currentPoint = { x: startingPoint.x - 1, y: startingPoint.y };
        directions = getDirection(startingPoint.x - 1, startingPoint.y, map);
        previousDirection = 'E';
    }
    let direction = directions.find((direction) => direction !== previousDirection);

    const visited = {
        [`${startingPoint.x},${startingPoint.y}`]: { start: true, nextPoint: { x: currentPoint.x, y: currentPoint.y } },
    };

    while (!visited[`${currentPoint.x},${currentPoint.y}`]) {
        const previousCurrentPoint = { ...currentPoint };
        switch (direction) {
            case 'N':
                currentPoint.y--;
                break;
            case 'S':
                currentPoint.y++;
                break;
            case 'E':
                currentPoint.x++;
                break;
            case 'W':
                currentPoint.x--;
                break;
        }
        previousDirection = getPreviousDirection(direction);
        visited[`${previousCurrentPoint.x},${previousCurrentPoint.y}`] = { nextPoint: { ...currentPoint } };
        if (map[currentPoint.y][currentPoint.x] !== STARTING_POINT) {
            directions = getDirection(currentPoint.x, currentPoint.y, map);
            direction = directions.find((direction) => direction !== previousDirection);
        } else {
            visited[`${currentPoint.x},${currentPoint.y}`].end = true;
        }
    }

    return visited;
};

const loop = findLoop(startingPoint, data);

const farthestPointSteps = Object.keys(loop).length / 2;

console.log(farthestPointSteps);
