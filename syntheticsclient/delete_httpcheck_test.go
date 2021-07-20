package syntheticsclient

import (
	"net/http"
	"testing"
)

var (
	deleteHttpRespBody = `{"result":"success","message":"testcase successfully deleted","errors":[]}`
)

func TestDeleteHttpCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/http/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(deleteHttpRespBody))
	})

	resp, err := testClient.DeleteHttpCheck(19)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "testcase successfully deleted" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Message, "testcase successfully deleted")
	}
	if resp.Result != "success" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Result, "success")
	}
}
