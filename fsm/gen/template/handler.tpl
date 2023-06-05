package handlers

import (
	"test/fsm"
	"{{.PackageDir}}/{{.Package}}/types"
)

func GetHandler_{{.Transition.Event}}_{{.Transition.From}}() fsm.Handler[types.Payload] {
	return fsm.Handler[types.Payload]{
		BeforeHandler: _BeforeHandler_{{.Transition.Event}}_{{.Transition.From}},
		Handler:       _Handler_{{.Transition.Event}}_{{.Transition.From}},
		AfterHandler:  _AfterHandler_{{.Transition.Event}}_{{.Transition.From}},
	}
}

func _BeforeHandler_{{.Transition.Event}}_{{.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    {{- range $fsmName := .FSMNameList}}
    case types.FSMName{{$fsmName}}:
		return _BeforeHandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f)
    {{- end}}
	default:
		return nil
	}
}

func _Handler_{{.Transition.Event}}_{{.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
	{{- range $fsmName := .FSMNameList}}
    case types.FSMName{{$fsmName}}:
		return _HandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f)
    {{- end}}
	default:
		return nil
	}

}

func _AfterHandler_{{.Transition.Event}}_{{.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
	{{- range $fsmName := .FSMNameList}}
    case types.FSMName{{$fsmName}}:
		return _AfterHandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f)
    {{- end}}
	default:
		return nil
	}
}

{{range $fsmName := .FSMNameList}}
func _BeforeHandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _HandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerFor{{$fsmName}}_{{$.Transition.Event}}_{{$.Transition.From}}(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
{{- end}}
