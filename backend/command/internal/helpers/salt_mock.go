package helpers

type MockSaltGenerator struct {}

func (m *MockSaltGenerator) Generate() (string, error) {
	return "9f86d081884c7d659a2feaa0c55ad015", nil
}