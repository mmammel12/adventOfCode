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

const numColors = {
    red: 12,
    green: 13,
    blue: 14,
};

let sum = 0;

for (game of games) {
    const { gameNum, rounds } = game;
    let isPossible = true;
    for (pulls of rounds) {
        if (!isPossible) break;
        for (pull of pulls) {
            const { color, value } = pull;
            if (value > numColors[color]) {
                isPossible = false;
                break;
            }
        }
    }

    if (isPossible) {
        sum += gameNum;
    }
}

console.log(sum);
