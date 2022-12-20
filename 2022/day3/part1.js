const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day3/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const CHARS = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

const split = (line) => {
    return [line.slice(0, line.length / 2), line.slice(line.length / 2)];
};

console.log(
    data.reduce((acc, line) => {
        let num = 0;
        const [left, right] = split(line);
        for (let i in left) {
            const char = left[i];
            if (right.indexOf(char) > -1) {
                num += CHARS.indexOf(char);
                break;
            }
        }
        return acc + num;
    }, 0)
);
