package walker

type Walker interface {
	WalkFromRoot() (string, error)
}
