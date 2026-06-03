package tmux

type HookType string
type HookCmd string

type Hook struct {
	Type HookType
	Cmd  HookCmd
}

const (
	HookTypeSessionClosed  HookType = "session-closed"
	HookTypeSessionCreated HookType = "session-created"
)

var (
	SessionCreatedHook = &Hook{
		Type: HookTypeSessionCreated,
		Cmd:  "run-shell 'tt session add #{hook_session_name} #{session_path}'",
	}
	SessionClosedHook = &Hook{
		Type: HookTypeSessionClosed,
		Cmd:  "run-shell 'tt session rm #{hook_session_name}'",
	}
)
