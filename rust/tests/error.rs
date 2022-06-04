#[test]
fn test1() {
    // panic 终止程序运行
    fn p() {
        panic!("big error");
    }
    p();
}

#[test]
fn test2() {
    // Result<T,E>
    use std::fs::File;
    let f = File::open("hello.txt");
    match f {
        Ok(file) => file,
        Err(err) => panic!("err={}", err),
    };
}

#[test]
fn test3() {
    use std::fs::File;
    // 简写1:打印错误信息
    File::open("hello.txt").unwrap();
    // 简写2:打印错误信息，附带自定义错误信息
    File::open("hello.txt").expect("expect error");
}

#[test]
fn test4() {
    // ? 号运算符

    use std::fs;
    use std::fs::File;
    use std::io;
    use std::io::Read; //todo 为什么需要引入这个

    // 分开写法
    fn read_file_to_string() -> Result<String, io::Error> {
        let mut f = File::open("hello.txt")?; // 打开失败直接return Err(e)
        let mut s = String::new();
        f.read_to_string(&mut s)?; // 读取错误直接return Err(e)
        Ok(s)
    }

    // 链式写法
    fn read_file_to_string2() -> Result<String, io::Error> {
        let mut s = String::new();
        File::open("hello.txt")?.read_to_string(&mut s)?;
        Ok(s)
    }

    // 库提供的函数
    fn read_file_to_string3() -> Result<String, io::Error> {
        fs::read_to_string("hello.txt")
    }
}
