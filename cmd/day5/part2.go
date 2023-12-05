/*
--- Part Two ---

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?
*/

package day5

import (
	"aoc2023/internal/util"
	"bufio"
	"fmt"
	"math"
	"strings"
)


var part1 = Part1{}

var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, 
        lightToTemperature, temperatureToHumidity, humidityToLocation [][]int


func (p *Part2) get_location_id (seed int, buffer *bufio.Scanner) int {
    var soilId, fertilizerId, waterId, lightId, temperatureId, humidityId, locationId int

    soilId = part1.get_mapping_id (seed, seedToSoil)
    fertilizerId = part1.get_mapping_id (soilId, soilToFertilizer)
    waterId = part1.get_mapping_id (fertilizerId, fertilizerToWater)
    lightId = part1.get_mapping_id (waterId, waterToLight)
    temperatureId = part1.get_mapping_id (lightId, lightToTemperature)
    humidityId = part1.get_mapping_id (temperatureId, temperatureToHumidity)
    locationId = part1.get_mapping_id (humidityId, humidityToLocation)

    return locationId
}


func (p Part2) Run (input string) {
    var seedsOriginal []int
    var lowest_location int = math.MaxInt

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
            seedsOriginal = part1.string_field_to_ints (line)
        }

        seedToSoil = part1.get_mapping (buffer, "seed-to-soil map:")
        soilToFertilizer = part1.get_mapping (buffer, "soil-to-fertilizer map:")
        fertilizerToWater = part1.get_mapping (buffer, "fertilizer-to-water map:")
        waterToLight = part1.get_mapping (buffer, "water-to-light map:")
        lightToTemperature = part1.get_mapping (buffer, "light-to-temperature map:")
        temperatureToHumidity = part1.get_mapping (buffer, "temperature-to-humidity map:")
        humidityToLocation = part1.get_mapping (buffer, "humidity-to-location map:")
    }

    for i := 0; i < len (seedsOriginal); i += 2 {
        if i == len (seedsOriginal) - 1 { break }
        for j := seedsOriginal[i]; j < seedsOriginal[i] + seedsOriginal[i+1]; j++ {
            locationId := p.get_location_id (j, buffer)
            if locationId < lowest_location {
                lowest_location = locationId
            }
        }
    }
  
    fmt.Println ("Lowest location:", lowest_location)
}

