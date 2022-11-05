// https://adventofcode.com/2021/day/3
use itertools::Itertools;
use std::{cmp::Ordering, fs};

fn main() {
    const FILEPATH: &str = "src/input.txt";
    let input: Vec<Vec<char>> = fs::read_to_string(FILEPATH)
        .unwrap()
        .split("\n")
        .map(|s| s.chars().collect_vec())
        .collect();
    // let rates = get_rates(input.clone());
    let oxygen_rating = get_rating(input.clone(), 'g');
    let co2_rating = get_rating(input.clone(), 'e');
    println!(
        "{}",
        usize::from_str_radix(&*oxygen_rating, 2).unwrap()
            * usize::from_str_radix(&*co2_rating, 2).unwrap()
    );
}

// Initialise an array of zeroes the length of the binary number
// For each bit, if it's a zero --, if one ++
// Inverse for epsilon rate

fn get_rates(readings: Vec<Vec<char>>) -> (String, String) {
    let mut gamma = String::new();
    let mut epsilon = String::new();
    for i in 0..readings[0].len() {
        let mut common = 0;
        for r in &readings {
            match r[i] {
                '0' => common -= 1,
                '1' => common += 1,
                _ => panic!(),
            }
        }
        match common.cmp(&0) {
            Ordering::Greater | Ordering::Equal => {
                gamma.push('1');
                epsilon.push('0');
            }
            Ordering::Less => {
                gamma.push('0');
                epsilon.push('1');
            }
        }
    }
    (gamma, epsilon)
}

fn get_rating(readings: Vec<Vec<char>>, mode: char) -> String {
    let mut filtered = readings.clone();
    for i in 0..filtered[0].len() {
        // Find the most common bit for the selected column
        let mut common = 0;
        for l in &filtered {
            match l[i] {
                '0' => common -= 1,
                '1' => common += 1,
                _ => panic!(),
            }
        }
        // Initialise the rate value for the column
        let bit: char;
        match common.cmp(&0) {
            Ordering::Greater | Ordering::Equal => match mode {
                'g' => bit = '1',
                'e' => bit = '0',
                _ => panic!(),
            },
            Ordering::Less => match mode {
                'g' => bit = '0',
                'e' => bit = '1',
                _ => panic!(),
            },
        }
        // Filter by rate
        filtered.retain(|l| l[i] == bit);
        if filtered.len() == 1 {
            break;
        }
    }
    let rating: String = filtered[0].iter().collect();
    rating
}
