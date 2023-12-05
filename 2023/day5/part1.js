const { getFileData } = require("../../readFile");

const data = getFileData("2023/day5/input.txt");

const SEEDS_TO_SOIL_START = 3;
// const SEEDS_TO_SOIL_END = 5;
const SEEDS_TO_SOIL_END = 25;

const SOIL_TO_FERTILIZER_START = SEEDS_TO_SOIL_END + 2;
// const SOIL_TO_FERTILIZER_END = 10;
const SOIL_TO_FERTILIZER_END = 55;

const FERTILIZER_TO_WATER_START = SOIL_TO_FERTILIZER_END + 2;
// const FERTILIZER_TO_WATER_END = 16;
const FERTILIZER_TO_WATER_END = 105;

const WATER_TO_LIGHT_START = FERTILIZER_TO_WATER_END + 2;
// const WATER_TO_LIGHT_END = 20;
const WATER_TO_LIGHT_END = 145;

const LIGHT_TO_TEMP_START = WATER_TO_LIGHT_END + 2;
// const LIGHT_TO_TEMP_END = 25;
const LIGHT_TO_TEMP_END = 194;

const TEMP_TO_HUMIDITY_START = LIGHT_TO_TEMP_END + 2;
// const TEMP_TO_HUMIDITY_END = 29;
const TEMP_TO_HUMIDITY_END = 236;

const HUMIDITY_TO_LOCATION_START = TEMP_TO_HUMIDITY_END + 2;
// const HUMIDITY_TO_LOCATION_END = 33;
const HUMIDITY_TO_LOCATION_END = 261;

const DESTINATION_INDEX = 0;
const SOURCE_INDEX = 1;
const RANGE_INDEX = 2;

// const createMap = (start, end) => {
//     const map = {};
//     for (let i = start; i < end; i++) {
//         const line = data[i].split(" ");
//         const destination = parseInt(line[DESTINATION_INDEX]);
//         const source = parseInt(line[SOURCE_INDEX]);
//         const range = parseInt(line[RANGE_INDEX]);

//         for (let j = 0; j < range; j++) {
//             map[source + j] = destination + j;
//         }
//     }

//     return map;
// };

const createRanges = (start, end) => {
    const ranges = [];

    for (let i = start; i < end; i++) {
        const line = data[i].split(" ");
        const destination = parseInt(line[DESTINATION_INDEX]);
        const source = parseInt(line[SOURCE_INDEX]);
        const range = parseInt(line[RANGE_INDEX]);

        ranges.push({
            destination,
            source,
            range,
        });
    }

    return ranges;
};

const seeds = data[0]
    .split(":")[1]
    .trim()
    .split(" ")
    .map((item) => parseInt(item));

const seedsToSoilRanges = createRanges(SEEDS_TO_SOIL_START, SEEDS_TO_SOIL_END);
const soilToFertilizerRanges = createRanges(SOIL_TO_FERTILIZER_START, SOIL_TO_FERTILIZER_END);
const fertilizerToWaterRanges = createRanges(FERTILIZER_TO_WATER_START, FERTILIZER_TO_WATER_END);
const waterToLightRanges = createRanges(WATER_TO_LIGHT_START, WATER_TO_LIGHT_END);
const lightToTempRanges = createRanges(LIGHT_TO_TEMP_START, LIGHT_TO_TEMP_END);
const tempToHumidityRanges = createRanges(TEMP_TO_HUMIDITY_START, TEMP_TO_HUMIDITY_END);
const humidityToLocationRanges = createRanges(HUMIDITY_TO_LOCATION_START, HUMIDITY_TO_LOCATION_END);

let minLocation = Infinity;

const getDestination = (ranges, source) => {
    const range = ranges.find((range) => range.source <= source && source < range.source + range.range);
    // if range not found, return source
    if (range === undefined) {
        return source;
    }
    return range.destination + (source - range.source);
};

seeds.forEach((seed) => {
    const soil = getDestination(seedsToSoilRanges, seed);
    const fertilizer = getDestination(soilToFertilizerRanges, soil);
    const water = getDestination(fertilizerToWaterRanges, fertilizer);
    const light = getDestination(waterToLightRanges, water);
    const temp = getDestination(lightToTempRanges, light);
    const humidity = getDestination(tempToHumidityRanges, temp);
    const location = getDestination(humidityToLocationRanges, humidity);

    if (location < minLocation) {
        minLocation = location;
    }
});

console.log(minLocation);
