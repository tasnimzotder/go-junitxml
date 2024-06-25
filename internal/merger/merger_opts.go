package merger

type MergeOptions struct {
	SingleTestSuite bool
}

func defaultMergeOptions() *MergeOptions {
	return &MergeOptions{
		SingleTestSuite: false,
	}
}

type MergeOption func(*MergeOptions)

func WithSingleTestSuites(singleTestSuite bool) MergeOption {
	return func(m *MergeOptions) {
		m.SingleTestSuite = singleTestSuite
	}
}

func parseOptions(opts ...MergeOption) *MergeOptions {
	options := defaultMergeOptions()

	for _, opt := range opts {
		opt(options)
	}

	return options
}
