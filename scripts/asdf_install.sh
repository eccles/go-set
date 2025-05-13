for p in golang golangci-lint just
do
	asdf plugin add $p
	asdf install $p latest
done
