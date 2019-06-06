my $DIST_NAME := "archie";
my $DIST_DIR := "../build";

my $builder = {
    shell "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $DIST_DIR/$DIST_NAME ../main.go"
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