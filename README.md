# devops-template
From push to k8s deploy

# Create a new ssh key for private use

# add alias for host and which ssh key to use
Host github-private
  HostName github.com
  User git
  IdentityFile /Users/andreas.hansson/.ssh/private
  IdentitiesOnly yes

# git remote add origin 
git remote add origin git@github-private:andreashanson/devops-template.git