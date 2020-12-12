$destDir = "D:\Projects\GItHub\InvidositeHtmlgit"
$confirmation = Read-Host "Do you want to publish a new version of vido-preproc on dir $destDir (y/n) ?"
if ($confirmation -ne 'y') {
    Write-Host "Nothing to do. All the best."
    return
}

#create the exe
Write-Host "Build the exe in $pwd"
go build

# Create the Zip package
Write-Host "Create a deploy package"
cd ./deploy
.\deploy.exe -target InvidositeHtmlgit

#Copy the latest zip to the target
Write-Host "Copy the zip to the target"
cd ..\vido-preproc-deployed
.\copy-latest-todest.ps1

#Update the target using update script
Write-Host "Continue with the deploy. Run the update script"
Invoke-Command  -ScriptBlock {cd $destDir; & '.\update-vidopre.ps1'}

Write-Host "That'all, I think."