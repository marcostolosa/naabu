package scan

import (
	"time"

	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/networkpolicy"
)

// Options of the scan
type Options struct {
	Timeout              time.Duration
	Retries              int
	Rate                 int
	PortThreshold        int
	ExcludeCdn           bool
	OutputCdn            bool
	ExcludedIps          []string
	Proxy                string
	ProxyAuth            string
	Stream               bool
	OnReceive            result.ResultFn
	ScanType             string
	NetworkPolicyOptions *networkpolicy.Options
}
