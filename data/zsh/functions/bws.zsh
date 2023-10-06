SESSION_FILE=($HOME/.bws-session)

function bwl { 
  if test -f $SESSION_FILE; then
    export BW_SESSION=$(cat $SESSION_FILE)
  fi

  if ! bw status | grep "unlocked" > /dev/null
  then
    export BW_SESSION=${bw unlock --raw}
    bw list items
  else
    echo $BW_SESSION > $SESSION_FILE
  fi
}

function bws {
  bwl
  bw sync
  bw get item "$(bw list items | jq '.[] | "\(.name) | username: \(.login.username) | id: \(.id)" ' \
    | fzf --with-nth 1..-4 | awk '{print $(NF -0)}' | sed 's/\"//g')" \
    | jq -r '.login.password' | sed 's/\"//g' | pbcopy
}


function bw-pass {
  bwl
  bw generate -lusn --length 18
}

function bw-passp { 
  bwl
  bw generate --passphrase --words 3 --separator - --includeNumber --capitalize
}
