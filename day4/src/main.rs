use std::{fs, collections::HashSet};

#[derive(Debug)]
struct Game {
    win_nums: HashSet<i32>,
    my_nums: Vec<i32>,
}

fn main() {
    let mut games = read_games();

    /*
    let mut answer = 0;
    for game in games {
        answer += score1(&game.win_nums, &game.my_nums);
    }
    */

    let mut cards: Vec<i32> = vec![1; games.len()];
    for (pos, game) in games.iter().enumerate() {
        let m = pos + 1 + matched(&game.win_nums, &game.my_nums);
        let mut i = pos + 1;
        while i < m && i < games.len() {
            println!("{}, {}", pos, i);
            cards[i] += cards[pos];
            i += 1;
        }
    }

    println!("{}", cards.iter().sum::<i32>());
}

fn matched(win_nums :&HashSet<i32>, my_nums :&Vec<i32>) -> usize {
    let mut matched = 0;
    for n in my_nums {
        if win_nums.contains(&n) {
            matched += 1;
        }
    }
    matched
}

fn score1(win_nums :&HashSet<i32>, my_nums :&Vec<i32>) -> i32 {
    let mut score = 0;

    for n in my_nums {
        if win_nums.contains(&n) {
            if score == 0 {
                score = 1;
            } else {
                score *= 2;
            }
        }
    }

    score
}

fn read_games() -> Vec<Game> {
    let contents = fs::read_to_string("./input").expect("failed to read input file");
    let lines = contents.lines();

    let mut games = Vec::new();
    for line in lines {
        let mut parts = line.split(":");
        
        // Ignore "Card N:"
        parts.next();
        let parts = parts.next().unwrap_or_default().trim();
        
        let mut parts = parts.split("|");

        // win numbers to hash set
        let win_nums_str = parts.next().unwrap_or_default().trim();
        let win_nums: HashSet<i32> = win_nums_str.split_whitespace().filter_map(|s| s.parse().ok()).collect();

        // my numbers to vector
        let my_nums_str = parts.next().unwrap_or_default().trim();
        let my_nums: Vec<i32> = my_nums_str.split_whitespace().filter_map(|s| s.parse().ok()).collect();
        
        games.push(Game{
            win_nums: win_nums,
            my_nums: my_nums,
        });
    }
    
    games
}
