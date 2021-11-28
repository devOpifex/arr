package cmd

import "github.com/spf13/cobra"

var packerCmd = &cobra.Command{
	Use:   "packer",
	Short: "interact with the {packer} package",
	Long: `bundle, test, and more from the command line.
Examples:
	arr packer bundle
	`,
	Args: cobra.MinimumNArgs(1),
	ValidArgs: []string{
		"bundle",
		"bundle_dev",
		"bundle_prov",
		"set_npm",
		"set_yarn",
		"loader_style",
		"loader_svelte",
		"loader_vue",
		"loader_coffee",
		"loader_css",
		"loader_rule",
		"loader_pug",
		"loader_sass",
		"Loader_framework7",
		"loader_mocha",
		"loader_babel",
		"loader_file",
		"loader_ts",
		"plugin_workbox",
		"plugin_html",
		"plugin_prettier",
		"plugin_tutorial",
		"plugin_eslint",
		"plugin_clean",
		"plugin_jsdoc",
		"tests_peeky",
		"tests_mocha",
		"action_check",
		"test",
		"precommit_hook",
		"rprofile_adapt",
		"scaffold_rmd",
		"scaffold_ambiorix",
		"scaffold_widget",
		"scaffold_output",
		"scaffold_bare",
		"scaffold_golem",
		"scaffold_extension",
		"scaffold_leprechaun",
		"scaffold_input",
	},
	Run: makeRun("packer", "-e"),
}

func init() {
	rootCmd.AddCommand(packerCmd)
}
