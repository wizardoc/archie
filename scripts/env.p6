say "Switch mode";

my Str $answer = prompt "Switch on release mode? (Y/ N): ";

loop {
    if !$answer.defined || $answer eq "" {
        $answer = prompt "Please choose Y or N: ";

        next;
    }

    last;
}

my Str $mode = $answer.uc eq "Y" ?? "release" !! "debug";

shell "export GIN_MODE=$mode";

say "$mode mode";