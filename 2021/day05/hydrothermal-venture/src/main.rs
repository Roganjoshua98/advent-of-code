use itertools::Itertools;
use std::fs;

fn main() {
    let path = "src/input.txt";
    let input = fs::read_to_string(path)
        .unwrap()
        .lines()
        .map(|s| s.replace(" -> ", " "))
        .collect_vec();
    let input = input
        .iter()
        .map(|s| s.split(" ").collect_vec())
        .collect_vec();
    for spl in input {
        println!("Hello!");
    }
}   
