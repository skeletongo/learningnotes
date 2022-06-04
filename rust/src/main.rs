use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("猜猜看");
    println!("请输入0~100中的数字");

    let n = rand::thread_rng().gen_range(1..101);
    //println!("rand num {}", n);
    let mut guess = String::new();

    loop {
        println!("请输入");
        guess.clear();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let num: i32 = match guess.trim().parse() {
            Ok(t) => t,
            Err(_) => {
                println!("无效数字");
                continue;
            }
        };

        if num < 0 || num > 100 {
            println!("请输入0~100中的数字");
            continue;
        }

        match num.cmp(&n) {
            Ordering::Less => {
                println!("小了");
            }
            Ordering::Equal => {
                println!("猜对了");
                break;
            }
            Ordering::Greater => {
                println!("大了");
            }
        }
    }
}
