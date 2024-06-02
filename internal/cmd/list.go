package cmd

import (
	"os"

	humanize "github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"

	"github.com/EmilyShepherd/ota-tool/pkg/payload"
)

type List struct {
}

func (l *List) Usage() string {
	return "List the partitions found in the payload"
}

func (l *List) Execute(f []string, update *payload.Payload) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Partition", "Old Size", "New Size"})
	for _, partition := range update.Partitions {
		table.Append([]string{
			partition.GetPartitionName(),
			humanize.Bytes(*partition.GetOldPartitionInfo().Size),
			humanize.Bytes(*partition.GetNewPartitionInfo().Size),
		})
	}
	table.SetBorder(false)
	table.Render()
}

