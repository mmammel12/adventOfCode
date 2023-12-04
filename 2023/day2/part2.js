const { readFile } = require("../../readFile");

let data = [];
readFile(
    "2023/day2/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const games = data.map((line) => {
    const gameAndPulls = line.split(":");
    const gameNum = parseInt(gameAndPulls[0].split(" ")[1]);
    const pulls = gameAndPulls[1].split(";");
    const rounds = pulls.map((round) => {
        const roundData = round.split(",");
        return roundData.map((pull) => {
            const trimmed = pull.trim().split(" ");
            const value = parseInt(trimmed[0]);
            const color = trimmed[1];
            return { color, value };
        });
    });
    return { gameNum, rounds };
});

let sum = 0;

for (game of games) {
    const { rounds } = game;
    let max = {
        red: 0,
        green: 0,
        blue: 0,
    };

    for (pulls of rounds) {
        for (pull of pulls) {
            const { color, value } = pull;
            if (value > max[color]) {
                max[color] = value;
            }
        }
    }

    const minProduct = max.red * max.green * max.blue;

    sum += minProduct;
}

console.log(sum);
