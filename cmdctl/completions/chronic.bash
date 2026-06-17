_chronic() {
    local CUR="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=()
    # Complete commands until one is seen; options (-ev, --) may precede it.
    local SEEN_COMMAND=false INDEX WORD
    for ((INDEX = 1; INDEX < COMP_CWORD; INDEX++)); do
        WORD="${COMP_WORDS[INDEX]}"
        if [[ "$WORD" == "--" ]]; then
            continue
        fi
        if [[ "$WORD" != -* ]]; then
            SEEN_COMMAND=true
            break
        fi
    done
    if [[ "$SEEN_COMMAND" == false ]]; then
        if [[ "$CUR" == -* ]]; then
            COMPREPLY=($(compgen -W "-e -v -ev --" -- "$CUR"))
        else
            COMPREPLY=($(compgen -c -- "$CUR"))
        fi
    fi
}
complete -o default -F _chronic chronic
