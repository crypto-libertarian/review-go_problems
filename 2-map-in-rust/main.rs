fn main() {
    let a: Vec<u32> = (0..10).collect();
    let b: Vec<(u32, u32)> = a.iter().map(|&n| (n, n.pow(2))).collect();
    for (n, nn) in b {
        println!("{}: {}", n, nn);
    }
}
