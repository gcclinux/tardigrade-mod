package tardigrade

// Updated - Sun 10 Sep 18:54:19 BST 2023
const Release = "0.2.5"
const Updated = "Sun 10 Sep 18:54:19 BST 2023"

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
