#!/bin/bash
source typer.sh

typer 'alias trhead="trelldo get card -b Work -l Stack --head 2>/dev/null"'
enter
sleep 2
typer 'alias trpop="trelldo delete card -b Work -l Stack --head 2>/dev/null"'
enter
sleep 2
typer 'trpush () { trelldo create card -b Work -l Stack --head "$*" 2>/dev/null }'
enter
sleep 2
typer 'trpush "Task 1"'
enter
sleep 2
typer 'trhead'
enter
sleep 2
typer 'trpush "Task 2"'
enter
sleep 2
typer 'trhead'
enter
sleep 2
typer 'trpop'
enter
sleep 2
typer 'trpop'
enter
sleep 2
typer "exit"
enter
