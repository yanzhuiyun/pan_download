package controller

type ParamBoolean struct {
	Flag bool `json:"bool"`
}

type ParamString struct {
	Msg string `json:"string"`
}

type ParamStringSlice struct {
	Slice []string `json:"files"`
}

type ParamLogin struct {
	ParamBoolean
}

type ParamSignUp struct {
	ParamString
	ParamBoolean
}

type ParamFiles struct {
	Filenames []string `json:"filenames"`
}

func (signup *ParamSignUp) Set(msg string, flag bool) {
	signup.Msg = msg
	signup.Flag = flag
}
