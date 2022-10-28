# psuedo code written on airplane

secret = ask-lastpass
projectConfigLoc = $Gitbasedir + '/project_config'
secretUID = tell-keeper --name $name --password $secret --folder $folderUID

#add new lines to project_config for keeper
tee -a $projectConfigLoc << END
keeperUID = $secretUID
END

#test if you can pull the secret from Keeper
ask-keeper --secret $secretUID
# if successful, remove old lastpass lines from project config
sed '/lastpass_uid/d' $projectConfigLoc

# create commit
git -add $projectConfigLoc
git commit -m "Update vault password from LastPass to Keeper"
gh pr create

