### SSHCheck
This utility queries GitHub for a users API keys and writes them to a file that can be used an ssh `authorized_hosts` file. This allows for making sure that your devices will always have to up to date keys based on what's in your GitHub profile.

**WARNING**
If someone is able to inject an SSH key into your GitHub profile, they will be able to remote in any device you have with this script.