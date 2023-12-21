package uuidgen

type MockUUID struct{}

func (*MockUUID) V4() string {
    return "sample-uuid-string"
}
