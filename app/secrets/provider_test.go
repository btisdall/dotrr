package secrets

import "testing"

func TestPassThruProvider(t *testing.T) {
	c := PassThruProvider{}
	expected := "bar"
	got, _ := c.GetSecret(expected)
	if got != expected {
		t.Errorf("GetSecret()= %v; want %v", got, expected)
	}
}
func TestEscapeProvider(t *testing.T) {
	c := EscapeProvider{}
	in := "\\aws-ssm-parameter:"
	expected := "aws-ssm-parameter:"
	got, _ := c.GetSecret(in)
	if got != expected {
		t.Errorf("GetSecret()= %v; want %v", got, expected)
	}
}

func TestGetProvider(t *testing.T) {
	prefix := "aws-ssm-parameter:"
	p := GetProvider(prefix)
	_, ok := p.(*SSMProvider)
	if !ok {
		t.Errorf("Expected *secrets.SSMProvider got %T", p)
	}
	prefix = "\\aws-ssm-parameter:"
	p = GetProvider(prefix)
	_, ok = p.(*EscapeProvider)
	if !ok {
		t.Errorf("Expected EscapeProvider got %T", p)
	}
	prefix = "test"
	p = GetProvider(prefix)
	_, ok = p.(*PassThruProvider)
	if !ok {
		t.Errorf("Expected *secrets.PassThruProvider got %T", p)
	}
}
