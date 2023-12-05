/*
--- Day 5: If You Give A Seed A Fertilizer ---

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48

The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53 corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51

With this map, you can look up the soil number required for each initial seed number:

    Seed number 79 corresponds to soil number 81.
    Seed number 14 corresponds to soil number 14.
    Seed number 55 corresponds to soil number 57.
    Seed number 13 corresponds to soil number 13.

The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds. To do this, you'll need to convert each seed number through other categories until you can find its corresponding location number. In this example, the corresponding types are:

    Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
    Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
    Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
    Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?
*/

package day5

import (
	"aoc2023/internal/util"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)


type Part1 struct { util.Part }
type Part2 struct { util.Part }

var Challenge util.Challenge = util.Challenge {
    Part1: Part1{}, Part2: Part2{},
}


func (p *Part1) string_field_to_ints (field string) (ints []int) {
    var int_field []int
    str_field := strings.Fields (field)
    for _, v := range str_field {
        n,_ := strconv.Atoi (v)
        int_field = append (int_field, n)  
    }
    return int_field 
}


func (p *Part1) get_mapping (buffer *bufio.Scanner, prefix string) ([][]int) {
    var mapping [][]int
    for buffer.Scan() {
        line := buffer.Text()
        if strings.HasPrefix (line, prefix) {
            for buffer.Scan() {
                intLine := buffer.Text()
                if intLine == "" { break }
                mapping = append (mapping, p.string_field_to_ints (intLine))
            }
            break
        }
    }
    return mapping
}


func (p *Part1) get_mapping_id (key int, mapping [][]int) (int) {
    var id int
    var IsMapped bool = false

    for _, values := range mapping {
        if key >= values[1] && key <= values[1] + values[2] {
            id = values[0] + (key - values[1])
            IsMapped = true
            break
        }
    }
    if !IsMapped { id = key }

    return id
}


func (p Part1) Run (input string) {
    var lowest_location int = math.MaxInt
    var seeds []int
    var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation [][]int

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    for buffer.Scan() {
        line := buffer.Text()
        if line == "" { continue }

        // Get seeds
        if strings.HasPrefix (line, "seeds:") {
            line = strings.TrimPrefix (line, "seeds:")
            seeds = p.string_field_to_ints (line)
        }

        seedToSoil = p.get_mapping (buffer, "seed-to-soil map:")
        soilToFertilizer = p.get_mapping (buffer, "soil-to-fertilizer map:")
        fertilizerToWater = p.get_mapping (buffer, "fertilizer-to-water map:")
        waterToLight = p.get_mapping (buffer, "water-to-light map:")
        lightToTemperature = p.get_mapping (buffer, "light-to-temperature map:")
        temperatureToHumidity = p.get_mapping (buffer, "temperature-to-humidity map:")
        humidityToLocation = p.get_mapping (buffer, "humidity-to-location map:")
    }

    for _, seed := range seeds {
        var soilId, fertilizerId, waterId, lightId, temperatureId, humidityId, locationId int

        soilId = p.get_mapping_id (seed, seedToSoil)
        fertilizerId = p.get_mapping_id (soilId, soilToFertilizer)
        waterId = p.get_mapping_id (fertilizerId, fertilizerToWater)
        lightId = p.get_mapping_id (waterId, waterToLight)
        temperatureId = p.get_mapping_id (lightId, lightToTemperature)
        humidityId = p.get_mapping_id (temperatureId, temperatureToHumidity)
        locationId = p.get_mapping_id (humidityId, humidityToLocation)

        if locationId < lowest_location {
            lowest_location = locationId
        }
    }

    fmt.Println ("Lowest location:", lowest_location)
}

