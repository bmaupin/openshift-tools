package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestIngressCleanup(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/ingress-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/ingress-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestSecretCleanup(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/secret-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/secret-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestListCleanup1(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-list-1-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-list-1-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestListCleanup2(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-list-2-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-list-2-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestListCleanupNodejsEx(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-list-nodejs-ex-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-list-nodejs-ex-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestListCleanupNodejsExJson(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-list-nodejs-ex-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-list-nodejs-ex-original.json")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Logf("Expected: \n%s", string(testContents))
		t.Logf("Received: \n%s", string(cleanedContents))
		t.Error("Cleaned list contents don't match")
	}
}

func TestTemplateCleanup1(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-template-1-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-template-1-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Error("Cleaned template contents don't match")
	}
}

func TestTemplateCleanup2(t *testing.T) {
	testContents, err := ioutil.ReadFile("testdata/openshift-template-2-cleaned.yaml")
	if err != nil {
		t.Errorf("Unexpected error reading test data file: %s", err)
	}

	cleanedContents := cleanKubernetesConfigFile("testdata/openshift-template-2-original.yaml")
	if bytes.Compare(cleanedContents, testContents) != 0 {
		t.Error("Cleaned template contents don't match")
	}
}

func TestIdempotency(t *testing.T) {
	const testdataDirectory = "testdata"

	testFiles, err := ioutil.ReadDir(testdataDirectory)
	if err != nil {
		t.Errorf("Unexpected error getting files in testdata directory: %s", err)
	}

	for _, testFile := range testFiles {
		if strings.Contains(testFile.Name(), "-cleaned.") {
			testFilePath := fmt.Sprintf("%s/%s", testdataDirectory, testFile.Name())

			testContents, err := ioutil.ReadFile(testFilePath)
			if err != nil {
				t.Errorf("Unexpected error reading test data file: %s", err)
			}

			t.Logf("Testing against %s for idempotency", testFilePath)
			cleanedContents := cleanKubernetesConfigFile(testFilePath)
			if bytes.Compare(cleanedContents, testContents) != 0 {
				t.Logf("Expected: \n%s", string(testContents))
				t.Logf("Received: \n%s", string(cleanedContents))
				t.Error("Cleaned list contents don't match")
			}
		}
	}
}
