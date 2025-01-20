use clap::{Parser, Subcommand};
use std::path::PathBuf;

mod formatter;

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
    Format {
        /// Path to file with results to format
        file: PathBuf,
        /// Format of file to parse. Should be `csv`, `tsv`, or `xlsx`.
        #[arg(short, long)]
        format: String,
        /// Number of digits for formatting point totals.
        #[arg(short, long, default_value_t = 3)]
        digits: usize,
    },
}

fn main() {
    let args = Args::parse();

    match args.cmd {
        Commands::Format {
            file,
            format,
            digits,
        } => formatter::format_results(file, format, digits),
        _ => todo!(),
    }
}
