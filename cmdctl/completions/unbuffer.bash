_unbuffer() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=()
    if [[ $COMP_CWORD -eq 1 ]]; then
        COMPREPLY=($(compgen -c -- "$cur"))
    fi
}
complete -o default -F _unbuffer unbuffer
