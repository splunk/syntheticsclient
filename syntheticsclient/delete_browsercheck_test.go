package syntheticsclient

import (
	"net/http"
	"testing"
)

var (
	deleteBrowserRespBody = `{"result":"success","message":"testtaste successfully deleted","errors":[]}`
)

func TestDeleteBrowseCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/real_browsers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(deleteBrowserRespBody))
	})

	resp, err := testClient.DeleteBrowserCheck(10)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "testtaste successfully deleted" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Message, "testtaste successfully deleted")
	}
	if resp.Result != "success" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Result, "success")
	}
}
