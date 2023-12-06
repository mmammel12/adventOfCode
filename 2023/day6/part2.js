const { getFileData } = require("../../readFile");

const data = getFileData("2023/day6/input.txt");

const time = parseInt(data[0].match(/\d+/g).join(""));
const distance = parseInt(data[1].match(/\d+/g).join(""));

const SPEED_CHANGE_INCREMENT = 1;

let sum = 0;

for (let i = 1; i <= time; i++) {
    const speed = i * SPEED_CHANGE_INCREMENT;

    if (speed * (time - i) > distance) {
        sum++;
    }
}

console.log(sum);
