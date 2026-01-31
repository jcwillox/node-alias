package utils

import "github.com/spf13/cobra"

func CompleteDependencies(_ *cobra.Command, args []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveDefault
	}
	pkg := GetPackageJson()
	if pkg == nil {
		return nil, cobra.ShellCompDirectiveDefault
	}
	completions := make([]cobra.Completion, 0, len(pkg.Scripts))

	for dep, version := range pkg.Dependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, version))
	}
	for dep, version := range pkg.DevDependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, version))
	}
	for dep, version := range pkg.PeerDependencies {
		completions = append(completions, cobra.CompletionWithDesc(dep, version))
	}

	return completions, cobra.ShellCompDirectiveDefault
}
