package tabletserver

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/youtube/vitess/go/acl"
)

var (
	streamqueryzHeader = []byte(`<thead>
		<tr>
			<th>Query</th>
			<th>RemoteAddr</th>
			<th>Username</th>
			<th>Duration</th>
			<th>Start</th>
			<th>SessionID</th>
			<th>TransactionID</th>
			<th>ConnectionID</th>
			<th>Current State</th>
		</tr>
        </thead>
	`)
	streamqueryzTmpl = template.Must(template.New("example").Parse(`
		<tr> 
			<td>{{.Query}}</td>
			<td>{{.RemoteAddr}}</td>
			<td>{{.Username}}</td>
			<td>{{.Duration}}</td>
			<td>{{.Start}}</td>
			<td>{{.SessionID}}</td>
			<td>{{.TransactionID}}</td>
			<td>{{.ConnID}}</td>
			<td>{{.State}}</td>
		</tr>
	`))
)

func init() {
	http.HandleFunc("/streamqueryz", streamqueryzHandler)
	http.HandleFunc("/streamqueryz/terminate", streamqueryzTerminateHandler)
}

func streamqueryzHandler(w http.ResponseWriter, r *http.Request) {
	if err := acl.CheckAccessHTTP(r, acl.DEBUGGING); err != nil {
		acl.SendError(w, err)
		return
	}
	startHTMLTable(w)
	defer endHTMLTable(w)
	w.Write(streamqueryzHeader)
	rows := SqlQueryRpcService.qe.streamQList.GetQueryzRows()
	for i := range rows {
		streamqueryzTmpl.Execute(w, rows[i])
	}
}

func streamqueryzTerminateHandler(w http.ResponseWriter, r *http.Request) {
	if err := acl.CheckAccessHTTP(r, acl.DEBUGGING); err != nil {
		acl.SendError(w, err)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form: %s", err), http.StatusInternalServerError)
		return
	}
	connID := r.FormValue("connID")
	c, err := strconv.Atoi(connID)
	if err != nil {
		http.Error(w, "invalid connID", http.StatusInternalServerError)
		return
	}
	if qd := SqlQueryRpcService.qe.streamQList.Get(int64(c)); qd != nil {
		qd.Terminate()
	}
	streamqueryzHandler(w, r)
}
