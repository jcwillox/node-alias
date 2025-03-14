package utils

import "github.com/spf13/cobra"

func CompleteDependencies(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
	pkg := GetPackageJson()
	completions := make([]cobra.Completion, 0, len(pkg.Scripts))

	for dep, version := range pkg.Dependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, dep+":"+version))
	}
	for dep, version := range pkg.DevDependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, dep+":"+version))
	}
	for dep, version := range pkg.PeerDependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, dep+":"+version))
	}

	return completions, cobra.ShellCompDirectiveDefault
}
