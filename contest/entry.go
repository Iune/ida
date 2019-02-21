package contest

type Entry struct {
	Country Country
	Artist  string
	Song    string
}

func (entry Entry) Equals(other Entry) bool {
	if entry.Artist == other.Artist && entry.Song == other.Song && entry.Country.GetPrimaryName() == other.Country.GetPrimaryName() {
		return true
	}
	return false
}
