pub fn add(a: i32, b: i32) -> i32 {
    return a + b;
}

fn scope(a: i32) {
    if a > 100 {
        panic!("gt 100");
    } else {
        panic!("le 100")
    }
}

// 函数测试
#[test]
//#[should_panic] // 忽略panic
//#[ignore] // 忽略这个测试
fn mmadd() {
    assert_eq!(2, add(1, 1));
}

// 模块测试
#[cfg(test)]
mod mm {
    use super::add; // super 当前模块所在的模块

    #[test]
    fn testadd() {
        let r = add(1, 1);
        assert_eq!(2, r);
    }

    #[test]
    #[should_panic(expected = "gt")] // 可以panic,但是错误信息要包含“gt”
    fn testscope() {
        super::scope(101)
    }
}

#[cfg(test)]
mod mm2 {
    #[test]
    fn it_works() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
}
