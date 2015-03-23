package cfmanifest

// NewSSHManifest prepares for a new cf-ssh.yml
func NewSSHManifest(appName string, command string) (manifest *Manifest) {
	manifest = NewManifest()
	cfssh := manifest.AddApplication(appName)
	if command == "" {
		command = "curl http://tmate-bootstrap.cfapps.io | sh"
	}
	cfssh["command"] = command
	cfssh["no-route"] = true
	cfssh["instances"] = 1
	return
}

// NewSSHManifestFromManifestPath prepares for a new cf-ssh.yml based on existing manifest.yml
func NewSSHManifestFromManifestPath(manifestPath string, command string) (manifest *Manifest, err error) {
	manifest, err = NewManifestFromPath(manifestPath)
	if err != nil {
		return
	}
	cfssh := manifest.FirstApplication()
	if command == "" {
		command = "curl http://tmate-bootstrap.cfapps.io | sh"
	}

	name := cfssh["name"].(string)
	cfssh["name"] = name + "-ssh"
	cfssh["command"] = command
	cfssh["no-route"] = true
	cfssh["instances"] = 1

	manifest.RemoveAllButFirstApplication()
	return
}
