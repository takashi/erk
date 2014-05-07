package main

type Adapter interface {
	Update() error
}

func HasRemoteConfiguration() bool {
	if config.Remote {
		if config.RemoteConfig.ApiToken != "" && config.RemoteConfig.Repo != "" {
			return true
		}
		return false
	}
	return false
}
