struct Country {
    forum: String,
    iso: String,
    names: Vec<String>,
}

impl Country {
    fn primary_name(&self) -> &String {
        &self.names[0]
    }

    fn contains_name(&self, line: &String) -> bool {
        let contains: Vec<&String> = self.names
            .iter()
            .filter(|name| line.to_lowercase().contains(&name.to_lowercase()))
            .collect();
        contains.len() > 0
    }
}

struct Entry {
    country: Country,
    artist: String,
    song: String,
}

impl Entry {
    fn flag(&self) -> String {
        format!("World/{}.png", self.country.iso)
    }
}

struct Contest {
    countries: Vec<Country>,
    entries: Vec<Entry>,
    voters: Vec<Country>,
}

impl Contest {
    fn find_voter_by_country_name(&self, voter_name: &String) -> Option<&Country> {
        let search: Vec<&Country> = self.voters.into_iter().filter(|voter| voter.contains_name(voter_name)).collect();
        if search.len() > 0 {
            return Some(&search[0]);
        }
        None
    }
}

fn main() {
    let estonia = Country {
        forum: "estonia".to_string(),
        iso: "ee".to_string(),
        names: vec!["Estonia".to_string()],
    };

    let ines = Entry {
        country: Country {
            forum: "estonia".to_string(),
            iso: "ee".to_string(),
            names: vec!["Estonia".to_string()],
        },
        artist: "Ines".to_string(),
        song: "Once in a Lifetime".to_string(),
    };

    let contest = Contest {
        countries: &vec![estonia],
        entries: &vec![ines],
        voters: &vec![estonia],
    };

    println!("{}", ines.flag());
    println!("{}", estonia.contains_name(&"12 malta".to_string()));
    match contest.find_voter_by_country_name(&"malta".to_string()) {
        Some(voter) => println!("{}", voter.primary_name()),
        None => println!("No voter was found.")
    };
}
