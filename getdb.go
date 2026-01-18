package tardigrade

// Updated - Sun Jan 18 09:38:18 PM GMT 2026
const Release = "0.3.0"
const Updated = "Sun Jan 18 09:38:18 PM GMT 2026"

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
