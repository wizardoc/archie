my $DEST_NAME := "archie";

my $builder = {
    shell "GOOS=linux GOARCH=amd64 go build -o ../build/$DEST_NAME ../main.go"
}

sub MAIN(Str :version($v)){
    if not $v.defined {
        say "Please input version";
        exit;
    }

    say "Begin build...\n";

    $builder();

    say "Begin docker build...\n";

    shell "docker build -t archie:v$v ..";
}