package core

type Commit struct {
	reporsitory *Repository
	hash        string
	Parent      *Commit
	Tree        *Tree
	Message     string
}

func (c Commit) Repository() *Repository {
	return c.reporsitory
}

func (c Commit) Hash() string {
	return c.hash
}

func (c Commit) Type() string {
	return "Commit"
}
