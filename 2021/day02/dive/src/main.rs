// https://adventofcode.com/2021/day/2
use std::fs;
use itertools::Itertools;

struct Position {
    horizontal: i32,
    depth: i32,
}

struct Aim {
    aim: i32,
    pos: Position,
}

fn main() {
    const FILEPATH: &str = "src/input.txt";
    let file = fs::read_to_string(FILEPATH).unwrap();
    let commands: Vec<(&str, i32)> = file
        .lines()
        .map(|l| l.split(" ").collect_vec())
        .map(|vec| (vec[0], vec[1].parse::<i32>().unwrap()))
        .collect_vec();
    let x = aim(&commands);
    let result = x.pos.horizontal * x.pos.depth;
    print!("{}", result);
}

fn depth(commands: &Vec<(&str, i32)>) -> Position {
    let mut direction = Position { horizontal: 0, depth: 0 };
    for c in commands {
        let n = c.1;
        match c.0 {
            "down"      => direction.depth += n,
            "up"        => direction.depth -= n,
            "forward"   => direction.horizontal += n,
            _           => panic!()
        }
    }
    direction
}

fn aim(commands: &Vec<(&str, i32)>) -> Aim {
    let mut direction = Aim { aim: 0, pos: Position { horizontal: 0, depth: 0 }};
    for c in commands {
        let n = c.1;
        match c.0 {
            "down"      => direction.aim += n,
            "up"        => direction.aim -= n,
            "forward"   => {
                direction.pos.horizontal += n;
                direction.pos.depth += n * direction.aim;
            },
            _           => panic!()
        }
    }
    direction
}
