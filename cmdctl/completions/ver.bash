_ver() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=()

    if [[ $COMP_CWORD -eq 1 ]]; then
        COMPREPLY=($(compgen -W "major minor patch" -- "$cur"))
    fi
}
complete -F _ver ver
