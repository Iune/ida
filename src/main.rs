use clap::{Parser, Subcommand};
use std::path::PathBuf;

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    cmd: Commands,
}

#[derive(Subcommand, Debug, Clone)]
enum Commands {
    Spreadsheet,
    Vote,
    Format { file: PathBuf },
}

fn main() {
    let args = Args::parse();

    match args.cmd {
        Commands::Format { file } => println!("Formatting {}", file.display()),
        _ => todo!(),
    }
}
