# Go-lang-quiz
maani@maani-VirtualBox:~/Documents$ git clone https://github.com/mirmujeeb9/Go-lang-quiz
Cloning into 'Go-lang-quiz'...
remote: Enumerating objects: 3, done.
remote: Counting objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (3/3), done.
maani@maani-VirtualBox:~/Documents$ cd Go-lang-quiz
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git branch dev
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git checkout dev
Switched to branch 'dev'
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ gedit blockchain.dev
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ gedit blockchain.go
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git add .
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git commmit -m "dev branch with file"
git: 'commmit' is not a git command. See 'git --help'.

The most similar command is
	commit
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git commit -m "dev branch with file"
[dev fd5db53] dev branch with file
 1 file changed, 78 insertions(+)
 create mode 100644 blockchain.go
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git push origin dev
Username for 'https://github.com': mirmujeeb9
Password for 'https://mirmujeeb9@github.com': 
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 6 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 880 bytes | 440.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
remote: 
remote: Create a pull request for 'dev' on GitHub by visiting:
remote:      https://github.com/mirmujeeb9/Go-lang-quiz/pull/new/dev
remote: 
To https://github.com/mirmujeeb9/Go-lang-quiz
 * [new branch]      dev -> dev
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git checkout main
Switched to branch 'main'
Your branch is up to date with 'origin/main'.
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git merge dev
Updating a48c2b2..fd5db53
Fast-forward
 blockchain.go | 78 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 78 insertions(+)
 create mode 100644 blockchain.go
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git add .
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git commit -m "after merging ,commiting main "
On branch main
Your branch is ahead of 'origin/main' by 1 commit.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ git push origin main
Username for 'https://github.com': mirmujeeb9
Password for 'https://mirmujeeb9@github.com': 
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/mirmujeeb9/Go-lang-quiz
   a48c2b2..fd5db53  main -> main
maani@maani-VirtualBox:~/Documents/Go-lang-quiz$ 

git stash
git pop
git log
