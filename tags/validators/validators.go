package validators

type FeatureValidator struct {
	Callback      func(string) bool
	AllowedValues []string
}

func (f *FeatureValidator) Validate(val string) bool {
	if f == nil || f.Callback == nil {
		return true
	}
	return f.Callback(val)
}
