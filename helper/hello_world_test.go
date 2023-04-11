package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("Khafid")
	assert.Equal(t, "Hello Khafid", result, "Result must be 'Hello Khafid'")
	fmt.Println("Hello Khafid With Assert")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("Khafid")
	require.Equal(t, "Hello Khafid", result, "Result must be 'Hello Khafid'")
	fmt.Println("Hello Khafid With Require")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Azhar")
	if result != "Hello Azhar" {
		//unit test failed
		// panic("Result is not 'Hello Azhar'")
		t.Fail()
	}

	fmt.Println("Hello Azhar")
}

func TestHelloWorldAzhar(t *testing.T) {
	result := HelloWorld("Majid")
	if result != "Hello Majid" {
		//unit test failed
		// panic("Result is not 'Hello Azhar'")
		t.FailNow()
	}

	fmt.Println("Hello Majid")

}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		//unit test failed
		// panic("Result is not 'Hello Azhar'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Eko")
	}

	fmt.Println("Hello Majid")

}
func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello Ilham" {
		//unit test failed
		// panic("Result is not 'Hello Azhar'")
		// t.FailNow()
		t.Error("Result must be 'Hello Ilham")
	}

	fmt.Println("Hello Majid")

}
