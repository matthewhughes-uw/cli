package cmd

import (
	"fmt"
	"github.com/opslevel/opslevel-go/v2023"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// CLIServiceDependencyCreateInput This is used to make the user facing CLI experience better
// than a straight passthrough to the API types which are overly verbose
type CLIServiceDependencyCreateInput struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Notes  string `json:"notes,omitempty"`
}

var createServiceDependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Create a service dependency",
	Example: `
cat << EOF | opslevel create service dependency -f -
source: my-service-alias # "source" and "target" fields support ID or Alias
target: XXXXXXX
notes: |
  Some extra information about the connection
EOF
`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readCreateServiceDependencyInput()
		cobra.CheckErr(err)
		result, err := getClientGQL().CreateServiceDependency(*input)
		cobra.CheckErr(err)
		fmt.Println(result.Id)
	},
}

var deleteServiceDependencyCmd = &cobra.Command{
	Use:   "dependency ID",
	Short: "Delete a service dependency",
	Example: `
opslevel delete service dependency XXX # ID of the dependency entity returned by the create command
`,
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"ID"},
	Run: func(cmd *cobra.Command, args []string) {
		key := opslevel.NewID(args[0])
		err := getClientGQL().DeleteServiceDependency(*key)
		cobra.CheckErr(err)
		fmt.Printf("deleted '%v' service dependency\n", key)
	},
}

func init() {
	createServiceCmd.AddCommand(createServiceDependencyCmd)
	deleteServiceCmd.AddCommand(deleteServiceDependencyCmd)
}

func readCreateServiceDependencyInput() (*opslevel.ServiceDependencyCreateInput, error) {
	data, err := readCreateFile()
	if err != nil {
		return nil, err
	}
	in := CLIServiceDependencyCreateInput{}
	err = yaml.Unmarshal(data, in)
	if err != nil {
		return nil, err
	}
	output := &opslevel.ServiceDependencyCreateInput{
		Key: opslevel.ServiceDependencyKey{
			Service:   *opslevel.NewIdentifier(in.Source),
			DependsOn: *opslevel.NewIdentifier(in.Target),
		},
		Notes: in.Notes,
	}
	return output, nil
}
