package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

// LogToFile stores logfile details
type LogToFile struct {
	flag     bool
	filename string
}

// URL is a struct to store Cortex XDR urls
type URL struct {
	url string
}

// TableData  stores data from connection results
type TableData struct {
	time   string
	url    string
	status string
}

// PopulateUrls This function builds out the Cortex XDR urls applying region and tenent details within the URLs
func PopulateUrls(region string, tenantname string) []URL {

	var CortexXDRUrls = []URL{
		{"distributions.traps.paloaltonetworks.com"},
		{"dc-" + tenantname + ".traps.paloaltonetworks.com"},
		{"ch-" + tenantname + ".traps.paloaltonetworks.com"},
		{"cc-" + tenantname + ".traps.paloaltonetworks.com"},
		{"lrc-" + region + ".paloaltonetworks.com"},
		{"panw-xdr-installers-prod-us.storage.googleapis.com"},
		{"panw-xdr-payloads-prod-us.storage.googleapis.com"},
		{"global-content-profiles-policy.storage.googleapis.com"},
		{"panw-xdr-evr-prod-" + region + ".storage.googleapis.com"},
	}

	return CortexXDRUrls
}

// TestUrlsGet this function accepts URL then does a connection test. Returns PASS or FAIL using net.DialTimeout
func TestUrlsGet(dataurl []URL) {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initilize a Table
	//table := tablewriter.NewWriter(os.Stdout)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	// Set Table Headers
	table.SetHeader([]string{"TIME", "Cortex XDR URL", "Result"})

	for _, row := range dataurl {
		t := time.Now()
		resp, err := client.Get("https://" + row.url)

		if err != nil {
			// Populate Table rows when connection fails
			d := []string{t.Format("[2006-01-02 15:04:05]"), row.url, "FAIL"}
			table.Append(d)
			continue
		}
		defer resp.Body.Close()

		// Populate Table rowns when connect is succesfull/pass
		d := []string{t.Format("[2006-01-02 15:04:05]"), row.url, "PASS"}
		table.Append(d)

	}
	// Display table to console

	table.Render()
	fmt.Println(tableString.String())
	//if LogfileInfo.flag == true {
	//	err := WriteToFileLog(LogfileInfo.filename, tableString.String())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

}
