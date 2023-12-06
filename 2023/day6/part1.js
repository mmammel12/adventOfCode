const { getFileData } = require("../../readFile");

const data = getFileData("2023/day6/input.txt");

const times = data[0].match(/\d+/g).map(Number);
const distances = data[1].match(/\d+/g).map(Number);

const SPEED_CHANGE_INCREMENT = 1;

let sum = 1;

times.forEach((time, i) => {
    const distance = distances[i];
    let numberOfWaysToWin = 0;
    for (let i = 0; i <= time; i++) {
        const speed = i * SPEED_CHANGE_INCREMENT;
        if (speed * (time - i) > distance) {
            numberOfWaysToWin++;
        }
    }
    sum *= numberOfWaysToWin;
});

console.log(sum);
