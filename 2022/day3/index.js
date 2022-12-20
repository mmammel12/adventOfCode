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

const groupedData = [];
for (let i = 0; i < data.length; i += 3) {
    groupedData.push([data[i], data[i + 1], data[i + 2]]);
}

console.log(
    groupedData.reduce((acc, group) => {
        let num = 0;
        const [left, middle, right] = group;
        for (let i in left) {
            const char = left[i];
            if (middle.indexOf(char) > -1 && right.indexOf(char) > -1) {
                num += CHARS.indexOf(char);
                break;
            }
        }
        return acc + num;
    }, 0)
);
