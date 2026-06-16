_cmdctl() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=()
    if [[ $COMP_CWORD -eq 1 ]]; then
        COMPREPLY=($(compgen -W "generate install" -- "$cur"))
    elif [[ $COMP_CWORD -eq 2 ]]; then
        COMPREPLY=($(compgen -W "bash" -- "$cur"))
    fi
}
complete -F _cmdctl cmdctl
