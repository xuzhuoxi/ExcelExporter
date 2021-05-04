package cmd

type Flags struct {
	Mode   RunMode
	Source string
	Target string
	Lang   []string
	Field  []string
}
