# ssh-helper
A simple ssh-agent wrapper for human and more

## Usages
- ```-h``` Print help text
- ```-e``` Print $SSH_AUTH_SOCK and $SSH_AGENT_PID
- ```-s``` Start a managed ssh-agent process. Managed means the environment variables will be recorded in config file to ensure only one managed agent is running
- ```-t``` Kill all other ssh-agent processes except managed one
- ```-k``` Kill all ssh-agent processes
- ```-a``` Run ssh-add command after setting environment variables. Only first option after -a is passed to ssh-add

## To-dos
- [ ] Deal with ssh-agent or ssh-add executable not in $PATH
- [ ] Support all ssh-agent and ssh-add options
- [ ] Run as system service
