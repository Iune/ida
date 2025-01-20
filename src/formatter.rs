use serde::Deserialize;
use std::path::PathBuf;

#[derive(Debug)]
enum FileFormat {
    Csv,
    Tsv,
    Xlsx,
}

#[derive(Debug, thiserror::Error)]
#[error("`{invalid_format}` is not a valid format. Use one of: `csv`, `tsv` or `xlsx`")]
pub struct InvalidFileFormatError {
    invalid_format: String,
}

impl TryFrom<String> for FileFormat {
    type Error = InvalidFileFormatError;

    fn try_from(value: String) -> Result<Self, Self::Error> {
        let value = value.to_lowercase();
        match value.as_str() {
            "csv" => Ok(FileFormat::Csv),
            "tsv" => Ok(FileFormat::Tsv),
            "xlsx" => Ok(FileFormat::Xlsx),
            _ => Err(InvalidFileFormatError {
                invalid_format: value,
            }),
        }
    }
}

#[derive(Debug, Deserialize, Clone, PartialEq)]
struct Entry {
    country: String,
    flag: String,
    artist: String,
    song: String,
    points: u32,
}

impl Entry {
    pub fn flag_icon(&self) -> String {
        // USA and Belarus need to be treated as special cases as their
        // flag icons don't map directly from their country names
        if self.country.to_lowercase() == "united states" {
            "usa".into()
        } else if self.country.to_lowercase() == "belarus" {
            "altbelarus".into()
        } else {
            self.country.to_lowercase().replace(" ", "")
        }
    }

    pub fn points(&self, num_digits: usize) -> String {
        format!("{:0zfill$}", self.points, zfill = num_digits)
    }
}

pub fn format_results(file: PathBuf, file_format: String, num_digits: usize) {
    let file_format = match FileFormat::try_from(file_format) {
        Ok(format) => format,
        Err(err) => panic!("{err}"),
    };

    let entries: Vec<Entry> = match file_format {
        FileFormat::Csv => parse_delimited_file(file, b','),
        FileFormat::Tsv => parse_delimited_file(file, b'\t'),
        FileFormat::Xlsx => panic!("xlsx file format is not yet supported"),
    };

    for entry in entries {
        println!("{}", format_entry(&entry, num_digits))
    }
}

fn parse_delimited_file(file: PathBuf, delimiter: u8) -> Vec<Entry> {
    let mut rdr = csv::ReaderBuilder::new()
        .has_headers(false)
        .delimiter(delimiter)
        .from_path(file)
        .unwrap();

    let mut entries: Vec<Entry> = Vec::new();
    for result in rdr.deserialize::<Entry>() {
        match result {
            Ok(entry) => entries.push(entry),
            Err(error) => panic!("{error}"),
        }
    }
    entries
}

fn format_entry(entry: &Entry, num_digits: usize) -> String {
    format!(
        "{} | :{}: [B]{}:[/B] {} - {}",
        entry.points(num_digits),
        entry.flag_icon(),
        entry.country,
        entry.artist,
        entry.song
    )
}
