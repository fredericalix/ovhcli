package project

import "github.com/spf13/cobra"

var projectid string

func init() {
	cmdProjectInfo.PersistentFlags().StringVarP(&projectid, "projectid", "", "", "Your ID Project")

}

var cmdProjectInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about a project: ovhcli project info",
	Run: func(cmd *cobra.Command, args []string) {

		/*		c := internal.Client()
				if c == nil {
					os.Exit(1)
				}
				project := internal.Project{}
				path := fmt.Sprintf("/cloud/project/%s", projectid)
				if err := c.Get(path, &project); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Project Name: %s\n", project.Name)
				fmt.Printf("Project Status: %s\n", project.Status)
				fmt.Printf("Creation Date: %s\n", project.CreationDate)
		*/
	},
}
