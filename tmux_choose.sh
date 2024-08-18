option=$(gum choose new existing kill)

if [[ $option == "kill" ]]
then
  tmux kill-session -t $(gum choose $(tmux ls -F '#{session_name}'))
  exit
fi

if [[ $option == "new" ]]
then
  session=$(gum input --placeholder "Name of new session")
else
  session=$(gum choose $(tmux ls -F '#{session_name}'))
fi

tmux new-session -A -s $session
