function vp() {
  if [ -z "$1" ]
    then 
      fzf -m | xargs nvim
    else
      nvim $1
  fi
}
