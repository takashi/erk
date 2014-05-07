package main

type Adapter interface {
	Update( /*[]*Issue*/) error
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
