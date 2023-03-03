package tardigrade

// Built Sat 25 Feb 16:15:57 GMT 2023

const DBFile = "tardigrade.db"
const Release = "0.1.3"
const Updated = "Fri  3 Mar 21:18:31 GMT 2023"

// Tardigrade is the main structure
type Tardigrade struct{}

// GetVersion function returns the current release version
func (tar *Tardigrade) GetVersion() (release string) {
	release = Release
	return release
}

// GetUpdated function returns the last updated time
func (tar *Tardigrade) GetUpdated() (updated string) {
	updated = Updated
	return updated
}
