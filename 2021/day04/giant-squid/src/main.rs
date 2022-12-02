// https://adventofcode.com/2021/day/4
use itertools::Itertools;
use std::fs;

fn main() {
    let path = "src/input.txt";
    let input = fs::read_to_string(path).unwrap();
    let input = input.split("\n\n").collect_vec();
    let numbers = String::from(input[0])
        .split(",")
        .map(|s| s.trim().parse::<i32>().unwrap())
        .collect_vec();
    let boards = input[1..]
        .iter()
        .map(|b| {
            b.split("\n")
                .map(|r| {
                    r.split(" ")
                        .filter(|s| *s != "")
                        .map(|s| (s.parse::<i32>().unwrap(), false))
                        .collect_vec()
                })
                .collect_vec()
        })
        .collect_vec();
    // let win = bingo(numbers, boards).unwrap();
    // let score = sum_score(win.0);
    // println!("{}", score * win.1)
    let lose = losing_board(numbers, boards);
    let score = sum_score(lose.0);
    println!("{}", score * lose.1)
}

fn losing_board(
    numbers: Vec<i32>,
    mut boards: Vec<Vec<Vec<(i32, bool)>>>,
) -> (Vec<Vec<(i32, bool)>>, i32) {
    let mut result: (Vec<Vec<(i32, bool)>>, i32) = (vec![], -1);
    for i in 0..5 {
        boards = call_number(numbers[i], boards);
    }
    for n in numbers[5..].to_vec() {
        boards = call_number(n, boards);
        while let Some(board) = check_boards(&boards) {
            let i = boards.iter().position(|x| *x == board).unwrap();
            if boards.len() == 1 {
                result = (boards.pop().unwrap(), n);
                break;
            } else {
                boards.remove(i);
            }
        }
    }
    result
}

fn _bingo(
    numbers: Vec<i32>,
    mut boards: Vec<Vec<Vec<(i32, bool)>>>,
) -> Option<(Vec<Vec<(i32, bool)>>, i32)> {
    let mut result: Option<(Vec<Vec<(i32, bool)>>, i32)> = None;
    for i in 0..5 {
        boards = call_number(numbers[i], boards);
    }
    for n in numbers[5..].to_vec() {
        boards = call_number(n, boards);
        match check_boards(&boards) {
            None => (),
            Some(b) => {
                result = Some((b, n));
                break;
            }
        }
    }
    result
}

fn call_number(n: i32, boards: Vec<Vec<Vec<(i32, bool)>>>) -> Vec<Vec<Vec<(i32, bool)>>> {
    let boards = boards
        .iter()
        .map(|board| {
            board
                .iter()
                .map(|row| {
                    row.iter()
                        .map(|x| match x.0 == n {
                            true => (x.0, true),
                            false => *x,
                        })
                        .collect_vec()
                })
                .collect_vec()
        })
        .collect_vec();
    boards
}

fn check_boards(boards: &Vec<Vec<Vec<(i32, bool)>>>) -> Option<Vec<Vec<(i32, bool)>>> {
    fn vertical_check(board: &Vec<Vec<(i32, bool)>>) -> bool {
        for col in 0..board.len() {
            let mut list: Vec<bool> = vec![];
            for row in 0..board.len() {
                if board[row][col].1 {
                    list.push(board[row][col].1);
                }
            }
            if list.len() == 5 {
                return true;
            }
        }
        false
    }
    fn horizontal_check(board: &Vec<Vec<(i32, bool)>>) -> bool {
        for row in board {
            if row.iter().filter(|x| x.1).collect_vec().len() == 5 {
                return true;
            }
        }
        false
    }
    // fn diagonal_check(board: &Vec<Vec<(i32, bool)>>) -> bool {
    //     let mut downwards: Vec<bool> = vec![];
    //     let mut upwards: Vec<bool> = vec![];
    //     let mut y: usize = board.len();
    //     for x in 0..board.len() {
    //         y -= 1;
    //         if board[x][y].1 {
    //             downwards.push(board[x][y].1);
    //         }
    //         if board[y][x].1 {
    //             upwards.push(board[y][x].1);
    //         }
    //     }
    //     downwards.len() == 5 || upwards.len() == 5
    // }
    for board in boards {
        if vertical_check(board) || horizontal_check(board) {
            return Some(board.clone());
        }
    }
    None
}

fn sum_score(board: Vec<Vec<(i32, bool)>>) -> i32 {
    let mut sum = 0;
    for row in board {
        let n: i32 = row.iter().filter(|x| !x.1).map(|x| x.0).sum();
        sum += n;
    }
    sum
}
