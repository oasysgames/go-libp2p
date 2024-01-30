package identify

type Override struct {
	ActivationThresh int
}

// Overriding package variables.
func OverrideUnsafe(o Override) {
	log.Info("Override pkg variables", "ActivationThresh", o.ActivationThresh)
	ActivationThresh = o.ActivationThresh
}
