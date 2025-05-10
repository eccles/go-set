for p in golang golangci-lint mockery bats shellcheck shfmt just
do
	asdf plugin add $p
	asdf install $p latest
done
