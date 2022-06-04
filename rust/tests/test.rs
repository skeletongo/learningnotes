// 集成测试

// 导入 library crate
use rust;

// 路径导入
use common::hehe as haha;

// 目录模块导入
mod common; // 公共库，目录中包含mod.rs，直接目录就是模块名不好吗

// 同名路径优化
//use common::hehe;
//use common::One;
//上面两个导入可以合并成如下所示
use common::{hehe, One};

#[test]
fn testadd() {
    assert_eq!(2, rust::add(1, 1));
}

#[test]
fn setup() {
    assert_eq!(2, rust::add(1, 1));

    // 路径导入的数据
    haha::hi();

    // 目录模块导入，只能使用mod.rs中导出的数据
    common::hehe::hi();
    common::setup();
}

// 模块:
//  1.文件就是一个模块
//  2.目录是一个模块，但是导出的内容必须在目录中的mod.rs文件中

// 模块导入: mod 模块名（文件名或目录名）
// 模块重新导出: pub mod 模块名（文件名或目录名）
