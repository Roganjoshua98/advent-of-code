use std::{fs, iter::Sum};

fn main() {
    const FILEPATH: &str = "input.txt";
    let measurements: Vec<i32> = fs::read_to_string(FILEPATH)
                                .unwrap()
                                .split("\n")
                                .map(|x| x.parse::<i32>().unwrap())
                                .collect();
    println!("{}", sonar_sweep(&measurements));
    println!("{}", group_sweep(&measurements));
}


fn sonar_sweep(measurements: &Vec<i32>) -> i32 {
    let mut count: i32 = 0;
    let mut prev: &i32 = &measurements[0];
    for m in measurements {
        if m > prev {
            count += 1;
        }
        prev = m;
    }
    count
}

fn group_sweep(measurements: &Vec<i32>) -> i32 {
    let mut count: i32 = 0;
    let mut prev_sum: i32 = measurements[2] + measurements[1] + measurements[0];
    for i in 3..measurements.len() {
        let sum = measurements[i] + measurements[i-1] + measurements[i-2];
        if sum > prev_sum {
            count += 1;
        }
        prev_sum = sum;
    }
    count
}
