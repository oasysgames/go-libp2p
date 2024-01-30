package identify

type Override struct {
	activationThresh int
}

// Overriding package variables.
func OverrideUnsafe(o Override) {
	log.Info("Override pkg variables", "activationThresh", o.activationThresh)
	ActivationThresh = o.activationThresh
}
