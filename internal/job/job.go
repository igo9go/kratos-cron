package job111

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewExampleJob)

type JobFunc func()

var DefaultJobs map[string]JobFunc
