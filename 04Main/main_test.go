package maintest

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Perform common tasks before testing")
	exitVal := m.Run()
	log.Println("Perform common tasks after testing!")

	//exit program with given status code
	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("Running Test A")
}

func TestB(t *testing.T) {
	log.Println("Running Test B")
}
