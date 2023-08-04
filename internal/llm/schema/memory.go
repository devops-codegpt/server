package schema

// Memory is the interface for memory
type Memory interface {
	// MemoryVariables Input keys this memory class will load dynamically
	MemoryVariables() []string
	// LoadMemoryVariables return all memories
	LoadMemoryVariables(inputs map[string]any) (map[string]any, error)
	// SaveContext Save the context of this models run to memory
	SaveContext(inputs map[string]any, outputs map[string]any) error
	// Clear memory contents
	Clear() error
}
